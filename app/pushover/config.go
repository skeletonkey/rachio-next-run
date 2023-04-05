// instance-gen: File auto generated -- DO NOT EDIT!!!
package pushover

import "rachionextrun/app/config"

var cfg *pushover

func getConfig() *pushover {
	config.LoadConfig("pushover", &cfg)
	return cfg
}

func reInitialize() bool {
	return config.Reset("pushover")
}
