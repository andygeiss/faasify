package hello

import (
	"encoding/json"
	"net/http"

	"github.com/andygeiss/faasify/internal/config"
)

type response struct {
	Data  string `json:"data"`
	Error error  `json:"error,omitempty"`
}

func HandlerFunc(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(response{Data: "Hello World!"})
	}
}
