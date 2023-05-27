package wasm_demo

import (
	"context"
	_ "embed"
	"encoding/json"
	"log"
	"net/http"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

//go:embed module/fn.wasm
var fn []byte

type response struct {
	Data  map[string]any `json:"data,omitempty"`
	Error error          `json:"error,omitempty"`
}

func HandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := context.Background()
		rt := wazero.NewRuntime(ctx)
		defer rt.Close(ctx)

		wasi_snapshot_preview1.MustInstantiate(ctx, rt)
		mod, _ := rt.Instantiate(ctx, fn)
		fn := mod.ExportedFunction("fn")

		res, err := fn.Call(ctx, 12, 30)
		if err != nil {
			log.Printf("error during function call: %v", err)
		}
		data := map[string]any{
			"result": res,
		}

		out := response{Data: data, Error: err}

		_ = json.NewEncoder(w).Encode(out)
	}
}
