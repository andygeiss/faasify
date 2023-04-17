package server

import (
	"log"
	"net/http"
	"strings"
)

// WithAuthentication checks for a valid API token.
func WithAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check for a valid token.
		if parts := strings.Split(r.Header.Get("Authorization"), " "); len(parts) == 2 {
			if parts[1] == securityToken {
				next.ServeHTTP(w, r)
				return
			}
		}
		// Handle unauthorized access.
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}

// WithLogging logs the current request to the terminal.
func WithLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Do middleware specific work.
		log.Printf("[%-20s] executes [%-20s]", r.RemoteAddr, r.RequestURI)
		// Delegate to the next handler.
		next.ServeHTTP(w, r)
	}
}
