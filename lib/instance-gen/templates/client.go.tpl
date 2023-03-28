package {{ . }}

import "rachionextrun/app/config"

var client *{{ . }}

func getClient() *{{ . }} {
	if client == nil {
		config.LoadConfig("{{ . }}", &client)
	}
	return client
}
