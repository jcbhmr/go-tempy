package tempy

import "crypto/rand"

func uniqueString() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return string(b)
}
