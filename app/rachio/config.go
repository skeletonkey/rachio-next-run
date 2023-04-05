// instance-gen: File auto generated -- DO NOT EDIT!!!
package rachio

import "rachionextrun/app/config"

var cfg *rachio

func getConfig() *rachio {
	config.LoadConfig("rachio", &cfg)
	return cfg
}

func reInitialize() bool {
	return config.Reset("rachio")
}
