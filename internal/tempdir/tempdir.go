package tempdir

import (
	"os"
	"path/filepath"
)

// Get the real path of the system temp directory.
//
// Constant.
var Default = func() string {
	t := os.TempDir()
	r, err := filepath.EvalSymlinks(t)
	if err != nil {
		panic(err)
	}
	return r
}()
