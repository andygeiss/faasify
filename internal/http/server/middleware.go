package server

import (
	"crypto/sha256"
	"crypto/subtle"
	"embed"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

// WithAuthentication checks for a valid API token
func WithAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check for a valid token
		if parts := strings.Split(r.Header.Get("Authorization"), " "); len(parts) == 2 {
			// Token-based Authentication
			if parts[1] == Token {
				w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
				next.ServeHTTP(w, r)
				return
			}
			// Basic Authentication (we still use the Token as password)
			if user, pass, ok := r.BasicAuth(); ok {
				userHash := sha256.Sum256([]byte(user))
				passHash := sha256.Sum256([]byte(pass))
				expectedUserHash := sha256.Sum256([]byte("faasify"))
				expectedPassHash := sha256.Sum256([]byte(Token))
				userMatch := (subtle.ConstantTimeCompare(userHash[:], expectedUserHash[:]) == 1)
				passMatch := (subtle.ConstantTimeCompare(passHash[:], expectedPassHash[:]) == 1)
				if userMatch && passMatch {
					w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
					next.ServeHTTP(w, r)
					return
				}
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
		//gz := gzip.NewWriter(w)
		//defer gz.Close()
		//gz.Write(content)
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

func WithStatistics(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.RequestURI
		start := time.Now()
		// Add pre function statistics
		stats.updatePreStats(name)
		// Delegate to the next handler
		next.ServeHTTP(w, r)
		// Add post function statistics
		stats.updatePostStats(name, start)
	}
}
