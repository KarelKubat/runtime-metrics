package base

import (
	"time"

	"github.com/stretchr/testify/assert"

	"testing"
)

func TestCounterPerDuration(t *testing.T) {
	const DURATION = time.Duration(0.1 * float64(time.Second))

	c := NewCounterPerDuration(DURATION)

	// First slice
	c.Mark()
	c.Mark()
	count, _ := c.Report()
	assert.Equal(t, int64(0), count)

	// Second slice
	time.Sleep(time.Duration(DURATION))
	c.Mark()
	c.Mark()
	c.Mark()
	count, _ = c.Report()
	assert.Equal(t, int64(2), count)

	// Third slice
	time.Sleep(time.Duration(DURATION))
	count, _ = c.Report()
	assert.Equal(t, int64(3), count)

	// Fourth slice
	time.Sleep(time.Duration(DURATION))
	count, _ = c.Report()
	assert.Equal(t, int64(0), count)

	// Reset test
	c.Reset()
	assert.Equal(t, int64(0), count)
	
	// First slice
	c.Mark()
	c.Mark()
	count, _ = c.Report()
	assert.Equal(t, int64(0), count)

	// Second slice
	time.Sleep(time.Duration(DURATION))
	c.Mark()
	c.Mark()
	c.Mark()
	c.Reset()
	count, _ = c.Report()
	assert.Equal(t, int64(0), count)
}
