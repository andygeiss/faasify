package server

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/andygeiss/faasify/internal/config"
)

// Validate is a generic function to test a handler function.
func Validate[IN any, OUT any](fn http.HandlerFunc, name string, req IN, cfg *config.Config) (res OUT, err error) {
	reqBytes, _ := json.Marshal(req)
	r := httptest.NewRequest("POST", "/"+name, bytes.NewReader(reqBytes))
	rr := httptest.NewRecorder()
	fn.ServeHTTP(rr, r)
	gr, _ := gzip.NewReader(rr.Body)
	defer gr.Close()
	err = json.NewDecoder(gr).Decode(&res)
	return res, err
}
