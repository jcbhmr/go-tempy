package tempdir_test

import (
	"fmt"
	"os"

	"github.com/jcbhmr/go-tempy/v3/internal/tempdir"
)

func ExampleDefault() {
	fmt.Printf("resolved path   : %s\n", tempdir.Default)
	fmt.Printf("possible symlink: %s\n", os.TempDir())
	// Possible output:
	// resolved path   : /private/var/folders/3x/jf5977fn79jbglr7rk0tq4d00000gn/T
	// possible symlink: /var/folders/3x/jf5977fn79jbglr7rk0tq4d00000gn/T
}
