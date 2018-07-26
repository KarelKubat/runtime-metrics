package runtimemetrics

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestInternalSum(t *testing.T) {
	s := newSum()

	for i := 1; i <= 10; i++ {
		s.mark(float64(i))
	}

	sum, n := s.report()

	assert.Equal(t, int64(10), n)
	assert.Equal(t, 55.0, sum)
}
