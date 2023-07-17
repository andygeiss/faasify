// Code generated by fassify DO NOT EDIT
package server

import (
	"embed"
	"net/http"

	"github.com/andygeiss/faasify/internal/http/server/functions/content"
	"github.com/andygeiss/faasify/internal/http/server/functions/index"
	"github.com/andygeiss/faasify/internal/http/server/functions/wasm_demo"
)

//go:embed bundle
var embedFS embed.FS

//go:generate go run ../../../cmd/update-functions/main.go

const Token = "UdnQBiJRd0UkaXKgX0eBwZY3YL/W446zD/zDn/BiLyE="

func router() (mux *http.ServeMux) {
	// Init multiplexer
	mux = http.NewServeMux()

	// Add functions
	mux.HandleFunc("/content", WithAuthentication(WithLogging(WithStatistics(content.HandlerFunc(Token)))))
	mux.HandleFunc("/wasm_demo", WithAuthentication(WithLogging(WithStatistics(wasm_demo.HandlerFunc(Token)))))

	// Serve statistics
	mux.HandleFunc("/index", WithLogging(index.HandlerFunc(Token)))
	mux.HandleFunc("/stats", WithAuthentication(WithLogging(statsHandler())))

	// Serve embedded files
	mux.HandleFunc("/", WithEmbeddedFiles(embedFS, "bundle"))
	return
}
