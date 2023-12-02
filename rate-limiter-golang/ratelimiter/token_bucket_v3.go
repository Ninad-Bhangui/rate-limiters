package ratelimiter

import (
	"sync"
	"time"
)

type TokenBucketV3 struct {
	tokenCount      int
	bucketSize      int
	refillRate      int
	lastElapsedTime time.Time
	mu              sync.RWMutex
}

func (t *TokenBucketV3) Request() bool {
	t.fill()
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.revokeToken()
}

func (t *TokenBucketV3) fill() {
	duration := time.Now().Sub(t.lastElapsedTime)
	fillCount := t.refillRate * int(duration/time.Second)
	t.mu.Lock()
	defer t.mu.Unlock()
	t.addTokens(int(fillCount))
	t.lastElapsedTime = time.Now()
}

func (t *TokenBucketV3) addTokens(fillCount int) {
	t.tokenCount = min(t.bucketSize, t.tokenCount+fillCount)
}

func (t *TokenBucketV3) revokeToken() bool {
	// fmt.Println("debug: current token count: ", t.tokenCount)
	if t.tokenCount > 0 {
		t.tokenCount -= 1
		return true
	}
	return false
}

func CreateTokenBucketV3(bucketSize int, refillRate int) *TokenBucketV3 {
	limiter := &TokenBucketV3{
		tokenCount:      bucketSize,
		bucketSize:      bucketSize,
		refillRate:      refillRate,
		lastElapsedTime: time.Now(),
	} // Initialzing bucket full of tokens
	return limiter
}
