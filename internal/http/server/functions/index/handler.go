package index

import (
	"compress/gzip"
	_ "embed"
	"encoding/hex"
	"html/template"
	"net/http"

	"github.com/andygeiss/faasify/internal/config"
)

//go:embed html.tmpl
var html string

//go:embed styles.css
var styles string

type response struct {
	Id     string
	Styles template.CSS
	Token  string
}

func HandlerFunc(cfg *config.Config) http.HandlerFunc {
	t, _ := template.New("t").Parse(html)
	return func(w http.ResponseWriter, r *http.Request) {
		// Ensure that the response will be compressed
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		w.Header().Add("Content-Encoding", "gzip")
		gw, _ := gzip.NewWriterLevel(w, gzip.BestCompression)
		defer gw.Close()
		// Verify account
		id := r.FormValue("femail")
		secret := r.FormValue("fpassword")
		if cfg.AccountAccess.VerifyAccount(id, secret) {
			hash := cfg.AccountAccess.GetAccount(id).Hash
			t.Execute(gw, response{id, template.CSS(styles), hex.EncodeToString(hash)})
			return
		}
		t.Execute(gw, response{"", template.CSS(styles), ""})
	}
}
