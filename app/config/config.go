package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	AppConfig AppConfig
}

type AppConfig struct {
	Port int `yaml:"port"`
}

func ReadAppConfig() (appConfig AppConfig) {
	config, err := os.ReadFile("config/app.yaml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	err = yaml.Unmarshal(config, &appConfig)
	if err != nil {
		log.Fatalf("Error parsing YAML file: %v", err)
	}

	return appConfig
}
