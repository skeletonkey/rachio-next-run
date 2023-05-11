// Package rachio provides an internal interface to the Rachio API
package rachio

import (
	"context"
	"time"

	libRachio "github.com/skeletonkey/rachio"

	"github.com/skeletonkey/lib-core-go/logger"
)

// GetNextScheduledRuns is overloaded and will be replaced
// returns number of hours, 'after' or 'before', bool to indicate if you should alert or not
func GetNextScheduledRuns() []NextScheduleData {
	config := getConfig()
	if !config.Enabled {
		log := logger.Get()
		log.Info().Msg("Rachio is disabled")
		return nil
	}

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

	rachioClient := libRachio.New(context.TODO(), config.BearerToken)
	deviceState, err := rachioClient.ExpGetDeviceState(d.ID)
	if err != nil {
		log.Panic().
			Str("device ID", d.ID).
			Err(err).
			Msg("Error from ExpGetDeviceState call")
	}
	var nextRunData nextRun

	if d.ID != deviceState.State.DeviceId {
		log.Panic().
			Str("local device ID", d.ID).
			Str("response device ID", nextRunData.State.DeviceID).
			Msg("devices IDs mismatch")
	}

	curTime := time.Now()
	var diff time.Duration
	if curTime.After(deviceState.State.NextRun) {
		diff = curTime.Sub(deviceState.State.NextRun)
		alertType = "after"
		diffHrs = int(diff.Hours())
	} else {
		diff = deviceState.State.NextRun.Sub(curTime)
		alertType = "before"
		diffHrs = int(diff.Hours())
	}
	log.Info().Str("current time", curTime.String()).
		Str("next run time", deviceState.State.NextRun.String()).
		Dur("difference microseconds", diff).
		Int("difference hours", diffHrs).
		Msg("Rachio actionable information")
	if diff < time.Duration(d.Hours[alertType])*time.Hour {
		alert = true
	}

	return diffHrs, alertType, alert
}
