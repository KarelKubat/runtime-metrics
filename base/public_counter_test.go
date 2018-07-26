package runtimemetrics

import (
	"github.com/stretchr/testify/assert"

	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	const THREADS = 1000
	const COUNTS_PER_THREAD = 10000

	c := NewCounter()

	var wg sync.WaitGroup
	for t := 0; t < THREADS; t++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < COUNTS_PER_THREAD; i++ {
				c.Mark()
			}
		}()
	}
	wg.Wait()

	assert.Equal(t, int64(COUNTS_PER_THREAD*THREADS), c.Report())
}
