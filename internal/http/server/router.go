// Code generated by fassify DO NOT EDIT
package server

import (
	"embed"
	"net/http"

	"github.com/andygeiss/faasify/internal/http/server/functions/hello"
	"github.com/andygeiss/faasify/internal/http/server/functions/index"
	"github.com/andygeiss/faasify/internal/http/server/functions/wasm_demo"
)

//go:embed bundle
var embedFS embed.FS

//go:generate go run ../../../cmd/update-functions/main.go

const Token = "wAJt1v8xpVU3olTH0VAf0SpYOa8Y0vXSqiN1JG+3BdQ="

func router() (mux *http.ServeMux) {
	// Init multiplexer
	mux = http.NewServeMux()

	// Add functions
	mux.HandleFunc("/hello", WithAuthentication(WithLogging(WithStatistics(hello.HandlerFunc()))))
	mux.HandleFunc("/wasm_demo", WithAuthentication(WithLogging(WithStatistics(wasm_demo.HandlerFunc()))))

	// Serve statistics
	mux.HandleFunc("/index", WithLogging(index.HandlerFunc(Token)))
	mux.HandleFunc("/stats", WithAuthentication(WithLogging(statsHandler())))

	// Serve embedded files
	mux.HandleFunc("/", WithEmbeddedFiles(embedFS, "bundle"))
	return
}
