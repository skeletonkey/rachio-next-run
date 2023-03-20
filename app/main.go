package main

import (
	"fmt"
	"rachionextrun/app/config"
)

type app struct {
	LogLevel string `json:"log_level"`
}

func main() {
	var appData app
	config.LoadConfig("app", &appData)
	fmt.Printf("App: %+v\n", appData)
}
