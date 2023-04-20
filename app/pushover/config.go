// instance-gen: File auto generated -- DO NOT EDIT!!!
package pushover

import "github.com/skeletonkey/rachio-next-run/app/config"

var cfg *pushover

func getConfig() *pushover {
	config.LoadConfig("pushover", &cfg)
	return cfg
}
