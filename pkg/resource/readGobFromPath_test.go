package resource_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/andygeiss/faasify/pkg/assert"
	"github.com/andygeiss/faasify/pkg/resource"
)

func TestReadGobFromPath(t *testing.T) {
	path := filepath.Join("testdata", "foo.gob")
	_ = os.Remove("testdata")
	type Data struct{ Name string }
	_ = resource.WriteGobToPath(Data{Name: "foo"}, path)
	got, err := resource.ReadGobFromPath[Data](path)
	assert.That("err should be nil", t, err, nil)
	assert.That("Name should be foo", t, got.Name, "foo")
}
