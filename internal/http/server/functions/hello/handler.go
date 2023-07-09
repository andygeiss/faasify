package hello

import (
	"net/http"
)

func HandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`Hello world from Go!`))
	}
}
