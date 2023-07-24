package resource

import (
	"errors"
	"os"
	"path/filepath"
)

// WriteFilesFromMap writes every file from the map to the file system.
func WriteFilesFromMap(in map[string][]byte, prefix string) (err error) {
	for path, content := range in {
		fqn := filepath.Join(prefix, path)
		base := filepath.Dir(fqn)
		_, err = os.Stat(base)
		if errors.Is(err, os.ErrNotExist) {
			if err := os.MkdirAll(base, 0755); err != nil {
				return err
			}
		}
		if err := os.WriteFile(fqn, content, 0644); err != nil {
			return err
		}
	}
	return nil
}
