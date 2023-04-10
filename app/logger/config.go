// instance-gen: File auto generated -- DO NOT EDIT!!!
package logger

import "github.com/skeletonkey/rachio-next-run/app/config"

var cfg *logger

func getConfig() *logger {
	config.LoadConfig("logger", &cfg)
	return cfg
}
