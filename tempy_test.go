package tempy_test

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jcbhmr/go-tempy/v3"
	"github.com/jcbhmr/go-tempy/v3/internal/tempdir"
	"github.com/stretchr/testify/assert"
)

func TestFile(t *testing.T) {
	assert.True(t, strings.Contains(tempy.TemporaryFile(nil), tempdir.Default))
	assert.False(t, strings.HasSuffix(tempy.TemporaryFile(nil), "."))
	assert.False(t, strings.HasSuffix(tempy.TemporaryFile(&tempy.FileOptions{Extension: nil}), "."))
	assert.True(t, strings.HasSuffix(tempy.TemporaryFile(&tempy.FileOptions{Extension: ptr("png")}), ".png"))
	assert.True(t, strings.HasSuffix(tempy.TemporaryFile(&tempy.FileOptions{Extension: ptr(".png")}), ".png"))
	assert.False(t, strings.HasSuffix(tempy.TemporaryFile(&tempy.FileOptions{Extension: ptr(".png")}), "..png"))
	assert.True(t, strings.HasSuffix(tempy.TemporaryFile(&tempy.FileOptions{Name: ptr("custom-name.md")}), "custom-name.md"))

	assert.Panics(t, func() {
		tempy.TemporaryFile(&tempy.FileOptions{Name: ptr("custom-name.md"), Extension: ptr(".ext")})
	})

	assert.Panics(t, func() {
		tempy.TemporaryFile(&tempy.FileOptions{Name: ptr("custom-name.md"), Extension: ptr("")})
	})

	assert.NotPanics(t, func() {
		tempy.TemporaryFile(&tempy.FileOptions{Name: ptr("custom-name.md"), Extension: nil})
	})
}

func touch(t *testing.T, path string) {
	f, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
}

func pathExists(t *testing.T, path string) bool {
	_, err := os.Stat(path)
	if err != nil && errors.Is(err, fs.ErrNotExist) {
		return false
	} else if err != nil {
		t.Fatal(err)
	}
	return true
}

func TestFileTask(t *testing.T) {
	var temporaryFilePath string
	rv, err := tempy.TemporaryFileTask(func(temporaryPath string) (string, error) {
		touch(t, temporaryPath)
		temporaryFilePath = temporaryPath
		return temporaryPath, nil
	}, nil)
	assert.NoError(t, err)
	assert.Equal(t, temporaryFilePath, rv)
	assert.False(t, pathExists(t, temporaryFilePath))
}

func TaskCleansUpEvenIfCallbackThrows(t *testing.T) {
	var temporaryDirectoryPath string
	_, err := tempy.TemporaryDirectoryTask(func(temporaryPath string) (string, error) {
		temporaryDirectoryPath = temporaryPath
		return "", errors.New("catch me if you can!")
	}, nil)
	assert.ErrorContains(t, err, "catch me if you can!")

	assert.False(t, pathExists(t, temporaryDirectoryPath))

	// Go-specific: make sure it also cleans up if the callback panics.
	assert.PanicsWithError(t, "catch me if you can!", func() {
		tempy.TemporaryDirectoryTask(func(temporaryPath string) (string, error) {
			temporaryDirectoryPath = temporaryPath
			panic(errors.New("catch me if you can!"))
		}, nil)
	})
	assert.False(t, pathExists(t, temporaryDirectoryPath))
}

func TestDirectory(t *testing.T) {
	prefix := "name_"

	s, err := tempy.TemporaryDirectory(nil)
	assert.NoError(t, err)
	assert.True(t, strings.Contains(s, tempdir.Default))

	s, err = tempy.TemporaryDirectory(&tempy.DirectoryOptions{Prefix: &prefix})
	assert.NoError(t, err)
	assert.True(t, strings.HasPrefix(filepath.Base(s), prefix))
}

func TestDirectoryTask(t *testing.T) {
	var temporaryDirectoryPath string
	rv, err := tempy.TemporaryDirectoryTask(func(temporaryPath string) (string, error) {
		temporaryDirectoryPath = temporaryPath
		return temporaryPath, nil
	}, nil)
	assert.NoError(t, err)
	assert.Equal(t, temporaryDirectoryPath, rv)
	assert.False(t, pathExists(t, temporaryDirectoryPath))
}

func readFileString(t *testing.T, filePath string) string {
	b, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func TestWriteString(t *testing.T) {
	filePath, err := tempy.TemporaryWrite("unicorn", &tempy.FileOptions{Name: ptr("test.png")})
	assert.NoError(t, err)
	assert.Equal(t, "unicorn", readFileString(t, filePath))
	assert.Equal(t, "test.png", filepath.Base(filePath))
}

func TestWriteTaskString(t *testing.T) {
	var temporaryFilePath string
	rv, err := tempy.TemporaryWriteTask("", func(temporaryPath string) (string, error) {
		temporaryFilePath = temporaryPath
		return temporaryPath, nil
	}, nil)
	assert.NoError(t, err)
	assert.Equal(t, temporaryFilePath, rv)
	assert.False(t, pathExists(t, temporaryFilePath))
}

func TestWriteBuffer(t *testing.T) {
	filePath, err := tempy.TemporaryWrite([]byte("unicorn"), nil)
	assert.NoError(t, err)
	assert.Equal(t, "unicorn", readFileString(t, filePath))
}

func TestWriteStream(t *testing.T) {
	readable := io.NopCloser(strings.NewReader("unicorn"))

	filePath, err := tempy.TemporaryWrite(readable, nil)
	assert.NoError(t, err)
	assert.Equal(t, "unicorn", readFileString(t, filePath))
}

type failingStream struct {
	state int
}

func (f *failingStream) Read(p []byte) (int, error) {
	if f.state == 0 {
		f.state++
		return copy(p, []byte("unicorn")), nil
	} else if f.state == 1 {
		f.state++
		return 0, errors.New("catch me if you can!")
	} else {
		panic(fmt.Sprintf("state=%d", f.state))
	}
}

func (f *failingStream) Close() error {
	if f.state == 2 {
		return nil
	} else {
		panic(fmt.Sprintf("state=%d", f.state))
	}
}

func TestWriteStreamFailingStream(t *testing.T) {
	readable := &failingStream{}

	// Must be a io.ReadCloser
	var _ io.ReadCloser = readable

	_, err := tempy.TemporaryWrite(readable, nil)
	assert.ErrorContains(t, err, "catch me if you can!")
}

func TestWriteSync(t *testing.T) {
	filePath, err := tempy.TemporaryWriteSync("unicorn", nil)
	assert.NoError(t, err)
	assert.Equal(t, "unicorn", readFileString(t, filePath))
}

func TestRoot(t *testing.T) {
	assert.True(t, len(tempy.RootTemporaryDirectory) > 0)
	assert.True(t, filepath.IsAbs(tempy.RootTemporaryDirectory))
}
