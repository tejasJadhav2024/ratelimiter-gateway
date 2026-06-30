package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tejasjadhav2024/ratelimiter-gateway/internal/middleware"
)

func main() {
	http.HandleFunc("/ping", middleware.AuthMiddleware(pingHandler))

	log.Println("Server starting on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}
