package redis

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"rate-limiter/pkg/shared"
	"sync"
)

var lockRedisProvider sync.Mutex
var lockRedisClient sync.Mutex
var redisProviderInstance *RedisProvider
var redisClientInstance *redis.Client

type RedisProvider struct {
	config shared.ConfigInterface
}

func NewRedisProvider(
	config shared.ConfigInterface,
) RedisProviderInterface {
	if redisProviderInstance == nil {
		lockRedisProvider.Lock()
		defer lockRedisProvider.Unlock()
		if redisProviderInstance == nil {
			redisProviderInstance = &RedisProvider{
				config: config,
			}
		}
	}
	return redisProviderInstance
}

func (p *RedisProvider) GetRedisClient() (*redis.Client, error) {
	if redisClientInstance == nil {
		lockRedisClient.Lock()
		defer lockRedisClient.Unlock()
		if redisClientInstance == nil {
			redisClientInstance = redis.NewClient(&redis.Options{
				Addr:     fmt.Sprintf("%s:%d", p.config.GetRedisHost(), p.config.GetRedisPort()),
				Password: p.config.GetRedisPassword(),
				DB:       0,
				Protocol: 2,
				PoolSize: 250,
			})
		}
	}

	return redisClientInstance, nil
}
