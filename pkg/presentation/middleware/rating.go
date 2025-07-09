package middleware

import (
	"fmt"
	"net"
	"net/http"
	"sync"
)

var ratingMiddlewareInstance *RatingMiddleware
var lockRating sync.Mutex

func NewRatingMiddleware() *RatingMiddleware {
	if ratingMiddlewareInstance == nil {
		lockRating.Lock()
		defer lockRating.Unlock()
		if ratingMiddlewareInstance == nil {
			ratingMiddlewareInstance = &RatingMiddleware{}
		}
	}
	return ratingMiddlewareInstance
}

type RatingMiddleware struct {
}

func (rm *RatingMiddleware) ServeRating(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		fmt.Println(ip)
		fmt.Println(r.RemoteAddr)
		fmt.Println(r.Header.Get("X-Real-IP"))

		next.ServeHTTP(w, r)
	})
}
