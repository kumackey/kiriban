package kiriban

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKind_String(t *testing.T) {
	tests := []struct {
		name string
		in   Kind
		out  string
	}{
		{name: "series of zero", in: SeriesOfZero, out: "Series of zero"},
		{name: "unknown", in: 999, out: "Unknown"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.out, test.in.String())
		})
	}
}
