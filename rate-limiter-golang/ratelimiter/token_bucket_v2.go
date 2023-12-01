package ratelimiter

import (
	"fmt"
	"time"
)

type TokenBucketV2 struct {
	tokenCount      int
	bucketSize      int
	refillRate      float32 
	lastElapsedTime time.Time
	// internalFunc    func()
}

func (t *TokenBucketV2) Request() bool {
	t.fill()
	return t.revokeToken()
}

// func (t *TokenBucketV2) CallInternalFunc() error {
// 	fmt.Println("debug: calling internal")
// 	t.fill()
// 	if t.revokeToken() == true {
//     fmt.Println("debug: revoked token successfully: ", t.tokenCount)
// 		t.internalFunc()
// 		return nil
// 	} else {
// 		fmt.Println("debug: rate limit exceeded")
// 		return errors.New("Rate limit exceeded")
// 	}
// }

func (t *TokenBucketV2) fill() {
	duration := time.Now().Sub(t.lastElapsedTime)
	fillCount := t.refillRate * float32(duration/time.Second)
	t.addTokens(int(fillCount))
	t.lastElapsedTime = time.Now()
}

func (t *TokenBucketV2) addTokens(fillCount int) {
	t.tokenCount = min(t.bucketSize, t.tokenCount+fillCount)
}

func (t *TokenBucketV2) revokeToken() bool {
	fmt.Println("debug: current token count: ", t.tokenCount)
	if t.tokenCount > 0 {
		t.tokenCount -= 1
		return true
	}
	return false
}

func CreateTokenBucketV2(bucketSize int, refillRate float32) *TokenBucketV2 {
	limiter := &TokenBucketV2{10, bucketSize, refillRate, time.Now()}
	return limiter
}
