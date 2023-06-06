package kiriban

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJudgeKinds(t *testing.T) {
	tests := []struct {
		in  int
		out []Kind
	}{
		{100, []Kind{SeriesOfZero}},
		{101, nil},
		{123, []Kind{Consecutive}},
		{124, nil},
		{321, []Kind{Consecutive}},
		{111, []Kind{Repdigit}},
		{110, nil},
		{90, nil},
	}

	for _, test := range tests {
		t.Run(strconv.Itoa(test.in), func(t *testing.T) {
			assert.Equal(t, test.out, JudgeKinds(test.in))
		})
	}
}

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
