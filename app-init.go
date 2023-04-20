// Package main in app-init.go is the generation script utilizing lib-instance-gen
package main

//go:generate go run app-init.go

import instance_gen "github.com/skeletonkey/rachio-next-run/lib/instance-gen"

func main() {
	app := instance_gen.NewApp("app")
	app.WithModules("logger", "pushover", "rachio")
}
