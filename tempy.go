package tempy

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"unsafe"
)

type FileOptions struct {
	// File extension.
	//
	// Mutually exclusive with the Name option.
	//
	// You usually won't need this option. Specify it only when actually needed.
	Extension *string
	// Filename.
	//
	// Mutually exclusive with the Extension option.
	//
	// You usually won't need this option. Specify it only when actually needed.
	Name *string
}

type DirectoryOptions struct {
	// Directory prefix.
	//
	// You usually won't need this option. Specify it only when actually needed.
	//
	// Useful for testing by making it easier to identify cache directories that are created.
	Prefix *string
}

// The temporary path created by the function.
//
// Will be changed to a generic type alias in Go 1.24.
type TaskCallback[T any] func(temporaryPath string) T

// prefix: defaults to ""
func getPath(prefix *string) string {
	var prefix2 string
	if prefix != nil {
		prefix2 = *prefix
	}
	return filepath.Join(os.TempDir(), prefix2+uniqueString())
}

func runTask[T any](temporaryPath string, callback TaskCallback[T]) (rv T, rerr error) {
	var v T
	defer func() {
		err := os.RemoveAll(temporaryPath)
		if err != nil {
			rerr = err
			return
		}
		rv = v
	}()
	v = callback(temporaryPath)
	return
}

// fileContent: string | [*bytes.Buffer] | [](uint8 | uint16 | uint32 | uint64 | int8 | int16 | in32 | int64 | float32 | float64) | []byte
func writeFile(filename string, fileContent any) error {
	var bytes2 []byte
	if v, ok := fileContent.(string); ok {
		bytes2 = []byte(v)
	} else if v, ok := fileContent.(*bytes.Buffer); ok {
		bytes2 = v.Bytes()
	} else if v, ok := fileContent.([]uint16); ok {
		bytes2 = unsafe.Slice((*byte)(unsafe.Pointer(&v[0])), len(v)*2)
	} else if v, ok := fileContent.([]uint32); ok {
		bytes2 = unsafe.Slice((*byte)(unsafe.Pointer(&v[0])), len(v)*4)
	} else if v, ok := fileContent.([]uint64); ok {
		bytes2 = unsafe.Slice((*byte)(unsafe.Pointer(&v[0])), len(v)*8)
	} else if v, ok := fileContent.([]int8); ok {
		bytes2 = unsafe.Slice((*byte)(unsafe.Pointer(&v[0])), len(v))
	} else if v, ok := fileContent.([]int16); ok {
		bytes2 = unsafe.Slice((*byte)(unsafe.Pointer(&v[0])), len(v)*2)
	} else if v, ok := fileContent.([]int32); ok {
		bytes2 = unsafe.Slice((*byte)(unsafe.Pointer(&v[0])), len(v)*4)
	} else if v, ok := fileContent.([]int64); ok {
		bytes2 = unsafe.Slice((*byte)(unsafe.Pointer(&v[0])), len(v)*8)
	} else if v, ok := fileContent.([]float32); ok {
		bytes2 = unsafe.Slice((*byte)(unsafe.Pointer(&v[0])), len(v)*4)
	} else if v, ok := fileContent.([]float64); ok {
		bytes2 = unsafe.Slice((*byte)(unsafe.Pointer(&v[0])), len(v)*8)
	} else if v, ok := fileContent.([]byte); ok {
		bytes2 = v
	} else {
		e := "string | *bytes.Buffer | [](uint8 | uint16 | uint32 | uint64 | int8 | int16 | in32 | int64 | float32 | float64) | []byte"
		panic(fmt.Sprintf("unsupported type: expected %s got %T", e, fileContent))
	}
	return os.WriteFile(filename, bytes2, 0o666)
}

func writeStream(filename string, stream io.ReadCloser) error {
	defer stream.Close()
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, stream)
	if err != nil {
		return err
	}
	return nil
}

// Get a temporary file path you can write to.
func TemporaryFile(options *FileOptions) string {
	if options == nil {
		options = &FileOptions{}
	}
	if options.Name != nil {
		if options.Extension != nil {
			panic(errors.New("the Name and Extension options are mutually exclusive"))
		}

		return filepath.Join(os.TempDir(), *options.Name)
	}

	path := getPath(nil)
	var suffix string
	if options.Extension != nil {
		extensionNoLeadingDot := *options.Extension
		if extensionNoLeadingDot[0] == '.' {
			extensionNoLeadingDot = extensionNoLeadingDot[1:]
		}
		suffix = "." + extensionNoLeadingDot
	}
	return path + suffix
}

// The callback resolves with a temporary file path you can write to. The file is automatically cleaned up after the callback is executed.
//
// Returns T after the callback is executed and the file is cleaned up.
func TemporaryFileTask[T any](callback TaskCallback[T], options *FileOptions) (T, error) {
	return runTask(TemporaryFile(options), callback)
}

// Get a temporary directory path. The directory is created for you.
func TemporaryDirectory(options *DirectoryOptions) (string, error) {
	if options == nil {
		options = &DirectoryOptions{}
	}
	var prefix string
	if options.Prefix != nil {
		prefix = *options.Prefix
	}
	directory := getPath(&prefix)
	err := os.Mkdir(directory, 0o777)
	if err != nil {
		return "", err
	}
	return directory, nil
}

// The callback resolves with a temporary directory path you can write to. The directory is automatically cleaned up after the callback is executed.
//
// Returns T after the callback is executed and the directory is cleaned up.
func TemporaryDirectoryTask[T any](callback TaskCallback[T], options *DirectoryOptions) (T, error) {
	d, err := TemporaryDirectory(options)
	if err != nil {
		var v T
		return v, err
	}
	return runTask(d, callback)
}

// Write data to a random temp file.
//
// fileContent: string | [*bytes.Buffer] | [](uint8 | uint16 | uint32 | uint64 | int8 | int16 | in32 | int64 | float32 | float64) | []byte | [io.ReadCloser]
func TemporaryWrite(fileContent any, options *FileOptions) (string, error) {
	filename := TemporaryFile(options)
	if v, ok := fileContent.(io.ReadCloser); ok {
		err := writeStream(filename, v)
		if err != nil {
			return "", err
		}
	} else {
		err := writeFile(filename, fileContent)
		if err != nil {
			return "", err
		}
	}
	return filename, nil
}

// Write data to a random temp file. The file is automatically cleaned up after the callback is executed.
//
// Returns T after the callback is executed and the file is cleaned up.
//
// fileContent: string | [*bytes.Buffer] | [](uint8 | uint16 | uint32 | uint64 | int8 | int16 | in32 | int64 | float32 | float64) | []byte | io.ReadCloser
func TemporaryWriteTask[T any](fileContent any, callback TaskCallback[T], options *FileOptions) (T, error) {
	f, err := TemporaryWrite(fileContent, options)
	if err != nil {
		var v T
		return v, err
	}
	return runTask(f, callback)
}

// Synchronously write data to a random temp file.
//
// fileContent: string | [*bytes.Buffer] | [](uint8 | uint16 | uint32 | uint64 | int8 | int16 | in32 | int64 | float32 | float64) | []byte
func TemporaryWriteSync(fileContent any, options *FileOptions) (string, error) {
	filename := TemporaryFile(options)
	err := writeFile(filename, fileContent)
	if err != nil {
		return "", err
	}
	return filename, nil
}

// Constant
var RootTemporaryDirectory = tempDir
