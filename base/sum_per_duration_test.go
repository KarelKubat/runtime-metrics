package base

import (
	"time"

	"github.com/stretchr/testify/assert"

	"testing"
)

func TestSumPerDuration(t *testing.T) {
	const DURATION = time.Duration(0.1 * float64(time.Second))

	s := NewSumPerDuration(DURATION)

	// First slice
	s.Mark(1.0)
	s.Mark(2.0)
	sum, count, _ := s.Report()
	assert.Equal(t, int64(0), count)
	assert.Equal(t, 0.0, sum)

	// Second slice
	time.Sleep(time.Duration(DURATION))
	s.Mark(3.0)
	s.Mark(4.0)
	s.Mark(5.0)
	sum, count, _ = s.Report()
	assert.Equal(t, int64(2), count)
	assert.Equal(t, 3.0, sum)

	// Third slice
	time.Sleep(time.Duration(DURATION))
	sum, count, _ = s.Report()
	assert.Equal(t, int64(3), count)
	assert.Equal(t, 12.0, sum)

	// Fourth slice
	time.Sleep(time.Duration(DURATION))
	sum, count, _ = s.Report()
	assert.Equal(t, int64(0), count)
	assert.Equal(t, 0.0, sum)

	// Reset test
	s.Reset()
	sum, count, _ = s.Report()
	assert.Equal(t, int64(0), count)
	assert.Equal(t, 0.0, sum)
	// First slice
	s.Mark(1.0)
	s.Mark(2.0)
	sum, count, _ = s.Report()
	assert.Equal(t, int64(0), count)
	assert.Equal(t, 0.0, sum)

	// Second slice
	time.Sleep(time.Duration(DURATION))
	s.Mark(3.0)
	s.Mark(4.0)
	s.Mark(5.0)
	s.Reset()
	sum, count, _ = s.Report()
	assert.Equal(t, int64(0), count)
	assert.Equal(t, 0.0, sum)
}
