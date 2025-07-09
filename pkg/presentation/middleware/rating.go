package middleware

import (
	"net/http"
	"rate-limiter/pkg/presentation/factory/check_limit_usecase_factory"
	"rate-limiter/pkg/shared/helpers"
	"sync"
)

var ratingMiddlewareInstance *RatingMiddleware
var lockRating sync.Mutex

func NewRatingMiddleware(
	checkLimitUseCaseFactory check_limit_usecase_factory.CheckLimitUseCaseFactoryInterface,
) *RatingMiddleware {
	if ratingMiddlewareInstance == nil {
		lockRating.Lock()
		defer lockRating.Unlock()
		if ratingMiddlewareInstance == nil {
			ratingMiddlewareInstance = &RatingMiddleware{
				checkLimitUseCaseFactory: checkLimitUseCaseFactory,
			}
		}
	}
	return ratingMiddlewareInstance
}

type RatingMiddleware struct {
	checkLimitUseCaseFactory check_limit_usecase_factory.CheckLimitUseCaseFactoryInterface
}

func (rm *RatingMiddleware) ServeRating(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		useCase := rm.checkLimitUseCaseFactory.Build()
		ok, err := useCase.Execute("", "")
		if err != nil {
			helpers.JsonResponse(w, http.StatusInternalServerError, "")
		}

		if !ok {
			helpers.JsonResponse(w, http.StatusTooManyRequests, "you have reached the maximum number of requests or actions allowed within a certain time frame")
		}

		//ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		//fmt.Println(ip)
		//fmt.Println(r.RemoteAddr)
		//fmt.Println(r.Header.Get("X-Real-IP"))

		next.ServeHTTP(w, r)
	})
}
