// Code generated by fassify DO NOT EDIT
package server

import (
	"embed"
	"net/http"

	"github.com/andygeiss/faasify/internal/config"
	"github.com/andygeiss/faasify/internal/http/server/functions/app"
	"github.com/andygeiss/faasify/internal/http/server/functions/index"
	"github.com/andygeiss/faasify/internal/http/server/functions/manifest"
	"github.com/andygeiss/faasify/internal/http/server/functions/secure"
)

//go:embed bundle
var embedFS embed.FS

//go:generate go run ../../../cmd/update-functions/main.go

func router(cfg *config.Config) (mux *http.ServeMux) {
	// Init multiplexer
	mux = http.NewServeMux()

	// Set generated security token
	cfg.Token = "mkM6MLBo11KLb7q4aZapfuqdhv/L70mWJJuDk5AmslY="

	// Add functions
	mux.HandleFunc("/app", WithLogging(app.HandlerFunc(cfg)))
	mux.HandleFunc("/index", WithLogging(index.HandlerFunc(cfg)))
	mux.HandleFunc("/manifest", WithLogging(manifest.HandlerFunc(cfg)))
	mux.HandleFunc("/secure", WithAuthentication(cfg, WithLogging(secure.HandlerFunc(cfg))))

	// Serve embedded files
	mux.HandleFunc("/", WithEmbeddedFiles(embedFS, "bundle"))
	return
}
