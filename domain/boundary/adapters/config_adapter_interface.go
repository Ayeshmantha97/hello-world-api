package adapters

import "hello-world-api/app/config"

type ConfigAdapterInterface interface {
	Read() *config.Config
}
