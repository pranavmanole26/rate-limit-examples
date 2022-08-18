package main

import (
	"fmt"
	"net/http"
	"time"

	// redisrate "rate-limit-examples/redis-rate"
	"rate-limit-examples/russell_luo"
	// redis_ratelimiter "rate-limit-examples/go-redis-ratelimiter"

	"github.com/gorilla/mux"
)

func main() {

	// 1. START: example with go mux
	router := mux.NewRouter()
	router.Path("/hello").Methods(http.MethodGet).HandlerFunc(func(resW http.ResponseWriter, req *http.Request) {
		ok := russell_luo.AllowRequest(req.URL.Path)
		if ok {
			resW.Write([]byte("Hello world"))
		} else {
			resW.Write([]byte("Too many requests"))
		}
	})

	server := http.Server{
		Addr:         "localhost:82",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      router,
	}

	fmt.Printf(server.ListenAndServe().Error())

	// 1. END: example with go mux
}
