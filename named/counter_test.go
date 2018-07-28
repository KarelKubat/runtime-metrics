package named

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounter(t *testing.T) {
	const HIGHEST = 10000

	c, err := NewCounter("counter")
	assert.NoError(t, err)

	for i := 0; i < HIGHEST; i++ {
		assert.Equal(t, int64(i), c.Report())
		c.Mark()
	}
}