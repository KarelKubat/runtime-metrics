package base

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T) {
	const HIGHEST = 10000

	c := NewCount()
	for i := 0; i < HIGHEST; i++ {
		assert.Equal(t, int64(i), c.Report())
		c.Mark()
	}

	c.Reset()
	assert.Equal(t, int64(0), c.Report())
}
