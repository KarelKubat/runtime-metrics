package threadsafe

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func markSumForDuration(c *SumPerDuration) {
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

func TestSumPerDuration(t *testing.T) {
	const THREADS = 1000
	const DURATION = time.Duration(0.1 * float64(time.Second))

	c := NewSumPerDuration(DURATION)

	var wg sync.WaitGroup
	for t := 0; t < THREADS; t++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			markSumForDuration(c)
		}()
	}

	// First slice
	sum, count, _ := c.Report()
	assert.Equal(t, int64(0), count)
	assert.Equal(t, 0.0, sum)

	// Second slice
	time.Sleep(time.Duration(DURATION))
	sum, count, _ = c.Report()
	assert.Equal(t, int64(2*THREADS), count)
	assert.Equal(t, 3000.0, sum)

	// Third slice
	time.Sleep(time.Duration(DURATION))
	sum, count, _ = c.Report()
	assert.Equal(t, int64(3*THREADS), count)
	assert.Equal(t, 12000.0, sum)

	// Fourth slice
	time.Sleep(time.Duration(DURATION))
	sum, count, _ = c.Report()
	assert.Equal(t, int64(0), count)
	assert.Equal(t, 0.0, sum)

	wg.Wait()
}