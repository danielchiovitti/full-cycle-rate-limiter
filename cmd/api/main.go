package main

import (
	"fmt"
	"net/http"
	full_cycle_rate_limiter "rate-limiter"
	"time"
)

func main() {
	time.Sleep(1 * time.Second)

	l := full_cycle_rate_limiter.InitializeLoader()
	r := l.GetRoutes()
	c := l.GetConfig()

	err := http.ListenAndServe(fmt.Sprintf(":%d", c.GetPort()), r)
	if err != nil {
		panic(err)
	}
}
