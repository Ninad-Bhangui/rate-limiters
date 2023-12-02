package main

import (
	"log"
	"net/http"

	"rate-limiter-golang/ratelimiter"
)

func rateLimiterMiddleWare(limiter ratelimiter.RateLimiter, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if limiter.Request() {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("429 - Too many requests!!"))
		}
	})
}

func simpleApi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Yo!"))
}

func main() {
	mux := http.NewServeMux()
	finalHandler := http.HandlerFunc(simpleApi)

	limiter := ratelimiter.CreateTokenBucketV3(100, 1)
	mux.Handle("/", rateLimiterMiddleWare(limiter, finalHandler))

	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
