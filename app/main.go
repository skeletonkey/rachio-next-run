// Package main is the entry point for the application
package main

import (
	"fmt"
	"time"

	"github.com/skeletonkey/lib-core-go/config"
	"github.com/skeletonkey/lib-core-go/logger"
	"github.com/skeletonkey/rachio-next-run/app/pushover"
	"github.com/skeletonkey/rachio-next-run/app/rachio"
)

type app struct {
}

func main() {
	var appData app
	config.LoadConfig("app", &appData)

	log := logger.Get()
	log.Info().Msg("Starting app")

	healthcheck.Healthcheck()

	time.Sleep(10 * time.Minute)
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
