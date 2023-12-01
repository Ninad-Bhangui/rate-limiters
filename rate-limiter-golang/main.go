package main

import (
	"fmt"
	"time"
  "rate-limiter-golang/ratelimiter"
)


func main() {
	limiter := ratelimiter.CreateTokenBucketV2(3, 1, func() { fmt.Println("internal function invoked") })
	limiter.CallInternalFunc()
	limiter.CallInternalFunc()
  fmt.Println("debug: Sleeping...")
  time.Sleep(2*time.Second)
	limiter.CallInternalFunc()
	limiter.CallInternalFunc()
	limiter.CallInternalFunc()
}
