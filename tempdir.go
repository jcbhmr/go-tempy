package tempy

import (
	"os"
	"path/filepath"
)

// From https://www.npmjs.com/package/temp-dir
var tempDir = func() string {
	t := os.TempDir()
	r, err := filepath.EvalSymlinks(t)
	if err != nil {
		panic(err)
	}
	return r
}()
