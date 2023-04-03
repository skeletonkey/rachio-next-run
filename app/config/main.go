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

const configFileString = "RACHIO_CONFIG_FILE"

var cfg *config
var lock = &sync.Mutex{}

func (c config) getConfigFile() string {
	if c.configFile == "" {
		// TODO: better way to get config file location
		if filename := os.Getenv(configFileString); filename == "" {
			panic(fmt.Errorf("env var %s is not set", configFileString))
		} else {
			c.configFile = filename
		}
	}
	return c.configFile
}

// getConfig returns the internal cfg object (creating it if needed)
// This method is usally called first and the application can not run without this information.  Since any errors
// encountered here are faital, panic is used instead of any type of error return or logging.
func getConfig() *config {
	if cfg == nil {
		lock.Lock()
		defer lock.Unlock()

		cfg = &config{}

		cfg.configs = make(configMapType)

		rawData, err := os.ReadFile(cfg.getConfigFile())
		if err != nil {
			panic(fmt.Errorf("unable to open config file (%s): %s", cfg.getConfigFile(), err))
		}

		if !json.Valid(rawData) {
			panic(fmt.Errorf("invalid JSON found in file (%s)", cfg.getConfigFile()))
		}

		data := map[string]interface{}{}
		err = json.Unmarshal(rawData, &data)
		if err != nil {
			panic(fmt.Errorf("unable to unmarshal config file (%s): %s", cfg.getConfigFile(), err))
		}

		for key, value := range data {
			valueJson, err := json.Marshal(value)
			if err != nil {
				panic(fmt.Errorf("unable to marchal key (%s) data: %s", key, err))
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
		panic(fmt.Errorf("key (%s) not found in config file (%s)", name, cfg.getConfigFile()))
	}
	json.Unmarshal(configData, &configStruct)
}

// #1 https://www.sohamkamani.com/golang/json/
// #2 https://blog.boot.dev/golang/anonymous-structs-golang/
// #3 https://refactoring.guru/design-patterns/singleton/go/example#example-0
