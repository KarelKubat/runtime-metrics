package threadsafe

import (
	"github.com/stretchr/testify/assert"

	"sync"
	"testing"
)

func TestSum(t *testing.T) {
	const THREADS = 1000
	const LOOPS = 10

	s := NewSum()
	var wg sync.WaitGroup
	for t := 0; t < THREADS; t++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 1; i <= LOOPS; i++ {
				s.Mark(float64(i))
			}
		}()
	}
	wg.Wait()

	sum, n := s.Report()
	assert.Equal(t, int64(THREADS*LOOPS), n)
	assert.Equal(t, 55.0*float64(THREADS), sum)
}
