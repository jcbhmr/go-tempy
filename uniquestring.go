package tempy

import (
	"crypto/rand"
	"fmt"
)

func uniqueString() string {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", b)
}
