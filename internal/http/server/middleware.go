package server

import (
	"embed"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/andygeiss/faasify/internal/config"
)

// WithAuthentication checks for a valid token
func WithAuthentication(cfg *config.Config, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if parts := strings.Split(r.Header.Get("Authorization"), " "); len(parts) == 2 {
			// Check for the admin token first
			if cfg.Token == parts[1] {
				w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
				next.ServeHTTP(w, r)
				return
			}
			// Check the account tokens
			id, secret, ok := r.BasicAuth()
			if ok && cfg.AccountAccess.VerifyAccount(id, secret) {
				w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
				next.ServeHTTP(w, r)
				return
			}
		}
		// Handle unauthorized access
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}

var mimetypes map[string]string = map[string]string{
	"avif": "image/avif",
	"css":  "text/css; charset=utf-8",
	"gif":  "image/gif",
	"htm":  "text/html; charset=utf-8",
	"html": "text/html; charset=utf-8",
	"jpeg": "image/jpeg",
	"jpg":  "image/jpeg",
	"js":   "text/javascript; charset=utf-8",
	"json": "application/json",
	"mjs":  "text/javascript; charset=utf-8",
	"pdf":  "application/pdf",
	"png":  "image/png",
	"svg":  "image/svg+xml",
	"wasm": "application/wasm",
	"webp": "image/webp",
	"xml":  "text/xml; charset=utf-8",
}

// WithEmbeddedFiles serves files from the embedded file system
func WithEmbeddedFiles(efs embed.FS, prefix string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the suffix
		suffix := r.RequestURI
		if suffix == "/" {
			suffix = "index.html"
		}
		// Read the file content
		path := filepath.Join(prefix, suffix)
		content, err := efs.ReadFile(path)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		// Add mime type
		mimeType := "text/plain"
		parts := strings.Split(suffix, ".")
		for ending, mt := range mimetypes {
			if ending == parts[len(parts)-1] {
				mimeType = mt
			}
		}
		w.Header().Add("Content-Type", mimeType)
		// Compress file content
		w.Header().Set("Content-Encoding", "gzip")
		w.Write(content)
	}
}

// WithLogging logs the current request to the terminal
func WithLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Do middleware specific work
		log.Printf("[%-20s] requests [%-20s]", r.RemoteAddr, r.RequestURI)
		// Delegate to the next handler
		next.ServeHTTP(w, r)
	}
}
