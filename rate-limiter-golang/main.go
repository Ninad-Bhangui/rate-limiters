package main

import (
	"log"
	"net/http"

	"rate-limiter-golang/ratelimiter"
)

func rateLimiterMiddleWare(limiter *ratelimiter.TokenBucketV2, next http.Handler) http.Handler {
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
	// limiter := ratelimiter.CreateTokenBucketV2(3, 1, func() { fmt.Println("internal function invoked") })
	// limiter.CallInternalFunc()
	// limiter.CallInternalFunc()
	//  fmt.Println("debug: Sleeping...")
	//  time.Sleep(2*time.Second)
	// limiter.CallInternalFunc()
	// limiter.CallInternalFunc()
	// limiter.CallInternalFunc()
	mux := http.NewServeMux()
	finalHandler := http.HandlerFunc(simpleApi)

	limiter := ratelimiter.CreateTokenBucketV2(3, 1)
	mux.Handle("/", rateLimiterMiddleWare(limiter, finalHandler))

	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
