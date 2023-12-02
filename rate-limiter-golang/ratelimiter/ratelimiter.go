package ratelimiter

type RateLimiter interface {
  Request() bool
}
