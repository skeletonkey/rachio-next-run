package pushover

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"rachionextrun/app/config"
)

type token struct {
	Account     string `json:"account"`
	Application string `json:"application"`
}
type pushover struct {
	Url   string `json:"url"`
	Token token  `json:"token"`
}

var client *pushover

func getClient() *pushover {
	if client == nil {
		config.LoadConfig("pushover", &client)
	}
	return client
}

func Notify(msg string) {
	client := getClient()
	requestUrl := fmt.Sprintf("%s/messages.json?token=%s&user=%s&message=%s",
		client.Url, client.Token.Account, client.Token.Application, url.QueryEscape(msg))
	fmt.Printf("Notification URL: %s", requestUrl)
	res, err := http.Post(requestUrl, "application/json", nil)
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	if res.StatusCode != 200 {
		panic(fmt.Errorf("non 200 code received from Notify call: (%d) %s", res.StatusCode, string(body[:])))
	}

	fmt.Printf("Response: %s\n", string(body[:]))
}
