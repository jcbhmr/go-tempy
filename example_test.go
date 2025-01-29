package tempy_test

import (
	"fmt"

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
	// /tmp/a1b2c3d4a1b2c3d4a1b2c3d4a1b2c3d4
	// /tmp/a1b2c3d4a1b2c3d4a1b2c3d4a1b2c3d4.png
	// /tmp/unicorn.png
	// /tmp/a1b2c3d4a1b2c3d4a1b2c3d4a1b2c3d4
}

func ExampleTemporaryFileTask() {
	type T = any
	tempy.TemporaryFileTask(func(tempFile string) (T, error) {
		fmt.Println(tempFile)
		return nil, nil
	}, nil)
	// Possible output: /tmp/a1b2c3d4a1b2c3d4a1b2c3d4a1b2c3d4
}

func ExampleTemporaryDirectory() {
	fmt.Println(tempy.TemporaryDirectory(nil))

	fmt.Println(tempy.TemporaryDirectory(&tempy.DirectoryOptions{
		Prefix: ptr("name_"),
	}))

	// Possible output:
	// /tmp/a1b2c3d4a1b2c3d4a1b2c3d4a1b2c3d4
	// /tmp/name_a1b2c3d4a1b2c3d4a1b2c3d4a1b2c3d4
}

func ExampleTemporaryDirectoryTask() {
	type T = any
	tempy.TemporaryDirectoryTask(func(tempDir string) (T, error) {
		fmt.Println(tempDir)
		return nil, nil
	}, nil)
	// Possible output: /tmp/a1b2c3d4a1b2c3d4a1b2c3d4a1b2c3d4
}

func ExampleTemporaryWrite() {
	fmt.Println(tempy.TemporaryWrite("🦄", nil))
	// Possible output: /tmp/a1b2c3d4a1b2c3d4a1b2c3d4a1b2c3d4
}

func ExampleTemporaryWriteTask() {
	type T = any
	tempy.TemporaryWriteTask("🦄", func(tempFile string) (T, error) {
		fmt.Println(tempFile)
		return nil, nil
	}, nil)
	// Possible output: /tmp/a1b2c3d4a1b2c3d4a1b2c3d4a1b2c3d4
}

func ExampleTemporaryWriteSync() {
	fmt.Println(tempy.TemporaryWriteSync("🦄", nil))
	// Possible output: /tmp/a1b2c3d4a1b2c3d4a1b2c3d4a1b2c3d4
}
