// instance-gen: File auto generated -- DO NOT EDIT!!!
package healthcheck

import "rachionextrun/app/config"

var client *healthcheck

func getConfig() *healthcheck {
	if client == nil {
		config.LoadConfig("healthcheck", &client)
	}
	return client
}
