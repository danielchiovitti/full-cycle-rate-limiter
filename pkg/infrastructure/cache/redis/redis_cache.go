package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"rate-limiter/pkg/domain/model"
	redis2 "rate-limiter/pkg/infrastructure/database/redis"
	"rate-limiter/pkg/shared"
	"strconv"
	"strings"
	"sync"
	"time"
)

var redisCacheLock sync.Mutex
var redisCacheInstance *RedisCache

func NewRedisCache(
	config shared.ConfigInterface,
	redisProvider redis2.RedisProviderInterface,
) *RedisCache {
	if redisCacheInstance == nil {
		redisCacheLock.Lock()
		defer redisCacheLock.Unlock()
		if redisCacheInstance == nil {
			list := strings.Split(config.GetConstraintList(), ";")
			var cList []model.Constraint
			for _, v := range list {
				innerValue := strings.Split(v, ",")
				requests, _ := strconv.Atoi(innerValue[2])
				blockTime, _ := strconv.Atoi(innerValue[3])
				cList = append(cList, model.Constraint{
					Key:       innerValue[1],
					KeyType:   model.ConstraintType(innerValue[0]),
					Requests:  int64(requests),
					BlockTime: int64(blockTime),
				})
			}

			rdb, _ := redisProvider.GetRedisClient()
			redisCacheInstance = &RedisCache{
				config:         config,
				constraintList: cList,
				rdb:            rdb,
			}
			redisCacheInstance.preload()
		}
	}

	return redisCacheInstance
}

type RedisCache struct {
	config         shared.ConfigInterface
	constraintList []model.Constraint
	rdb            *redis.Client
}

func (r *RedisCache) SetValue(ctx context.Context, keyType model.ConstraintType, key string, value interface{}, ttl int) error {
	builtKey := r.buildKey(keyType, key)
	err := r.rdb.Set(ctx, builtKey, value, time.Duration(ttl)*time.Second).Err()
	return err
}

func (r *RedisCache) IncrValue(ctx context.Context, keyType model.ConstraintType, key string) (int64, error) {
	builtKey := r.buildKey(keyType, key)
	val, err := r.rdb.Incr(ctx, builtKey).Result()
	if err != nil {
		return 0, nil
	}

	exists, err := r.rdb.TTL(ctx, builtKey).Result()
	if err != nil {
		return 0, err
	}

	if exists < 0 {
		r.rdb.Expire(ctx, builtKey, time.Duration(r.config.GetTimeRange())*time.Second)
	}

	return val, nil
}

func (r *RedisCache) GetValue(ctx context.Context, keyType model.ConstraintType, key string) (interface{}, error) {
	builtKey := r.buildKey(keyType, key)
	val, err := r.rdb.Get(ctx, builtKey).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return val, nil
}

func (r *RedisCache) buildKey(keyType model.ConstraintType, key string) string {
	return fmt.Sprintf("%s_%s", string(keyType), key)
}

func (r *RedisCache) preload() {
	for _, v := range r.constraintList {
		val, err := json.Marshal(v)
		if err != nil {
			panic("error marshalling preload values")
		}

		ctx := context.Background()

		var keyType model.ConstraintType
		if v.KeyType == "CONSTRAINT_LI" {
			keyType = model.CONSTRAINT_LI
		} else if v.KeyType == "CONSTRAINT_LT" {
			keyType = model.CONSTRAINT_LT
		}

		err = r.SetValue(ctx, keyType, v.Key, val, 0)
		if err != nil {
			panic("error inserting preload values")
		}
	}
}
