package {{ . }}

import "rachionextrun/app/config"

var client *{{ . }}

func getConfig() *{{ . }} {
	if client == nil {
		config.LoadConfig("{{ . }}", &client)
	}
	return client
}
