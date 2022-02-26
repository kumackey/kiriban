package kiriban

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnpassAPIClient_FetchEventList(t *testing.T) {
	tests := map[string]struct {
		input  int
		output bool
	}{
		"100 is kiriban":     {100, true},
		"101 is not kiriban": {101, false},
		"123 is kiriban":     {123, true},
		"124 is not kiriban": {124, false},
		"321 is kiriban":     {321, true},
		"111 is kiriban":     {111, true},
		"110 is not kiriban": {110, false},
		"90 is not kiriban":  {90, false},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			checker := Checker{}
			assert.Equal(t, test.output, checker.IsKiriban(test.input))
		})
	}
}
