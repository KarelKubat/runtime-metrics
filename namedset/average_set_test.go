package namedset

import (
	"github.com/stretchr/testify/assert"

	"github.com/KarelKubat/runtime-metrics/base"

	"testing"
)

func TestAverageSet(t *testing.T) {
	var NAMES = []string{
		"a", "b", "c", "d", "e",
	}

	set := NewAverageSet()

	// Create and add some base metrics
	for _, name := range NAMES {
		err := set.Add(name, base.NewAverage())
		assert.NoError(t, err)
	}

	// Re-adding won't work
	for _, name := range NAMES {
		err := set.Add(name, base.NewAverage())
		assert.Error(t, err)
	}

	// Names are available
	assert.Equal(t, set.Names(), NAMES)
}
