package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

const defaultConfigPath = "config.json"

// Config is the struct for the
type Config struct {
	SSL         bool   `json:"ssl"`
	DefaultURL  string `json:"default_url"`
	DefaultPort int64  `json:"default_port"`
	Rules       []Rule `json:"rules"`
}

// Rule define the struct for a single rule
type Rule struct {
	Matcher        string `json:"matcher"`
	DownstreamURL  string `json:"downstream_url"`
	DownstreamPort int64  `json:"downstream_port"`
}

// InitConfig initialise the configuration
func InitConfig(configPath string) Config {
	// When the config path is not passed we set as default the `defaultConfigPath`
	if configPath == "" {
		configPath = defaultConfigPath
	}

	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	var config Config

	err = json.Unmarshal([]byte(file), &config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
