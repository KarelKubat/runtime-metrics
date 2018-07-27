package named

import (
	"time"

	"github.com/stretchr/testify/assert"

	"testing"
)

func TestRegistry(t *testing.T) {
	const NAME = "test_name"
	const DURATION = time.Duration(0.1 * float64(time.Second))

	// Registering new names may not cause problems
	_, err := NewAverage(NAME)
	assert.NoError(t, err)

	// Re-registering it must error out
	_, err = NewAverage(NAME)
	assert.Error(t, err)
	_, err = NewAveragePerDuration(NAME, DURATION)
	assert.Error(t, err)
	_, err = NewCounter(NAME)
	assert.Error(t, err)
	_, err = NewCounterPerDuration(NAME, DURATION)
	assert.Error(t, err)
	_, err = NewSum(NAME)
	assert.Error(t, err)
	_, err = NewSumPerDuration(NAME, DURATION)
	assert.Error(t, err)
}
