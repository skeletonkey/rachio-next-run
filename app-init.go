package main

//go:generate go run app-init.go

import "rachionextrun/lib/instance-gen"

func main() {
	app := instance_gen.NewApp("app")
	app.WithClients("logger", "pushover", "rachio")
}
