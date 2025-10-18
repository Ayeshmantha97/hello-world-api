package adapters

import (
	"hello-world-api/app/config"
	"hello-world-api/domain/boundary/adapters"
)

type YamlConfigAdapter struct {
}

func NewYamlConfigAdapter() adapters.ConfigAdapterInterface {
	return &YamlConfigAdapter{}
}

func (y *YamlConfigAdapter) Read() *config.Config {
	cfg := &config.Config{}
	cfg.AppConfig = config.ReadAppConfig()

	return cfg
}
