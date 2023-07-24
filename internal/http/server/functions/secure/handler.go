package secure

import (
	"compress/gzip"
	_ "embed"
	"encoding/json"
	"net/http"

	"github.com/andygeiss/faasify/internal/config"
)

type response struct {
	Data string `json:"data"`
}

func HandlerFunc(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Encoding", "gzip")
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		gw, _ := gzip.NewWriterLevel(w, gzip.BestCompression)
		defer gw.Close()
		json.NewEncoder(gw).Encode(response{Data: "secure"})
	}
}
