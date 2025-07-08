package get_value_usecase_factory

import (
	"rate-limiter/pkg/domain/usecase/get_value"
	"rate-limiter/pkg/infrastructure/cache"
	"rate-limiter/pkg/infrastructure/cache/redis"
	"rate-limiter/pkg/shared"
	"strings"
)

func NewGetValueUseCaseFactory(
	config shared.ConfigInterface,
) *GetValueUseCaseFactory {
	return &GetValueUseCaseFactory{
		config: config,
	}
}

type GetValueUseCaseFactory struct {
	config shared.ConfigInterface
}

func (g *GetValueUseCaseFactory) Build() *get_value.GetValueUseCase {
	var engine cache.CacheInterface

	switch strings.ToUpper(g.config.GetCacheEngine()) {
	case "REDIS":
		engine = redis.NewRedisCache(g.config)
	case "MEMCACHED":
		panic("Engine not implemented")
	default:
		panic("Engine not implemented")
	}

	return get_value.NewGetValueUseCase(engine)
}
