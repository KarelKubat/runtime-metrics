package threadsafe

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func markCountForDuration(c *CountPerDuration, duration time.Duration) {
	// First slice
	c.Mark()
	c.Mark()

	// Second slice
	time.Sleep(time.Duration(duration))
	c.Mark()
	c.Mark()
	c.Mark()

	// Third slice
	time.Sleep(time.Duration(duration))

	// Fourth slice
	time.Sleep(time.Duration(duration))
}

func TestCountPerDuration(t *testing.T) {
	const THREADS = 1000
	const DURATION = time.Duration(0.1 * float64(time.Second))

	c := NewCountPerDuration(DURATION)

	var wg sync.WaitGroup
	for t := 0; t < THREADS; t++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			markCountForDuration(c, DURATION)
		}()
	}

	// First slice
	count, _ := c.Report()
	assert.Equal(t, int64(0), count)

	// Second slice
	time.Sleep(time.Duration(DURATION))
	count, _ = c.Report()
	assert.Equal(t, int64(2*THREADS), count)

	// Third slice
	time.Sleep(time.Duration(DURATION))
	count, _ = c.Report()
	assert.Equal(t, int64(3*THREADS), count)

	// Fourth slice
	time.Sleep(time.Duration(DURATION))
	count, _ = c.Report()
	assert.Equal(t, int64(0), count)

	wg.Wait()
}
