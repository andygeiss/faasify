package private

import (
	_ "embed"
	"html/template"
	"net/http"
	"strings"

	"github.com/andygeiss/faasify/internal/config"
)

//go:embed html.tmpl
var html string

//go:embed styles.css
var styles string

type response struct {
	Styles template.CSS
	Token  string
}

func HandlerFunc(cfg *config.Config) http.HandlerFunc {
	t, _ := template.New("t").Parse(html)
	return func(w http.ResponseWriter, r *http.Request) {
		bearerToken := strings.Split(r.Header.Get("Authorization"), " ")
		t.Execute(w, response{template.CSS(styles), bearerToken[1]})
	}
}
