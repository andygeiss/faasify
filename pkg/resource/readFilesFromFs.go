package resource

import (
	"io/fs"
	"os"
)

// ReadFilesFromFs reads every file from a file system into a memory map.
func ReadFilesFromFs(sys fs.FS) (out map[string][]byte, err error) {
	out = make(map[string][]byte)
	if err := fs.WalkDir(sys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path == "." {
			return nil
		}
		if d.IsDir() {
			return nil
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		out[path] = content
		return nil
	}); err != nil {
		return nil, err
	}
	return out, nil
}
