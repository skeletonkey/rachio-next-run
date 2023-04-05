package {{ . }}

import "rachionextrun/app/config"

var cfg *{{ . }}

func getConfig() *{{ . }} {
	config.LoadConfig("{{ . }}", &cfg)
	return cfg
}

func reInitialize() bool {
	return config.Reset("{{ . }}")
}
