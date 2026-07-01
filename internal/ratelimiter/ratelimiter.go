package ratelimiter

type RateLimiter interface {
	Allow(clientID string) bool
}
