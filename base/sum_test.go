package base

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestSum(t *testing.T) {
	s := NewSum()

	for i := 1; i <= 10; i++ {
		s.Mark(float64(i))
	}

	sum, n := s.Report()
	assert.Equal(t, int64(10), n)
	assert.Equal(t, 55.0, sum)

	s.Reset()
	sum, n = s.Report()
	assert.Equal(t, int64(0), n)
	assert.Equal(t, 0.0, sum)
}
