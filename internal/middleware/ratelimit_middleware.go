package middleware

import (
	"net/http"

	"github.com/tejasjadhav2024/ratelimiter-gateway/internal/ratelimiter"
)

func RateLimitMiddleware(limiter *ratelimiter.FixedWindowLimiter) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			clientID := r.Header.Get("X-API-Key")

			if !limiter.Allow(clientID) {
				w.WriteHeader(http.StatusTooManyRequests)
				w.Write([]byte("429 Too Many Requests"))
				return
			}

			next(w, r)
		}
	}
}
