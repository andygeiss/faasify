package server

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

type statistics struct {
	ActiveCount        map[string]int64 `json:"active_count"`
	TotalCount         map[string]int64 `json:"total_count"`
	LastResponseTimeMs map[string]int64 `json:"last_response_time_ms"`
	mutex              sync.Mutex
}

func (a *statistics) updatePreStats(name string) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.ActiveCount[name]++
}

func (a *statistics) updatePostStats(name string, start time.Time) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.ActiveCount[name]--
	a.LastResponseTimeMs[name] = time.Since(start).Milliseconds()
	a.TotalCount[name]++
}

// shared data
var stats = &statistics{
	ActiveCount:        make(map[string]int64),
	TotalCount:         make(map[string]int64),
	LastResponseTimeMs: make(map[string]int64),
}

func statsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stats.mutex.Lock()
		defer stats.mutex.Unlock()
		_ = json.NewEncoder(w).Encode(stats)
	}
}
