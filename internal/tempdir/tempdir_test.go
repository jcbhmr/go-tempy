package tempdir_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jcbhmr/go-tempy/v3/internal/tempdir"
	"github.com/stretchr/testify/assert"
)

func readFileString(t *testing.T, path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func TestEnsureTheReturnedFilepathIsNotASymlink(t *testing.T) {
	// TODO: Somehow mock os.TempDir() to a symlinked folder BEFORE the tempdir module is loaded.

	filePath := filepath.Join(tempdir.Default, "unicorn")
	err := os.WriteFile(filePath, []byte("🦄"), 0o644)
	assert.NoError(t, err)

	filePathRealpath, err := filepath.EvalSymlinks(filePath)
	assert.NoError(t, err)
	assert.Equal(t, filePath, filePathRealpath)
	assert.Equal(t, "🦄", readFileString(t, filePath))

	os.Remove(filePath)
}

func TestMain2(t *testing.T) {
	assert.True(t, filepath.IsAbs(tempdir.Default))
}
