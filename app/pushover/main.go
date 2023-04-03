package pushover

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"rachionextrun/app/logger"
)

func Notify(msg string) {
	client := getConfig()
	log := logger.Get()
	requestUrl := fmt.Sprintf("%s/messages.json?token=%s&user=%s&message=%s",
		client.Url, client.Token.Account, client.Token.Application, url.QueryEscape(msg))
	log.Debug().Str("URL", requestUrl).Msg("notification URL")
	res, err := http.Post(requestUrl, "application/json", nil)
	if err != nil {
		log.Panic().Err(err).Str("URL", requestUrl).Msg("unable to post to url")
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Panic().Err(err).Interface("response", res).Msg("unable to read response body")
	}
	if res.StatusCode != 200 {
		log.Panic().Int("Status Code", res.StatusCode).Bytes("response body", body).Msg("non-200 response received")
	}

	log.Info().Bytes("response body", body).Msg("pushover response")
}
