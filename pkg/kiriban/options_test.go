package kiriban

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChecker_IsKiriban_Option_Error(t *testing.T) {
	tests := []struct {
		name string
		in   OptionFunc
		out  error
	}{
		{"min value -50", MinValue(-50), ErrorInvalidMinValue},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.in(&options{})
			assert.Equal(t, test.out, err)
		})
	}
}
