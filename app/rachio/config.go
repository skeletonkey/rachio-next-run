// instance-gen: File auto generated -- DO NOT EDIT!!!
package rachio

import "github.com/skeletonkey/rachio-next-run/app/config"

var cfg *rachio

func getConfig() *rachio {
	config.LoadConfig("rachio", &cfg)
	return cfg
}
