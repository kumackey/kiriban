package kiriban

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChecker_IsKiriban(t *testing.T) {
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

	c, _ := NewChecker()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.out, c.IsKiriban(test.in))
		})
	}
}

func TestChecker_JudgeKinds(t *testing.T) {
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

	c, _ := NewChecker()
	for _, test := range tests {
		t.Run(strconv.Itoa(test.in), func(t *testing.T) {
			assert.Equal(t, test.out, c.JudgeKinds(test.in))
		})
	}
}

func TestChecker_Next(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{0, 100},
		{99, 100},
		{100, 111},
		{101, 111},
		{111, 123},
		{123, 200},
		{123456789, 200000000}, // too late...
	}

	c, _ := NewChecker()
	for _, test := range tests {
		t.Run(strconv.Itoa(test.in), func(t *testing.T) {
			assert.Equal(t, test.out, c.Next(test.in))
		})
	}
}
