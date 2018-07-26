package runtimemetrics

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const THREADS = 1000
const DURATION = time.Duration(0.1 * float64(time.Second))

func makeMarks(c *AveragePerDuration) {
	// First slice
	c.Mark(1.0)
	c.Mark(2.0)

	// Second slice
	time.Sleep(time.Duration(DURATION))
	c.Mark(3.0)
	c.Mark(4.0)
	c.Mark(5.0)

	// Third slice
	time.Sleep(time.Duration(DURATION))

	// Fourth slice
	time.Sleep(time.Duration(DURATION))
}

func TestAveragePerDuration(t *testing.T) {
	c := NewAveragePerDuration(DURATION)

	var wg sync.WaitGroup
	for t := 0; t < THREADS; t++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			makeMarks(c)
		}()
	}

	// First slice
	avg, count, _ := c.Report()
	assert.Equal(t, int64(0), count)
	assert.Equal(t, 0.0, avg)

	// Second slice
	time.Sleep(time.Duration(DURATION))
	avg, count, _ = c.Report()
	assert.Equal(t, int64(2*THREADS), count)
	assert.Equal(t, 1.5, avg)

	// Third slice
	time.Sleep(time.Duration(DURATION))
	avg, count, _ = c.Report()
	assert.Equal(t, int64(3*THREADS), count)
	assert.Equal(t, 4.0, avg)

	// Fourth slice
	time.Sleep(time.Duration(DURATION))
	avg, count, _ = c.Report()
	assert.Equal(t, int64(0), count)
	assert.Equal(t, 0.0, avg)

	wg.Wait()
}
