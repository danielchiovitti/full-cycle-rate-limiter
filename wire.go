//go:build wireinject
// +build wireinject

package full_cycle_rate_limiter

import (
	"github.com/google/wire"
	"rate-limiter/pkg/presentation"
	"rate-limiter/pkg/presentation/factory/check_limit_usecase_factory"
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
	check_limit_usecase_factory.NewCheckLimitUseCaseFactory,
)

func InitializeLoader() *presentation.Loader {
	wire.Build(superSet)
	return &presentation.Loader{}
}
