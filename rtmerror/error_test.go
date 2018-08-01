package rtmerror

import (
	"github.com/stretchr/testify/assert"

	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	err := NewError("%d is the answer", 42)
	assert.Error(t, err)
	assert.Equal(t, "42 is the answer", err.Error())
	assert.Equal(t, false, err.Retryable())
	assert.Error(t, err)

	err.WithError(fmt.Errorf("but what is the question"))
	assert.Equal(t, "42 is the answer (but what is the question)", err.Error())
	assert.Equal(t, false, err.Retryable())
	assert.Error(t, err)

	err = NewError("pi is not %.2f", 3.14).
		WithError(fmt.Errorf("e is not %.2f", 2.71)).
		WithRetryable(true)
	assert.Equal(t, "pi is not 3.14 (e is not 2.71)", err.Error())
	assert.Equal(t, true, err.Retryable())
	assert.Error(t, err)
}
