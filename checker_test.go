package kiriban

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChecker_IsKiriban(t *testing.T) {
	type out struct {
		isKiriban bool
		kinds     []Kind
	}

	tests := map[string]struct {
		input  int
		output out
	}{
		"100 is series of zero": {100, out{isKiriban: true, kinds: []Kind{KindSeriesOfZero{}}}},
		"101 is not kiriban":    {101, out{isKiriban: false}},
		"123 is consecutive":    {123, out{isKiriban: true, kinds: []Kind{KindConsecutive{}}}},
		"124 is not kiriban":    {124, out{isKiriban: false}},
		"321 is consecutive":    {321, out{isKiriban: true, kinds: []Kind{KindConsecutive{}}}},
		"111 is repdigit":       {111, out{isKiriban: true, kinds: []Kind{KindRepdigit{}}}},
		"110 is not kiriban":    {110, out{isKiriban: false}},
		"90 is not kiriban":     {90, out{isKiriban: false}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			checker := Checker{}
			isKiriban, kinds := checker.IsKiriban(test.input)
			assert.Equal(t, test.output.isKiriban, isKiriban)
			assert.Equal(t, test.output.kinds, kinds)
		})
	}
}
