package redis

import "github.com/redis/go-redis/v9"

type RedisProviderInterface interface {
	GetRedisClient() (*redis.Client, error)
}
