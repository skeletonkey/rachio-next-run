package main

import (
	"fmt"
	"rachionextrun/app/config"
	"rachionextrun/app/pushover"
	"rachionextrun/app/rachio"
)

type app struct {
	LogLevel string `json:"log_level"`
}

func main() {
	var appData app
	config.LoadConfig("app", &appData)
	fmt.Printf("App: %+v\n", appData)
	timeUntil, alertType, alert := rachio.GetNextScheduledRun()
	fmt.Printf("Making the call to rachio: %d - %s - %t\n", timeUntil, alertType, alert)
	if alert {
		if alertType == "after" {
			pushover.Notify("Scheduled watering has completed")
		} else {
			pushover.Notify(fmt.Sprintf("Next scheduled watering is %d hour(s) away", timeUntil))
		}
	}
}
