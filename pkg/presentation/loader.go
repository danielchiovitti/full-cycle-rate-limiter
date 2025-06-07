package presentation

import (
	"github.com/go-chi/chi/v5"
	"rate-limiter/pkg/presentation/route"
	"rate-limiter/pkg/shared"
)

func NewLoader(
	config shared.ConfigInterface,
	healthRoute *route.HealthRoute,
	gameRoute *route.GameRoute,
) *Loader {
	return &Loader{
		config:      config,
		healthRoute: healthRoute,
		gameRoute:   gameRoute,
	}
}

type Loader struct {
	config      shared.ConfigInterface
	healthRoute *route.HealthRoute
	gameRoute   *route.GameRoute
}

func (l *Loader) GetRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Mount("/health", l.healthRoute.HealthRoutes())
	r.Mount("/game", l.gameRoute.GameRoutes())
	return r
}

func (l *Loader) GetConfig() shared.ConfigInterface {
	return l.config
}
