package uniquestring

import (
	"crypto/rand"
	"fmt"
)

// Generate a unique random string.
//
// Returns a 32 character unique string. Matches the length of MD5, which is [unique enough] for non-crypto purposes.
//
// Panics if crypto/rand.Read() returns an error.
//
// [unique enough]: https://stackoverflow.com/a/2444336/64949
func Default() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", b)
}
