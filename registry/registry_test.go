package registry

import (
	"testing"
	"time"

	"github.com/KarelKubat/runtime-metrics/named"
	"github.com/stretchr/testify/assert"
)

func TestRegistry(t *testing.T) {
	const D = time.Duration(0.1 * float64(time.Second))

	var err error

	// Add some metrics. No errors expected.
	err = AddAverage(
		named.NewAverage("testname-average"))
	assert.NoError(t, err)
	err = AddAveragePerDuration(
		named.NewAveragePerDuration("testname-average-per-duration", D))
	assert.NoError(t, err)
	err = AddCount(
		named.NewCount("testname-counter"))
	assert.NoError(t, err)
	err = AddCountPerDuration(
		named.NewCountPerDuration("testname-counter-per-duration", D))
	assert.NoError(t, err)
	err = AddSum(
		named.NewSum("testname-sum"))
	assert.NoError(t, err)
	err = AddSumPerDuration(
		named.NewSumPerDuration("testname-sum-per-duration", D))
	assert.NoError(t, err)

	// Redo, errors expected.
	err = AddAverage(
		named.NewAverage("testname-average"))
	assert.Error(t, err)
	err = AddAveragePerDuration(
		named.NewAveragePerDuration("testname-average-per-duration", D))
	assert.Error(t, err)
	err = AddCount(
		named.NewCount("testname-counter"))
	assert.Error(t, err)
	err = AddCountPerDuration(
		named.NewCountPerDuration("testname-counter-per-duration", D))
	assert.Error(t, err)
	err = AddSum(
		named.NewSum("testname-sum"))
	assert.Error(t, err)
	err = AddSumPerDuration(
		named.NewSumPerDuration("testname-sum-per-duration", D))
	assert.Error(t, err)
}
