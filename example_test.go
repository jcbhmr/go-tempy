package tempy_test

import (
	"fmt"
	"log"

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

	p, err := tempy.TemporaryDirectory(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)

	// Possible output:
	// /tmp/a1b2c3d4a1b2c3d4a1b2c3d4a1b2c3d4
	// /tmp/a1b2c3d4a1b2c3d4a1b2c3d4a1b2c3d4.png
	// /tmp/unicorn.png
	// /tmp/a1b2c3d4a1b2c3d4a1b2c3d4a1b2c3d4
}

func ExampleTemporaryFileTask() {
	type T = any
	_, err := tempy.TemporaryFileTask(func(tempFile string) (T, error) {
		fmt.Println(tempFile)
		return nil, nil
	}, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Possible output: /tmp/a1b2c3d4a1b2c3d4a1b2c3d4a1b2c3d4
}

func ExampleTemporaryDirectory() {
	p, err := tempy.TemporaryDirectory(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)

	p, err = tempy.TemporaryDirectory(&tempy.DirectoryOptions{
		Prefix: ptr("name_"),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)

	// Possible output:
	// /tmp/a1b2c3d4a1b2c3d4a1b2c3d4a1b2c3d4
	// /tmp/name_a1b2c3d4a1b2c3d4a1b2c3d4a1b2c3d4
}

func ExampleTemporaryDirectoryTask() {
	type T = any
	_, err := tempy.TemporaryDirectoryTask(func(tempDir string) (T, error) {
		fmt.Println(tempDir)
		return nil, nil
	}, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Possible output: /tmp/a1b2c3d4a1b2c3d4a1b2c3d4a1b2c3d4
}

func ExampleTemporaryWrite() {
	p, err := tempy.TemporaryWrite("🦄", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)
	// Possible output: /tmp/a1b2c3d4a1b2c3d4a1b2c3d4a1b2c3d4
}

func ExampleTemporaryWriteTask() {
	type T = any
	_, err := tempy.TemporaryWriteTask("🦄", func(tempFile string) (T, error) {
		fmt.Println(tempFile)
		return nil, nil
	}, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Possible output: /tmp/a1b2c3d4a1b2c3d4a1b2c3d4a1b2c3d4
}

func ExampleTemporaryWriteSync() {
	p, err := tempy.TemporaryWriteSync("🦄", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)
	// Possible output: /tmp/a1b2c3d4a1b2c3d4a1b2c3d4a1b2c3d4
}
