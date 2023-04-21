// Package main in app-init.go is the generation script utilizing lib-instance-gen
package main

//go:generate go run app-init.go

import instance_gen "github.com/skeletonkey/lib-instance-gen-go/app"

func main() {
	app := instance_gen.NewApp("rachio-next-run", "app")
	app.WithPackages("logger", "pushover", "rachio").
		WithGithubWorkflows("linter", "test").
		WithMakefile()
}
