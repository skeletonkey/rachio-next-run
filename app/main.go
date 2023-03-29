package main

import (
	"fmt"
	"rachionextrun/app/config"
	"rachionextrun/app/logger"
	"rachionextrun/app/pushover"
	"rachionextrun/app/rachio"
)

type app struct {
}

func main() {
	var appData app
	config.LoadConfig("app", &appData)

	log := logger.Get()
	log.Info().Msg("Starting app")

	log.Debug().Interface("application Config", appData).Msg("App Data")
	timeUntil, alertType, alert := rachio.GetNextScheduledRun()
	log.Debug().Int("hours until next run", timeUntil).
		Str("alert type", alertType).
		Bool("Notify", alert).
		Msg("Rachio data")
	if alert {
		if alertType == "after" {
			pushover.Notify("Scheduled watering has completed")
		} else {
			pushover.Notify(fmt.Sprintf("Next scheduled watering is %d hour(s) away", timeUntil))
		}
	}

	log.Info().Msg("Finished")
}
