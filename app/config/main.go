package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type configMapType map[string][]byte
type config struct {
	configFile string        // location of the config json file
	configs    configMapType // map containing json representation of each high level key
}

var cfg *config

var lock = &sync.Mutex{}

func (c config) getConfigFile() string {
	if c.configFile == "" {
		// TODO: better way to get config file location
		c.configFile = os.Getenv("RACHIO_CONFIG_FILE")
	}
	return c.configFile
}

func getConfig() *config {
	if cfg == nil {
		lock.Lock()
		defer lock.Unlock()

		cfg = &config{}

		cfg.configs = make(configMapType)

		rawData, err := os.ReadFile(cfg.getConfigFile())
		if err != nil {
			//	TODO: logging will happen later
			panic(err)
		}

		if !json.Valid(rawData) {
			// TODO: logging
			panic(fmt.Errorf("invalid JSON in %s", cfg.getConfigFile()))
		}

		data := map[string]interface{}{}
		err = json.Unmarshal(rawData, &data)
		if err != nil {
			//	TODO: logging will happen later
			panic(err)
		}

		for key, value := range data {
			valueJson, err := json.Marshal(value)
			if err != nil {
				//	TODO: logging will happen later
				panic(err)
			}

			cfg.configs[key] = valueJson
		}
	}

	return cfg
}

func LoadConfig(name string, configStruct interface{}) {
	cfg = getConfig()

	configData, ok := cfg.configs[name]
	if !ok {
		panic(fmt.Errorf("%s does not exist in the config file %s", name, cfg.configFile))
	}
	json.Unmarshal(configData, &configStruct)
}

// #1 https://www.sohamkamani.com/golang/json/
// #2 https://blog.boot.dev/golang/anonymous-structs-golang/
// #3 https://refactoring.guru/design-patterns/singleton/go/example#example-0