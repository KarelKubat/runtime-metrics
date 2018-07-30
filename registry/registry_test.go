package registry

import (
	"testing"
	"time"

	"github.com/KarelKubat/runtime-metrics/base"
	"github.com/stretchr/testify/assert"
)

func TestRegistry(t *testing.T) {
	const D = time.Duration(0.1 * float64(time.Second))

	var err error

	// Add some metrics. No errors expected.
	err = AddAverage("testname-average",
		base.NewAverage())
	assert.NoError(t, err)
	err = AddAveragePerDuration("testname-average-per-duration",
		base.NewAveragePerDuration(D))
	assert.NoError(t, err)
	err = AddCount("testname-counter",
		base.NewCount())
	assert.NoError(t, err)
	err = AddCountPerDuration("testname-counter-per-duration",
		base.NewCountPerDuration(D))
	assert.NoError(t, err)
	err = AddSum("testname-sum",
		base.NewSum())
	assert.NoError(t, err)
	err = AddSumPerDuration("testname-sum-per-duration",
		base.NewSumPerDuration(D))
	assert.NoError(t, err)

	// Redo, errors expected.
	err = AddAverage("testname-average",
		base.NewAverage())
	assert.Error(t, err)
	err = AddAveragePerDuration("testname-average-per-duration",
		base.NewAveragePerDuration(D))
	assert.Error(t, err)
	err = AddCount("testname-counter",
		base.NewCount())
	assert.Error(t, err)
	err = AddCountPerDuration("testname-counter-per-duration",
		base.NewCountPerDuration(D))
	assert.Error(t, err)
	err = AddSum("testname-sum",
		base.NewSum())
	assert.Error(t, err)
	err = AddSumPerDuration("testname-sum-per-duration",
		base.NewSumPerDuration(D))
	assert.Error(t, err)
}
