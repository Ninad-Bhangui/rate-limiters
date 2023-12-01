package ratelimiter

import (
	"errors"
	"fmt"
	"time"
)

type TokenBucketV2 struct {
	tokenCount      int
	bucketSize      int
	refillRate      int
	lastElapsedTime time.Time
	internalFunc    func()
}

func (t *TokenBucketV2) CallInternalFunc() error {
	fmt.Println("debug: calling internal")
	t.fill()
	if t.revokeToken() == true {
    fmt.Println("debug: revoked token successfully: ", t.tokenCount)
		t.internalFunc()
		return nil
	} else {
		fmt.Println("debug: rate limit exceeded")
		return errors.New("Rate limit exceeded")
	}
}

func (t *TokenBucketV2) fill() {
	duration := time.Now().Sub(t.lastElapsedTime)
	fillCount := t.refillRate * int(duration/time.Second)
	t.addTokens(fillCount)
  t.lastElapsedTime = time.Now()
}

func (t *TokenBucketV2) addTokens(fillCount int) {
	if t.tokenCount < t.bucketSize {
		newBucketSize := t.tokenCount + fillCount
		if newBucketSize <= t.bucketSize {
			t.tokenCount = newBucketSize
		} else {
			t.tokenCount = t.bucketSize
		}
	}
}

func (t *TokenBucketV2) revokeToken() bool {
	fmt.Println("debug: current token count: ", t.tokenCount)
	if t.tokenCount > 0 {
		t.tokenCount -= 1
		return true
	}
	return false
}

func CreateTokenBucketV2(bucketSize int, refillRate int, internalFunc func()) *TokenBucketV2 {
	limiter := &TokenBucketV2{1, bucketSize, refillRate, time.Now(), internalFunc}
	return limiter
}
