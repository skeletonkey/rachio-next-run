package pushover

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"rachionextrun/app/logger"
)

// Notify sends `msg` using the Pushover API
func Notify(msg string) {
	client := getConfig()
	log := logger.Get()
	requestURL := fmt.Sprintf("%s/messages.json?token=%s&user=%s&message=%s",
		client.URL, client.Token.Account, client.Token.Application, url.QueryEscape(msg))
	log.Debug().Str("URL", requestURL).Msg("notification URL")
	res, err := http.Post(requestURL, "application/json", nil)
	if err != nil {
		log.Panic().Err(err).Str("URL", requestURL).Msg("unable to post to url")
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			log.Error().
				Err(err).
				Msg("unable to close response body")
		}
	}()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Panic().Err(err).Interface("response", res).Msg("unable to read response body")
	}
	if res.StatusCode != 200 {
		log.Panic().Int("Status Code", res.StatusCode).Bytes("response body", body).Msg("non-200 response received")
	}

	log.Info().Bytes("response body", body).Msg("pushover response")
}
