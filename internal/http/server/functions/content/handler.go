package content

import (
	_ "embed"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

//go:embed html.tmpl
var html string

//go:embed styles.css
var styles string

type response struct {
	Data   map[string]any
	Styles template.CSS
}

func HandlerFunc(token, domain, url string) http.HandlerFunc {
	t, _ := template.New("t").Funcs(template.FuncMap{}).Parse(html)
	return func(w http.ResponseWriter, r *http.Request) {
		client := http.Client{}
		req, err := http.NewRequest("POST", url+"/stats", nil)
		if err != nil {
			log.Printf("error during http.NewRequest: %v", err)
		}
		req.Header.Set("Authorization", "Bearer "+token)
		res, err := client.Do(req)
		if err != nil {
			log.Printf("error during client.Do: %v", err)
		}
		defer res.Body.Close()
		var data map[string]any
		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			log.Printf("error during json.Decode: %v", err)
		}
		t.Execute(w, response{data, template.CSS(styles)})
	}
}
