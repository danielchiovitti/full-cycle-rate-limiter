package redis

import (
	"rate-limiter/pkg/domain/model"
	"rate-limiter/pkg/shared"
	"strconv"
	"strings"
	"sync"
)

var redisCacheLock sync.Mutex
var redisCacheInstance *RedisCache

func NewRedisCache(
	config shared.ConfigInterface,
) *RedisCache {
	if redisCacheInstance == nil {
		redisCacheLock.Lock()
		defer redisCacheLock.Unlock()
		if redisCacheInstance == nil {
			list := strings.Split(config.GetConstraintList(), ";")
			cList := make([]model.Constraint, len(list))
			for _, v := range list {
				innerValue := strings.Split(v, ",")
				requests, _ := strconv.Atoi(innerValue[2])
				blockTime, _ := strconv.Atoi(innerValue[3])
				cList = append(cList, model.Constraint{
					Key:       innerValue[1],
					KeyType:   model.ConstraintType(innerValue[0]),
					Requests:  requests,
					BlockTime: blockTime,
				})
			}

			redisCacheInstance = &RedisCache{
				config: config,
			}
		}
	}

	return redisCacheInstance
}

type RedisCache struct {
	config         shared.ConfigInterface
	constraintList []model.Constraint
}

func (r *RedisCache) SetValue(key, value string, ttl int) error {
	return nil
}

func (r *RedisCache) GetValue(key string) (string, error) {
	return "", nil
}
