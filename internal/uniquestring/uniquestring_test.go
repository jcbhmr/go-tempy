package uniquestring_test

import (
	"testing"

	"github.com/jcbhmr/go-tempy/v3/internal/uniquestring"
	"github.com/stretchr/testify/assert"
)

func TestMain2(t *testing.T) {
	assert.Equal(t, 32, len(uniquestring.Default()))

	created := map[string]struct{}{}

	for i := 0; i < 100_000; i++ {
		string2 := uniquestring.Default()

		if _, ok := created[string2]; ok {
			t.Errorf("%s already exists", string2)
		}

		assert.Equal(t, 32, len(string2))

		created[string2] = struct{}{}
	}
}
