package named

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverage(t *testing.T) {
	a, err := NewAverage("average")
	assert.NoError(t, err)

	a.Mark(1.0)
	a.Mark(2.0)
	a.Mark(3.0)

	avg, n := a.Report()

	assert.Equal(t, n, int64(3))
	assert.Equal(t, avg, 2.0)
}
