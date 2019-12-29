package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Config is the struct for the
type Config struct {
	SSL         bool   `json:"ssl"`
	DefaultURL  string `json:"default_url"`
	DefaultPort string `json:"default_port"`
	Rules       []Rule `json:"rules"`
}

// Rule define the struct for a single rule
type Rule struct {
	Matcher        string `json:"matcher"`
	DownstreamPort string `json:"downstream_port"`
}

// InitConfig initialise the configuration
func InitConfig(configPath string) Config {
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
