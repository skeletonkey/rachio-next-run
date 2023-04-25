// lib-instance-gen-go: File auto generated -- DO NOT EDIT!!!
package logger

import "github.com/skeletonkey/lib-core-go/config"

var cfg *logger

func getConfig() *logger {
	config.LoadConfig("logger", &cfg)
	return cfg
}
