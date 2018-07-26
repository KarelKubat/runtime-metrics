package runtimemetrics

import (
	"time"

	"github.com/stretchr/testify/assert"

	"testing"
)

func TestInternalSumPerDuration(t *testing.T) {
	const DURATION = time.Duration(0.1 * float64(time.Second))

	c := newSumPerDuration(DURATION)

	// First slice
	c.mark(1.0)
	c.mark(2.0)
	sum, count, _ := c.report()
	assert.Equal(t, int64(0), count)
	assert.Equal(t, 0.0, sum)

	// Second slice
	time.Sleep(time.Duration(DURATION))
	c.mark(3.0)
	c.mark(4.0)
	c.mark(5.0)
	sum, count, _ = c.report()
	assert.Equal(t, int64(2), count)
	assert.Equal(t, 3.0, sum)

	// Third slice
	time.Sleep(time.Duration(DURATION))
	sum, count, _ = c.report()
	assert.Equal(t, int64(3), count)
	assert.Equal(t, 12.0, sum)

	// Fourth slice
	time.Sleep(time.Duration(DURATION))
	sum, count, _ = c.report()
	assert.Equal(t, int64(0), count)
	assert.Equal(t, 0.0, sum)
}
