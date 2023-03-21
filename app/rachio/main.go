// Package rachio provides an internal interface to the Rachio API
package rachio

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"rachionextrun/app/config"
	"time"
)

type url struct {
	Public   string `json:"public"`
	Internal string `json:"internal"`
}
type device struct {
	Name  string         `json:"name"`
	Id    string         `json:"device_id"`
	Hours map[string]int `json:"hours"`
}
type rachio struct {
	Url         url      `json:"url"`
	BearerToken string   `json:"bearer_token"`
	Devices     []device `json:"devices"`
}

var client *rachio

func getClient() *rachio {
	if client == nil {
		config.LoadConfig("rachio", &client)
	}
	return client
}

type deviceStateType struct {
	DeviceId string `json:"deviceId"`
	NextRun  string `json:"nextRun"`
}
type nextRun struct {
	State deviceStateType `json:"state"`
}

// GetNextScheduledRun is overloaded and will be replaced
// returns number of hours, 'after' or 'before', bool to indicate if you should alert or not
func GetNextScheduledRun() (diffHrs int, alertType string, alert bool) {
	client := getClient()
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/device/getDeviceState/%s", client.Url.Internal, client.Devices[0].Id), nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.BearerToken))
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	if res.StatusCode != 200 {
		panic(fmt.Errorf("non 200 code received from GetNextScheduledRun call: (%d) %s", res.StatusCode, string(body[:])))
	}

	fmt.Printf("Response: %s\n", string(body[:]))
	var nextRunData nextRun
	json.Unmarshal(body, &nextRunData)
	fmt.Printf("Unmarshalled: %+v\n", nextRunData)

	if client.Devices[0].Id != nextRunData.State.DeviceId {
		panic(fmt.Errorf("device IDs do not match: %s - %s", client.Devices[0].Id, nextRunData.State.DeviceId))
	}

	t, err := time.Parse(time.RFC3339, nextRunData.State.NextRun)
	if err != nil {
		panic(err)
	}

	curTime := time.Now()
	var diff time.Duration
	if curTime.After(t) {
		diff = curTime.Sub(t)
		alertType = "after"
		diffHrs = int(diff / time.Hour)
	} else {
		diff = t.Sub(curTime)
		alertType = "before"
		diffHrs = int(diff / time.Hour)
	}
	fmt.Printf("Current time: %s\nNext Run time: %s\nDifference: %d (%d Hours)\n", curTime.String(), t.String(), diff, diffHrs)
	if diff < time.Duration(client.Devices[0].Hours[alertType])*time.Hour {
		alert = true
	}

	return
}
