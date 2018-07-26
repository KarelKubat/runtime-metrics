package runtimemetrics

import (
	"github.com/stretchr/testify/assert"

	"sync"
	"testing"
)

func TestAverage(t *testing.T) {
	const THREADS = 1000
	const LOOPS = 10

	a := NewAverage()
	var wg sync.WaitGroup
	for t := 0; t < THREADS; t++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 1; i <= LOOPS; i++ {
				a.Mark(float64(i))
			}
		}()
	}
	wg.Wait()

	avg, n := a.Report()
	assert.Equal(t, int64(THREADS*LOOPS), n)
	assert.Equal(t, 5.5, avg)
}
