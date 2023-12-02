package ratelimiter

import (
	"time"
)

type TokenBucketV2 struct {
	tokenCount      int
	bucketSize      int
	refillRate      int 
	lastElapsedTime time.Time
}

func (t *TokenBucketV2) Request() bool {
	t.fill()
	return t.revokeToken()
}

func (t *TokenBucketV2) fill() {
	duration := time.Now().Sub(t.lastElapsedTime)
	fillCount := t.refillRate * int(duration/time.Second)
	t.addTokens(int(fillCount))
	t.lastElapsedTime = time.Now()
}

func (t *TokenBucketV2) addTokens(fillCount int) {
	t.tokenCount = min(t.bucketSize, t.tokenCount+fillCount)
}

func (t *TokenBucketV2) revokeToken() bool {
	// fmt.Println("debug: current token count: ", t.tokenCount)
	if t.tokenCount > 0 {
		t.tokenCount -= 1
		return true
	}
	return false
}

func CreateTokenBucketV2(bucketSize int, refillRate int) *TokenBucketV2 {
	limiter := &TokenBucketV2{10, bucketSize, refillRate, time.Now()}
	return limiter
}
