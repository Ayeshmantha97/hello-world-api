package container

import (
	"hello-world-api/app/config"
	"hello-world-api/externals/adapters"
)

func ResolveConfig() (cfg *config.Config) {
	configAdapter := adapters.NewYamlConfigAdapter()

	cfg = configAdapter.Read()

	return cfg
}
