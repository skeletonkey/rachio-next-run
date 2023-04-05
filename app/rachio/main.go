// Package rachio provides an internal interface to the Rachio API
package rachio

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"rachionextrun/app/logger"
	"time"
)

// GetNextScheduledRuns is overloaded and will be replaced
// returns number of hours, 'after' or 'before', bool to indicate if you should alert or not
func GetNextScheduledRuns() []NextScheduleData {
	config := getConfig()

	scheduleData := make([]NextScheduleData, 0)

	for _, device := range config.Devices {
		diff, alertType, alert := getNextScheduleData(device)
		scheduleData = append(scheduleData,
			NextScheduleData{device.Name, diff, alertType, alert})
	}

	return scheduleData
}

func getNextScheduleData(d device) (diffHrs int, alertType string, alert bool) {
	log := logger.Get()
	config := getConfig()

	url := fmt.Sprintf("%s/device/getDeviceState/%s", client.URL.Internal, d.ID)
	log.Debug().Str("URL", url).Msg("Connect to Rachio")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Panic().Err(err).Str("URL", url).Msg("unable to create new http request")
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.BearerToken))
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panic().Err(err).Str("URL", url).Msg("unable to execute http request")
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			log.Error().Interface("response", res).Err(err).Msg("Unable to close body for response")
		}
	}()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Panic().Err(err).Interface("response", res).Msg("unable to read response Body")
	}
	if res.StatusCode != 200 {
		log.Panic().Int("status code", res.StatusCode).Bytes("body", body).Msg("non-200 response code")
	}

	log.Debug().Bytes("response", body).Msg("Rachio response received")
	var nextRunData nextRun
	err = json.Unmarshal(body, &nextRunData)
	if err != nil {
		log.Panic().
			Bytes("response", body).
			Err(err).
			Msg("unable to unmarshal body")
	}
	log.Debug().Interface("un-marshalled body", nextRunData).Msg("Parsed body")

	if d.ID != nextRunData.State.DeviceID {
		log.Panic().
			Str("local device ID", d.ID).
			Str("response device ID", nextRunData.State.DeviceID).
			Msg("devices IDs mismatch")
	}

	t, err := time.Parse(time.RFC3339, nextRunData.State.NextRun)
	if err != nil {
		log.Panic().
			Err(err).
			Str("next run", nextRunData.State.NextRun).
			Msg("unable to parse time/date")
	}

	curTime := time.Now()
	var diff time.Duration
	if curTime.After(t) {
		diff = curTime.Sub(t)
		alertType = "after"
		diffHrs = int(diff.Hours())
	} else {
		diff = t.Sub(curTime)
		alertType = "before"
		diffHrs = int(diff.Hours())
	}
	log.Info().Str("current time", curTime.String()).
		Str("next run time", t.String()).
		Dur("difference microseconds", diff).
		Int("difference hours", diffHrs).
		Msg("Rachio actionable information")
	if diff < time.Duration(d.Hours[alertType])*time.Hour {
		alert = true
	}

	return diffHrs, alertType, alert
}
