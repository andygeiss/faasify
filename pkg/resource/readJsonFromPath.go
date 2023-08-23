package resource

import (
	"encoding/json"
	"os"
)

// ReadJsonFromPath reads data from a JSON file into a struct.
func ReadJsonFromPath[T any](path string) (out T, err error) {
	file, err := os.Open(path)
	if err != nil {
		return out, err
	}
	var tmp T
	if err := json.NewDecoder(file).Decode(&tmp); err != nil {
		return out, err
	}
	if err := file.Close(); err != nil {
		return out, err
	}
	return tmp, nil
}
