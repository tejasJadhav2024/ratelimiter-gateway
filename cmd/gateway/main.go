package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/tejasjadhav2024/ratelimiter-gateway/internal/cache"
	"github.com/tejasjadhav2024/ratelimiter-gateway/internal/middleware"
	"github.com/tejasjadhav2024/ratelimiter-gateway/internal/ratelimiter"
)

func main() {
	err := cache.InitRedis("localhost:6379")
	if err != nil {
		log.Fatal(err)
	}

	limiter := ratelimiter.NewRedisFixedWindowLimiter(
		cache.Client,
		5,
		60*time.Second,
	)

	handler := middleware.AuthMiddleware(
		middleware.RateLimitMiddleware(limiter)(pingHandler),
	)

	http.HandleFunc("/ping", handler)

	log.Println("Server starting on port 8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}
