package resource_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/andygeiss/faasify/pkg/assert"
	"github.com/andygeiss/faasify/pkg/resource"
)

func TestWriteJsonToPath(t *testing.T) {
	path := filepath.Join("testdata", "foo.json")
	_ = os.Remove("testdata")
	type Data struct {
		Name string `json:"name"`
	}
	err := resource.WriteJsonToPath(Data{Name: "bar"}, path)
	assert.That("err should be nil", t, err, nil)
}
