package redis

import "sync"

var redisCacheLock sync.Mutex
var redisCacheInstance *RedisCache

func NewRedisCache() *RedisCache {
	if redisCacheInstance == nil {
		redisCacheLock.Lock()
		defer redisCacheLock.Unlock()
		if redisCacheInstance == nil {
			redisCacheInstance = &RedisCache{}
		}
	}

	return redisCacheInstance
}

type RedisCache struct {
}

func (r *RedisCache) SetValue(key, value string, ttl int) error {
	return nil
}

func (r *RedisCache) GetValue(key string) (string, error) {
	return "", nil
}
