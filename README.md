# tempy for Go

📂 The [tempy npm package](https://www.npmjs.com/package/tempy) ported to Go

<table align=center><td>

```go
fmt.Println(tempy.TemporaryFile(nil))
fmt.Println(tempy.TemporaryFile(&tempy.FileOptions{Extension: "png"}))
fmt.Println(tempy.TemporaryFile(&tempy.FileOptions{Name: "unicorn.png"}))
fmt.Println(tempy.TemporaryDirectory(nil))
fmt.Println(tempy.TemporaryDirectory(&tempy.DirectoryOptions{Prefix: "name"}))
// Possible output:
// /tmp/a1b2c3d4a1b2c3d4a1b2c3d4a1b2c3d4
// /tmp/a1b2c3d4a1b2c3d4a1b2c3d4a1b2c3d4.png
// /tmp/unicorn.png
// /tmp/a1b2c3d4a1b2c3d4a1b2c3d4a1b2c3d4
// /tmp/name_a1b2c3d4a1b2c3d4a1b2c3d4a1b2c3d4
```

</table>

## Installation

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=Go&logoColor=FFFFFF)

```sh
go get github.com/jcbhmr/go-tempy/v3
```

## Usage

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=Go&logoColor=FFFFFF)

```go
package main

import (
  "fmt"

  "github.com/jcbhmr/go-tempy/v3"
)

func main() {
  fmt.Println(tempy.TemporaryFile(nil))
  fmt.Println(tempy.TemporaryFile(&tempy.FileOptions{Extension: "png"}))
  fmt.Println(tempy.TemporaryFile(&tempy.FileOptions{Name: "unicorn.png"}))
  fmt.Println(tempy.TemporaryDirectory(nil))
  fmt.Println(tempy.TemporaryDirectory(&tempy.DirectoryOptions{Prefix: "name"}))
  // Possible output:
  // /tmp/a1b2c3d4a1b2c3d4a1b2c3d4a1b2c3d4
  // /tmp/a1b2c3d4a1b2c3d4a1b2c3d4a1b2c3d4.png
  // /tmp/unicorn.png
  // /tmp/a1b2c3d4a1b2c3d4a1b2c3d4a1b2c3d4
  // /tmp/name_a1b2c3d4a1b2c3d4a1b2c3d4a1b2c3d4
}
```

[📚 See pkg.go.dev/github.com/jcbhmr/go-tempy/v3 for more docs](https://pkg.go.dev/github.com/jcbhmr/go-tempy/v3)

## Development

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=Go&logoColor=FFFFFF)

This package attempts to mirror the functionality and API surface of [the tempy npm package](https://www.npmjs.com/package/tempy). To that end, we need to convert some JavaScript concepts to Go concepts.

- Union types like `TypedArray | Buffer` are mapped to `any` with a `switch x.(type)` or `if v, ok := x.(T); ok`.
- If possible, JavaScript standard library or Node.js standard library types are mapped to Go standard library types.
- Node.js `Buffer` is `bytes.Buffer`.
- ECMA `TypedArray` variants are `[]<numeric>`. Note that `uint8` and `byte` are the same type.
- ECMA `DataView` is `[]byte`. Byte slices can be viewed into other byte slices so this works OK.
- Node.js streams are `io.*Closer` interfaces. Node.js streams are closed by the stream's consumer. We do the same here.
- All `Promise<T>` values are flattened to just `T`. We let the user spawn goroutines if they want things to be asynchronous.

Also try to keep the version tags in sync. v1.0.0 of tempy on npm should correspond with v1.0.0 of this module.
