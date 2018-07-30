package baseset

import (
	"github.com/stretchr/testify/assert"

	"github.com/KarelKubat/runtime-metrics/base"

	"testing"
)

func TestCountSet(t *testing.T) {
	var NAMES = []string{
		"a", "b", "c", "d", "e",
	}

	set := NewCountSet()

	// Create and add some base metrics
	for _, name := range NAMES {
		err := set.Add(name, base.NewCount())
		assert.NoError(t, err)
	}

	// Re-adding won't work
	for _, name := range NAMES {
		err := set.Add(name, base.NewCount())
		assert.Error(t, err)
	}

	// Names are available
	assert.Equal(t, set.Names(), NAMES)
}
