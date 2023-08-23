package resource_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/andygeiss/faasify/pkg/assert"
	"github.com/andygeiss/faasify/pkg/resource"
)

func TestWriteFilesFromMap(t *testing.T) {
	path := filepath.Join("testdata")
	_ = os.Remove("testdata")
	filesMap := map[string][]byte{
		"testdata/bar.txt": []byte("bar"),
		"testdata/foo.txt": []byte("foo"),
	}
	err := resource.WriteFilesFromMap(filesMap, path)
	_, err1 := os.Stat(filepath.Join("testdata", "bar.txt"))
	_, err2 := os.Stat(filepath.Join("testdata", "foo.txt"))
	assert.That("err should be nil", t, err, nil)
	assert.That("err1 should be nil", t, err1, nil)
	assert.That("err2 should be nil", t, err2, nil)
}
