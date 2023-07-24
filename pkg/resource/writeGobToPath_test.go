package resource_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/andygeiss/faasify/pkg/assert"
	"github.com/andygeiss/faasify/pkg/resource"
)

func TestWriteGobToPath(t *testing.T) {
	path := filepath.Join("testdata", "foo.gob")
	_ = os.Remove("testdata")
	gob := struct {
		Bar string
	}{
		Bar: "bar",
	}
	err := resource.WriteGobToPath(gob, path)
	assert.That("err should be nil", t, err, nil)
}
