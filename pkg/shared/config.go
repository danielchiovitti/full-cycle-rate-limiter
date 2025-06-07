package shared

import (
	"github.com/kelseyhightower/envconfig"
	_ "github.com/kelseyhightower/envconfig"
	"log"
	"sync"
)

var lockConfig sync.Mutex
var configInstance *Config

func NewConfig() ConfigInterface {
	if configInstance == nil {
		lockConfig.Lock()
		defer lockConfig.Unlock()
		if configInstance == nil {
			configInstance = &Config{}
			err := envconfig.Process("", configInstance)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return configInstance
}

type Config struct {
	Port uint16 `envconfig:"PORT" required:"true"`
}

func (c *Config) GetPort() uint16 {
	return c.Port
}
