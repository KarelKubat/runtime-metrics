package namedset

import (
	"github.com/stretchr/testify/assert"

	named "github.com/KarelKubat/runtime-metrics/named"

	"testing"
	"time"
)

func TestAveragePerDurationSet(t *testing.T) {
	var NAMES = []string{
		"a", "b", "c", "d", "e",
	}
	const DURATION = time.Duration(0.1 * float64(time.Second))

	set := NewAveragePerDurationSet()

	// Create and add some named metrics
	for _, name := range NAMES {
		err := set.Add(named.NewAveragePerDuration(name, DURATION))
		assert.NoError(t, err)
	}

	// Re-adding won't work
	for _, name := range NAMES {
		err := set.Add(named.NewAveragePerDuration(name, DURATION))
		assert.Error(t, err)
	}

	// Names are available
	assert.Equal(t, set.Names(), NAMES)
}
