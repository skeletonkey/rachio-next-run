// lib-instance-gen-go: File auto generated -- DO NOT EDIT!!!
package rachio

import "github.com/skeletonkey/lib-core-go/config"

var cfg *rachio

func getConfig() *rachio {
	config.LoadConfig("rachio", &cfg)
	return cfg
}
