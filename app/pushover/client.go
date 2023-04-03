// instance-gen: File auto generated -- DO NOT EDIT!!!
package pushover

import "rachionextrun/app/config"

var client *pushover

func getConfig() *pushover {
	if client == nil {
		config.LoadConfig("pushover", &client)
	}
	return client
}
