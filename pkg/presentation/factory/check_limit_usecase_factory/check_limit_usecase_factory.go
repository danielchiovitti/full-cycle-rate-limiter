package check_limit_usecase_factory

import (
	"rate-limiter/pkg/domain/usecase/check_limit"
	"rate-limiter/pkg/infrastructure/cache"
	"rate-limiter/pkg/infrastructure/cache/redis"
	redis2 "rate-limiter/pkg/infrastructure/database/redis"
	"rate-limiter/pkg/shared"
	"strings"
)

func NewCheckLimitUseCaseFactory(
	config shared.ConfigInterface,
	redisProvider redis2.RedisProviderInterface,
) CheckLimitUseCaseFactoryInterface {
	return &CheckLimitUseCaseFactory{
		config:        config,
		redisProvider: redisProvider,
	}
}

type CheckLimitUseCaseFactory struct {
	config        shared.ConfigInterface
	redisProvider redis2.RedisProviderInterface
}

func (c *CheckLimitUseCaseFactory) Build() check_limit.CheckLimitUseCaseInterface {
	var engine cache.CacheInterface

	switch strings.ToUpper(c.config.GetCacheEngine()) {
	case "REDIS":
		engine = redis.NewRedisCache(c.config, c.redisProvider)
	case "MEMCACHED":
		panic("cache engine not implemented")
	default:
		panic("cache engine not implemented")
	}

	return check_limit.NewCheckLimitUseCase(engine, c.config)
}
