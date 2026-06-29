package middleware

import (
	"net/http"
)

const validAPIKey = "secret-key-123"

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")

		if apiKey != validAPIKey {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("401 Unauthorized: invalid or missing API key"))
			return
		}

		next(w, r)
	}
}
