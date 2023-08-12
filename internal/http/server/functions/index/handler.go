package index

import (
	"compress/gzip"
	_ "embed"
	"html/template"
	"net/http"

	"github.com/andygeiss/faasify/internal/config"
)

//go:embed html.tmpl
var html string

//go:embed styles.css
var styles string

type response struct {
	AppName string
	Styles  template.CSS
	Token   string
}

func HandlerFunc(cfg *config.Config) http.HandlerFunc {
	t, _ := template.New("t").Parse(html)
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		w.Header().Add("Content-Encoding", "gzip")
		gw, _ := gzip.NewWriterLevel(w, gzip.BestCompression)
		defer gw.Close()
		t.Execute(gw, response{cfg.AppName, template.CSS(styles), cfg.Token})
	}
}
