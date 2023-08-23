package resource_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/andygeiss/faasify/pkg/assert"
	"github.com/andygeiss/faasify/pkg/resource"
)

func TestReadJsonFromPath(t *testing.T) {
	path := filepath.Join("testdata", "foo.json")
	_ = os.Remove("testdata")
	type Data struct {
		Name string `json:"name"`
	}
	_ = resource.WriteJsonToPath(Data{Name: "foo"}, path)
	got, err := resource.ReadJsonFromPath[Data](path)
	assert.That("err should be nil", t, err, nil)
	assert.That("Name should be foo", t, got.Name, "foo")
}
