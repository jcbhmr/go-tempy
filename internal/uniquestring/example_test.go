package uniquestring_test

import (
	"fmt"

	"github.com/jcbhmr/go-tempy/v3/internal/uniquestring"
)

func ExampleDefault() {
	fmt.Println(uniquestring.Default())
	// Possible output: b4de2a49c8ffa3fbee04446f045483b2
}
