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
	Port           uint16 `envconfig:"PORT" required:"true"`
	MaxRequests    int    `envconfig:"MAX_REQUESTS" required:"true"`
	BlockTime      int    `envconfig:"BLOCK_TIME" required:"true"`
	ConstraintList string `envconfig:"CONSTRAINT_LIST" required:"true"`
	CacheEngine    string `envconfig:"CACHE_ENGINE" required:"true"`
}

func (c *Config) GetPort() uint16 {
	return c.Port
}

func (c *Config) GetMaxRequests() int {
	return c.MaxRequests
}

func (c *Config) GetBlockTime() int {
	return c.BlockTime
}

func (c *Config) GetConstraintList() string {
	return c.ConstraintList
}

func (c *Config) GetCacheEngine() string {
	return c.CacheEngine
}
