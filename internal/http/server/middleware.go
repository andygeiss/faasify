package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

// WithAuthentication checks for a valid API token
func WithAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check for a valid token
		if parts := strings.Split(r.Header.Get("Authorization"), " "); len(parts) == 2 {
			if parts[1] == securityToken {
				next.ServeHTTP(w, r)
				return
			}
		}
		// Handle unauthorized access
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}

// WithLogging logs the current request to the terminal
func WithLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Do middleware specific work
		log.Printf("[%-20s] executes [%-20s]", r.RemoteAddr, r.RequestURI)
		// Delegate to the next handler
		next.ServeHTTP(w, r)
	}
}

type statistics struct {
	counter      map[string]uint64
	responseTime map[string]time.Duration
	mutex        sync.Mutex
}

// WithStatistics collects data about the function calls
func WithStatistics(stats *statistics, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stats.mutex.Lock()
		defer stats.mutex.Unlock()
		start := time.Now()
		// Delegate to the next handler
		next.ServeHTTP(w, r)
		// Create statistics
		name := strings.TrimPrefix("/function/", r.RequestURI)
		stats.counter[name]++
		stats.responseTime[name] = time.Since(start)
	}
}

func statsHandler(stats *statistics) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stats.mutex.Lock()
		defer stats.mutex.Unlock()
		_ = json.NewEncoder(w).Encode(stats)
	}
}
