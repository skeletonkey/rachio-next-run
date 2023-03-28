// instance-gen: File auto generated -- DO NOT EDIT!!!
package rachio

import "rachionextrun/app/config"

var client *rachio

func getClient() *rachio {
	if client == nil {
		config.LoadConfig("rachio", &client)
	}
	return client
}
