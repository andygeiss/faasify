package function

import (
	"compress/gzip"
	"encoding/json"
	"net/http"
)

func Process[REQ any, RES any](w http.ResponseWriter, r *http.Request, fn func(req REQ) (RES, error)) (*RES, error) {
	w.Header().Add("Content-Encoding", "gzip")
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	gw, _ := gzip.NewWriterLevel(w, gzip.BestCompression)
	defer gw.Close()
	var res RES
	var req REQ
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		json.NewEncoder(gw).Encode(res)
		return nil, err
	}
	res, err := fn(req)
	if err != nil {
		json.NewEncoder(gw).Encode(res)
		return nil, err
	}
	json.NewEncoder(gw).Encode(res)
	return &res, nil
}
