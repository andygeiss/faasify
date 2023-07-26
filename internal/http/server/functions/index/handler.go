package index

import (
	_ "embed"
	"html/template"
	"net/http"
)

//go:embed html.tmpl
var html string

//go:embed styles.css
var styles string

type response struct {
	Styles template.CSS
	Token  string
}

func HandlerFunc(token string) http.HandlerFunc {
	t, _ := template.New("t").Funcs(template.FuncMap{}).Parse(html)
	return func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, response{template.CSS(styles), token})
	}
}
