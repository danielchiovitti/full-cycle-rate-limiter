// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package full_cycle_rate_limiter

import (
	"github.com/google/wire"
	"rate-limiter/pkg/infrastructure/database/redis"
	"rate-limiter/pkg/presentation"
	"rate-limiter/pkg/presentation/factory/check_limit_usecase_factory"
	"rate-limiter/pkg/presentation/middleware"
	"rate-limiter/pkg/presentation/route"
	"rate-limiter/pkg/shared"
)

// Injectors from wire.go:

func InitializeLoader() *presentation.Loader {
	configInterface := shared.NewConfig()
	healthRoute := route.NewHealthRoute()
	redisProviderInterface := redis.NewRedisProvider(configInterface)
	checkLimitUseCaseFactoryInterface := check_limit_usecase_factory.NewCheckLimitUseCaseFactory(configInterface, redisProviderInterface)
	ratingMiddleware := middleware.NewRatingMiddleware(checkLimitUseCaseFactoryInterface)
	gameRoute := route.NewGameRoute(ratingMiddleware)
	loader := presentation.NewLoader(configInterface, healthRoute, gameRoute)
	return loader
}

// wire.go:

var superSet = wire.NewSet(presentation.NewLoader, shared.NewConfig, route.NewHealthRoute, route.NewGameRoute, middleware.NewRatingMiddleware, check_limit_usecase_factory.NewCheckLimitUseCaseFactory, redis.NewRedisProvider)
