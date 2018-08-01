package base

import (
	"time"

	"github.com/stretchr/testify/assert"

	"testing"
)

func TestAveragePerDuration(t *testing.T) {
	const DURATION = time.Duration(0.1 * float64(time.Second))

	a := NewAveragePerDuration(DURATION)

	// First slice
	a.Mark(1.0)
	a.Mark(2.0)
	avg, count, _ := a.Report()
	assert.Equal(t, int64(0), count)
	assert.Equal(t, 0.0, avg)

	// Second slice
	time.Sleep(time.Duration(DURATION))
	a.Mark(3.0)
	a.Mark(4.0)
	a.Mark(5.0)
	avg, count, _ = a.Report()
	assert.Equal(t, int64(2), count)
	assert.Equal(t, 1.5, avg)

	// Third slice
	time.Sleep(time.Duration(DURATION))
	avg, count, _ = a.Report()
	assert.Equal(t, int64(3), count)
	assert.Equal(t, 4.0, avg)

	// Fourth slice
	time.Sleep(time.Duration(DURATION))
	avg, count, _ = a.Report()
	assert.Equal(t, int64(0), count)
	assert.Equal(t, 0.0, avg)

	// Reset test
	a.Reset()
	avg, count, _ = a.Report()
	assert.Equal(t, int64(0), count)
	assert.Equal(t, 0.0, avg)

	// First slice
	a.Mark(1.0)
	a.Mark(2.0)
	avg, count, _ = a.Report()
	assert.Equal(t, int64(0), count)
	assert.Equal(t, 0.0, avg)

	// Second slice
	time.Sleep(time.Duration(DURATION))
	a.Mark(3.0)
	a.Mark(4.0)
	a.Mark(5.0)
	a.Reset()

	avg, count, _ = a.Report()
	assert.Equal(t, int64(0), count)
	assert.Equal(t, 0.0, avg)
}
