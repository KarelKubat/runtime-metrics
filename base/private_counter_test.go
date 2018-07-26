package runtimemetrics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInternalCounter(t *testing.T) {
	const HIGHEST = 10000

	c := newCounter()
	for i := 0; i < HIGHEST; i++ {
		assert.Equal(t, int64(i), c.report())
		c.mark()
	}
}
