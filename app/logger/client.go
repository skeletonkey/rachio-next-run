// instance-gen: File auto generated -- DO NOT EDIT!!!
package logger

import "rachionextrun/app/config"

var client *logger

func getConfig() *logger {
	if client == nil {
		config.LoadConfig("logger", &client)
	}
	return client
}
