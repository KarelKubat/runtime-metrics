package namedset

import (
	"github.com/stretchr/testify/assert"

	"github.com/KarelKubat/runtime-metrics/base"

	"testing"
	"time"
)

func TestCountPerDurationSet(t *testing.T) {
	var NAMES = []string{
		"a", "b", "c", "d", "e",
	}
	const DURATION = time.Duration(0.1 * float64(time.Second))

	set := NewCountPerDurationSet()

	// Create and add some base metrics
	for _, name := range NAMES {
		err := set.Add(name, base.NewCountPerDuration(DURATION))
		assert.Nil(t, err)
	}

	// Re-adding won't work
	for _, name := range NAMES {
		err := set.Add(name, base.NewCountPerDuration(DURATION))
		assert.Error(t, err)
	}

	// Names are available
	assert.Equal(t, set.Names(), NAMES)
}
