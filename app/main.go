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

	for _, nextRun := range rachio.GetNextScheduledRuns() {
		log.Debug().
			Str("device name", nextRun.DeviceName).
			Int("hours until next run", nextRun.HoursUntil).
			Str("alert type", nextRun.AlertType).
			Bool("Notify", nextRun.Alert).
			Msg("Rachio data")
		if nextRun.Alert {
			if nextRun.AlertType == "after" {
				pushover.Notify(fmt.Sprintf("%s: Scheduled watering has completed", nextRun.DeviceName))
			} else {
				pushover.Notify(fmt.Sprintf("%s: Next scheduled watering is %d hour(s) away", nextRun.DeviceName, nextRun.HoursUntil))
			}
		}
	}

	log.Info().Msg("Finished")
}
