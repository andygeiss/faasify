package resource

import (
	"os"
)

// WriteTempFilesFromMap writes every file from the map to the temporary file system.
func WriteTempFilesFromMap(in map[string][]byte) (prefix string, err error) {
	file, err := os.CreateTemp("", "asset")
	if err != nil {
		return "", err
	}
	prefix = file.Name()
	if err := file.Close(); err != nil {
		return "", err
	}
	if err := os.Remove(prefix); err != nil {
		return "", err
	}
	if err := WriteFilesFromMap(in, prefix); err != nil {
		return "", err
	}
	return prefix, nil
}
