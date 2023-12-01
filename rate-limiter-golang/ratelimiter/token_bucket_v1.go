package ratelimiter

import (
	"errors"
	"fmt"
	"time"
)

type TokenBucket struct {
	tokenCount   int
	bucketSize   int
	refillRate   int
	internalFunc func()
}

func (t *TokenBucket) CallInternalFunc() error {
	fmt.Println("debug: calling internal")
	if t.revokeToken() == true {
		fmt.Println("debug: revoked token successfully")
		t.internalFunc()
		return nil
	} else {
    fmt.Println("debug: rate limit exceeded")
		return errors.New("Rate limit exceeded")
	}
}
func (t *TokenBucket) StartFill() {
	for range time.Tick(time.Second * 1) {
		t.addToken()
	}
}

func (t *TokenBucket) addToken() {
	if t.tokenCount < t.bucketSize {
		newBucketSize := t.tokenCount + t.refillRate
		if newBucketSize <= t.bucketSize {
			t.tokenCount = newBucketSize
		} else {
			t.tokenCount = t.bucketSize
		}
	}
}

func (t *TokenBucket) revokeToken() bool {
  fmt.Println("debug: current token count: ", t.tokenCount)
	if t.tokenCount > 0 {
		t.tokenCount -= 1
		return true
	}
	return false
}

func CreateTokenBucketV1(bucketSize int, refillRate int, internalFunc func()) *TokenBucket {
	limiter := &TokenBucket{1, bucketSize, refillRate, internalFunc}
	go limiter.StartFill()
	return limiter
}

