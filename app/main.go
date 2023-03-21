package main

import (
	"fmt"
	"rachionextrun/app/config"
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
}
