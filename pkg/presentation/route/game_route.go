package route

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"rate-limiter/pkg/presentation/middleware"
	"rate-limiter/pkg/shared/helpers"
)

func NewGameRoute(
	ratingMiddleware *middleware.RatingMiddleware,
) *GameRoute {
	return &GameRoute{
		ratingMiddleware: ratingMiddleware,
	}
}

type GameRoute struct {
	ratingMiddleware *middleware.RatingMiddleware
}

func (g *GameRoute) GameRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.With(g.ratingMiddleware.ServeRating).Get("/", getGame)
	return r
}

func getGame(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{
		"game-01": "running",
		"game-02": "running",
		"game-03": "running",
	}

	helpers.JsonResponse(w, http.StatusOK, res)
}
