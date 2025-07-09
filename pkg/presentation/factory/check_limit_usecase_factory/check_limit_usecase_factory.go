package check_limit_usecase_factory

import (
	"rate-limiter/pkg/domain/usecase/check_limit"
	"rate-limiter/pkg/infrastructure/cache"
	"rate-limiter/pkg/infrastructure/cache/redis"
	"rate-limiter/pkg/shared"
	"strings"
)

func NewCheckLimitUseCaseFactory(
	config shared.ConfigInterface,
) CheckLimitUseCaseFactoryInterface {
	return &CheckLimitUseCaseFactory{
		config: config,
	}
}

type CheckLimitUseCaseFactory struct {
	config shared.ConfigInterface
}

func (c *CheckLimitUseCaseFactory) Build() check_limit.CheckLimitUseCaseInterface {
	var engine cache.CacheInterface

	switch strings.ToUpper(c.config.GetCacheEngine()) {
	case "REDIS":
		engine = redis.NewRedisCache(c.config)
	case "MEMCACHED":
		panic("cache engine not implemented")
	default:
		panic("cache engine not implemented")
	}

	return check_limit.NewCheckLimitUseCase(engine)
}
