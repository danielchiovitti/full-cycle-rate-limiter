package route

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	errors "rate-limiter/pkg/shared/error"
	"rate-limiter/pkg/shared/helpers"
	"sync"
)

var healthRouteLock sync.Mutex
var healthRouteInstance *HealthRoute

func NewHealthRoute() *HealthRoute {
	if healthRouteInstance == nil {
		healthRouteLock.Lock()
		defer healthRouteLock.Unlock()
		if healthRouteInstance == nil {
			healthRouteInstance = &HealthRoute{}
		}
	}

	return healthRouteInstance
}

type HealthRoute struct{}

func (HealthRoute) HealthRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", getHealth)
	return r
}

func getHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res := map[string]string{
		"status": "ok",
	}

	resJson, err := json.Marshal(res)
	if err != nil {
		helpers.JsonResponse(w, http.StatusBadRequest, errors.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	_, err = w.Write(resJson)
	if err != nil {
		helpers.JsonResponse(w, http.StatusBadRequest, errors.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
}
