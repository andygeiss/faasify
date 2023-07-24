package resource

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

// WriteJsonToPath encodes data as JSON and saves it to the filesystem by path.
func WriteJsonToPath[T any](in T, path string) (err error) {
	_, err = os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			return err
		}
	}
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_SYNC, 0644)
	if err != nil {
		return err
	}
	if err := json.NewEncoder(file).Encode(in); err != nil {
		return err
	}
	if err := file.Close(); err != nil {
		return err
	}
	return nil
}
