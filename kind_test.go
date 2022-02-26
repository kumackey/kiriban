package kiriban

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKind(t *testing.T) {
	type out struct {
		kind string
	}

	tests := map[string]struct {
		input  Kind
		output out
	}{
		"series of zero": {KindSeriesOfZero{}, out{kind: "series of zero"}},
		"consecutive":    {KindConsecutive{}, out{kind: "consecutive"}},
		"repdigit":       {KindRepdigit{}, out{kind: "repdigit"}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.output.kind, test.input.Kind())
		})
	}
}
