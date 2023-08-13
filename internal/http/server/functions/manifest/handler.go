package manifest

import (
	"bytes"
	"compress/gzip"
	_ "embed"
	"encoding/json"
	"net/http"

	"github.com/andygeiss/faasify/internal/config"
)

//go:embed manifest.json
var manifest []byte

type icon struct {
	Purpose string `json:"purpose"`
	Sizes   string `json:"sizes"`
	Src     string `json:"src"`
	Type    string `json:"type"`
}

type response struct {
	BackgroundColor string `json:"background_color"`
	Display         string `json:"display"`
	Icons           []icon `json:"icons"`
	Name            string `json:"name"`
	Orientation     string `json:"orientation"`
	ShortName       string `json:"short_name"`
	StartUrl        string `json:"start_url"`
	ThemeColor      string `json:"theme_color"`
}

func HandlerFunc(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Encoding", "gzip")
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		gw, _ := gzip.NewWriterLevel(w, gzip.BestCompression)
		defer gw.Close()
		var res response
		json.NewDecoder(bytes.NewReader(manifest)).Decode(&res)
		res.Name = cfg.AppName
		res.ShortName = cfg.AppName
		json.NewEncoder(gw).Encode(res)
	}
}
