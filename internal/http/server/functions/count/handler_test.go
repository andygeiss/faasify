package count_test

import (
	"testing"

	"github.com/andygeiss/faasify/internal/config"
	"github.com/andygeiss/faasify/internal/http/server"
	"github.com/andygeiss/faasify/internal/http/server/functions/count"
	"github.com/andygeiss/faasify/pkg/assert"
)

func TestCountSuccess(t *testing.T) {
	cfg := &config.Config{}
	fn := count.HandlerFunc(cfg)
	req := count.Request{}
	res, err := server.Validate[count.Request, count.Response](fn, "count", req, cfg)
	assert.That("no error is returned", t, err, nil)
	assert.That("count is 1", t, res.Count, 1)
}
