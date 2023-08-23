package resource_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/andygeiss/faasify/pkg/assert"
	"github.com/andygeiss/faasify/pkg/resource"
)

func TestWriteTempFilesFromMap(t *testing.T) {
	filesMap := map[string][]byte{
		"testdata/bar.txt": []byte("bar"),
		"testdata/foo.txt": []byte("foo"),
	}
	prefix, err := resource.WriteTempFilesFromMap(filesMap)
	_, err1 := os.Stat(filepath.Join(prefix, "testdata", "bar.txt"))
	_, err2 := os.Stat(filepath.Join(prefix, "testdata", "foo.txt"))
	assert.That("err should be nil", t, err, nil)
	assert.That("err1 should be nil", t, err1, nil)
	assert.That("err2 should be nil", t, err2, nil)
	assert.That("prefix should not be empty", t, prefix != "", true)
}
