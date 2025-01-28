package tempy_test

import (
	"fmt"
	"testing"

	"github.com/jcbhmr/go-tempy/v3"
)

func ptr[T any](v T) *T {
	return &v
}

func ExampleTemporaryFile() {
	fmt.Println(tempy.TemporaryFile(nil))

	fmt.Println(tempy.TemporaryFile(&tempy.FileOptions{
		Extension: ptr("png"),
	}))

	fmt.Println(tempy.TemporaryFile(&tempy.FileOptions{
		Name: ptr("unicorn.png"),
	}))

	fmt.Println(tempy.TemporaryDirectory(nil))

	// Possible output:
	// /tmp/tempy-123456789
	// /tmp/tempy-123456789.png
	// /tmp/unicorn.png
	// /tmp/tempy-123456789
}
func TestExampleTemporaryFile(t *testing.T) {
	ExampleTemporaryFile()
}

func ExampleTemporaryFileTask() {
	tempy.TemporaryFileTask(func(tempFile string) struct{} {
		fmt.Println(tempFile)
		return struct{}{}
	}, nil)
	// Possible output: /tmp/tempy-123456789
}
func TestExampleTemporaryFileTask(t *testing.T) {
	ExampleTemporaryFileTask()
}

func ExampleTemporaryDirectory() {
	fmt.Println(tempy.TemporaryDirectory(nil))

	fmt.Println(tempy.TemporaryDirectory(&tempy.DirectoryOptions{
		Prefix: ptr("a"),
	}))

	// Possible output:
	// /tmp/tempy-123456789
	// /tmp/a_123456789
}
func TestExampleTemporaryDirectory(t *testing.T) {
	ExampleTemporaryDirectory()
}

func ExampleTemporaryDirectoryTask() {
	tempy.TemporaryDirectoryTask(func(tempDir string) struct{} {
		fmt.Println(tempDir)
		return struct{}{}
	}, nil)
	// Possible output: /tmp/tempy-123456789
}
func TestExampleTemporaryDirectoryTask(t *testing.T) {
	ExampleTemporaryDirectoryTask()
}

func ExampleTemporaryWrite() {
	fmt.Println(tempy.TemporaryWrite("🦄", nil))
	// Possible output: /tmp/tempy-123456789
}
func TestExampleTemporaryWrite(t *testing.T) {
	ExampleTemporaryWrite()
}

func ExampleTemporaryWriteTask() {
	tempy.TemporaryWriteTask("🦄", func(tempFile string) struct{} {
		fmt.Println(tempFile)
		return struct{}{}
	}, nil)
	// Possible output: /tmp/tempy-123456789
}
func TestExampleTemporaryWriteTask(t *testing.T) {
	ExampleTemporaryWriteTask()
}

func ExampleTemporaryWriteSync() {
	fmt.Println(tempy.TemporaryWriteSync("🦄", nil))
	// Possible output: /tmp/tempy-123456789
}
func TestExampleTemporaryWriteSync(t *testing.T) {
	ExampleTemporaryWriteSync()
}
