package kiriban

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsKiriban(t *testing.T) {
	tests := []struct {
		name string
		in   int
		out  bool
	}{
		{"100 is kiriban", 100, true},
		{"101 is not kiriban", 101, false},
		{"123 is kiriban", 123, true},
		{"124 is not kiriban", 124, false},
		{"321 is kiriban", 321, true},
		{"111 is kiriban", 111, true},
		{"110 is not kiriban", 110, false},
		{"90 is not kiriban", 90, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.out, IsKiriban(test.in))
		})
	}
}
