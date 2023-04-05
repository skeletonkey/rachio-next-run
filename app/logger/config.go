// instance-gen: File auto generated -- DO NOT EDIT!!!
package logger

import "rachionextrun/app/config"

var cfg *logger
var initialized bool

func getConfig() *logger {
	config.LoadConfig("logger", &cfg)
	return cfg
}

func reInitialize() bool {
	return config.Reset("logger")
}
