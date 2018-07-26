package runtimemetrics

import (
	"time"

	"github.com/stretchr/testify/assert"

	"testing"
)

func TestInternalCounterPerDuration(t *testing.T) {
	const DURATION = time.Duration(0.1 * float64(time.Second))

	c := newCounterPerDuration(DURATION)

	// First slice
	c.mark()
	c.mark()
	count, _ := c.report()
	assert.Equal(t, int64(0), count)

	// Second slice
	time.Sleep(time.Duration(DURATION))
	c.mark()
	c.mark()
	c.mark()
	count, _ = c.report()
	assert.Equal(t, int64(2), count)

	// Third slice
	time.Sleep(time.Duration(DURATION))
	count, _ = c.report()
	assert.Equal(t, int64(3), count)

	// Fourth slice
	time.Sleep(time.Duration(DURATION))
	count, _ = c.report()
	assert.Equal(t, int64(0), count)
}
