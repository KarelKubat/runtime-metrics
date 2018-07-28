package namedset

import (
	"github.com/stretchr/testify/assert"

	named "github.com/KarelKubat/runtime-metrics/named"

	"testing"
)

func TestAverageSet(t *testing.T) {
	var NAMES = []string{
		"a", "b", "c", "d", "e",
	}

	set := NewAverageSet()

	// Create and add some named metrics
	for _, name := range NAMES {
		err := set.Add(named.NewAverage(name))
		assert.NoError(t, err)
	}

	// Re-adding won't work
	for _, name := range NAMES {
		err := set.Add(named.NewAverage(name))
		assert.Error(t, err)
	}

	// Names are available
	assert.Equal(t, set.Names(), NAMES)
}
