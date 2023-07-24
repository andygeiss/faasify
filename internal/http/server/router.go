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

var Domain string

const Token = "0p55AY366wdHOxpj6zM7q8ZWei97JbYCuufLy977iEc="

var Url string

func router() (mux *http.ServeMux) {
	// Init multiplexer
	mux = http.NewServeMux()

	// Add functions
	mux.HandleFunc("/content", WithAuthentication(WithLogging(WithStatistics(content.HandlerFunc(Token, Domain, Url)))))
	mux.HandleFunc("/wasm_demo", WithAuthentication(WithLogging(WithStatistics(wasm_demo.HandlerFunc(Token, Domain, Url)))))

	// Serve statistics
	mux.HandleFunc("/index", WithLogging(index.HandlerFunc(Token)))
	mux.HandleFunc("/stats", WithAuthentication(WithLogging(statsHandler())))

	// Serve embedded files
	mux.HandleFunc("/", WithEmbeddedFiles(embedFS, "bundle"))
	return
}
