package main

import (
	"rachionextrun/app/config"
)

type app struct {
	LogLevel string `json:"log_level"`
}

func main() {
	var appData app
	config.LoadConfig("app", &appData)
}
