package namedset

import (
	"github.com/stretchr/testify/assert"

	named "github.com/KarelKubat/runtime-metrics/named"

	"testing"
)

func TestSumSet(t *testing.T) {
	var NAMES = []string{
		"a", "b", "c", "d", "e",
	}

	set := NewSumSet()

	// Create and add some named metrics
	for _, name := range NAMES {
		err := set.Add(named.NewSum(name))
		assert.NoError(t, err)
	}

	// Re-adding won't work
	for _, name := range NAMES {
		err := set.Add(named.NewSum(name))
		assert.Error(t, err)
	}

	// Names are available
	assert.Equal(t, set.Names(), NAMES)
}
