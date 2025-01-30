# tempy for Go

📂 The [tempy npm package](https://www.npmjs.com/package/tempy) ported to Go

<table align=center><td>

```go
fmt.Println(tempy.TemporaryFile(&tempy.FileOptions{Extension: "png"}))
// Possible output: /tmp/a1b2c3d4a1b2c3d4a1b2c3d4a1b2c3d4.png
```

```go
p, err = tempy.TemporaryDirectory(&tempy.DirectoryOptions{Prefix: "name_"})
if err != nil {
  log.Fatal(err)
}
log.Println(p)
// Possible output: /tmp/name_a1b2c3d4a1b2c3d4a1b2c3d4a1b2c3d4
```

</table>

<p align=center>
  <a href="https://pkg.go.dev/github.com/jcbhmr/go-tempy/v3">Docs</a>
  | <a href="https://github.com/jcbhmr/go-tempy">GitHub</a>
</p>

🔗 Properly resolves `TMPDIR=/dirsymlink` symlinks \
🐿️ Uses Go idioms while maintaining the same API surface as tempy \
📁 Great for quickly writing some data to a temporary file

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
  log.Println(tempy.TemporaryFile(nil))

  log.Println(tempy.TemporaryFile(&tempy.FileOptions{Extension: "png"}))

  log.Println(tempy.TemporaryFile(&tempy.FileOptions{Name: "unicorn.png"}))

  p, err = tempy.TemporaryDirectory(nil)
  if err != nil {
    log.Fatal(err)
  }
  log.Println(p)

  p, err = tempy.TemporaryDirectory(&tempy.DirectoryOptions{Prefix: "name_"})
  if err != nil {
    log.Fatal(err)
  }
  log.Println(p)

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

- Union types like `TypedArray | Buffer` are mapped to `any` with a `switch x.(type)` or `if v, ok := x.(T); ok`. Doc comments should be included to clarify possible types.
- If possible, JavaScript standard library or Node.js standard library types are mapped to Go standard library types.
- Node.js `Buffer` is `[]byte`. `bytes.Buffer` might seem like a good fit but it implements `Reader` and `Writer` interfaces which we don't want.
- `ArrayBuffer` is `[]byte`. `bytes.Buffer` offers too much functionality.
- ECMA `TypedArray` variants are `[]<numeric>`. Note that `byte` is an alias of `uint8`.
- ECMA `DataView` is `[]byte`. Byte slices can be views into other byte slices so this works OK.
- Node.js R/W streams are `io.(Read|Write)Closer` interfaces. Node.js streams are closed by the stream's consumer. We do the same here.
- All `Promise<T>` values are flattened to just `T`. We let the user spawn goroutines if they want things to be asynchronous.

Also try to keep the version tags in sync. v1.0.0 of tempy on npm should correspond with v1.0.0 of this module.
