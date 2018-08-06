package namedset

import (
	"github.com/stretchr/testify/assert"

	base "github.com/KarelKubat/runtime-metrics/base"

	"testing"
)

func TestSumSet(t *testing.T) {
	var NAMES = []string{
		"a", "b", "c", "d", "e",
	}

	set := NewSumSet()

	// Create and add some base metrics
	for _, name := range NAMES {
		err := set.Add(name, base.NewSum())
		assert.Nil(t, err)
	}

	// Re-adding won't work
	for _, name := range NAMES {
		err := set.Add(name, base.NewSum())
		assert.Error(t, err)
	}

	// Names are available
	assert.Equal(t, set.Names(), NAMES)
}
