//go:build wireinject
// +build wireinject

package full_cycle_rate_limiter

import (
	"github.com/google/wire"
	"rate-limiter/pkg/infrastructure/cache/redis"
	"rate-limiter/pkg/presentation"
	"rate-limiter/pkg/presentation/middleware"
	"rate-limiter/pkg/presentation/route"
	"rate-limiter/pkg/shared"
)

var superSet = wire.NewSet(
	presentation.NewLoader,
	shared.NewConfig,
	route.NewHealthRoute,
	route.NewGameRoute,
	middleware.NewRatingMiddleware,
	redis.NewRedisCache,
)

func InitializeLoader() *presentation.Loader {
	wire.Build(superSet)
	return &presentation.Loader{}
}
