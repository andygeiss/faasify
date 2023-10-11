package count

import (
	"net/http"
	"sync"

	"github.com/andygeiss/faasify/internal/config"
	"github.com/andygeiss/faasify/internal/http/server"
)

type Request struct{}
type Response struct {
	Count int    `json:"count"`
	Error string `json:"error"`
}

var (
	count int
	mutex sync.Mutex
)

func HandlerFunc(cfg *config.Config) http.HandlerFunc {
	count = 0
	return func(w http.ResponseWriter, r *http.Request) {
		server.Process[Request, Response](w, r, func(req Request) (Response, error) {
			mutex.Lock()
			defer mutex.Unlock()
			count++
			return Response{Count: count}, nil
		})
	}
}
