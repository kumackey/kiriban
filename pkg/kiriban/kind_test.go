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
		{name: "Trailing zeros", in: KindTrailingZeros{}, out: "Trailing zeros"},
		{name: "Exceptional kiriban", in: KindExceptionalKiriban{&ExceptionalKiriban{1101, "birthday"}}, out: "Exceptional kiriban: 1101, birthday"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.out, test.in.String())
		})
	}
}
