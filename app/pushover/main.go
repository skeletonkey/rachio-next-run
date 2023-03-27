package pushover

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

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
