package index

import (
	_ "embed"
	"html/template"
	"net/http"
)

//go:embed html.tmpl
var tmpl string

func HandlerFunc(token string) http.HandlerFunc {
	t, _ := template.New("t").Parse(tmpl)
	return func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, struct{ Token string }{Token: token})
	}
}
