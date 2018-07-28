package named

import (
	"time"

	named "github.com/KarelKubat/runtime-metrics/named"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestRegistry(t *testing.T) {
	const NAME = "test_name"
	const DURATION = time.Duration(0.1 * float64(time.Second))

	// Registering new names may not cause problems
	err := RegisterAverage(named.NewAverage(NAME))
	assert.NoError(t, err)

	// Re-registering it must error out
	assert.Error(t, RegisterAverage(named.NewAverage(NAME)))
	assert.Error(t, RegisterCounter(named.NewCounter(NAME)))
	assert.Error(t, RegisterSum(named.NewSum(NAME)))
	assert.Error(t, RegisterCounterPerDuration(
		named.NewCounterPerDuration(NAME, DURATION)))
	assert.Error(t, RegisterAveragePerDuration(
		named.NewAveragePerDuration(NAME, DURATION)))
	assert.Error(t, RegisterSumPerDuration(
		named.NewSumPerDuration(NAME, DURATION)))
}
