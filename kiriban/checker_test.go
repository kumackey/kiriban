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
		{"90 is kiriban", 90, true},
		{"1100 is not kiriban", 1100, false},
		{"1 is not kiriban", 1, false},
		{"9 is not kiriban", 9, false},
		{"0 is not kiriban", 0, false},
		{"-100 is kiriban", -100, true},
		{"-111 is kiriban", -111, true},
		{"-123 is kiriban", -123, true},
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
		{100, []Kind{KindRoundNumber{}}},
		{2000, []Kind{KindRoundNumber{}}},
		{110000000, nil},
		{101, nil},
		{123, []Kind{KindConsecutive{}}},
		{124, nil},
		{321, []Kind{KindConsecutive{}}},
		{111, []Kind{KindRepdigit{}}},
		{110, nil},
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
		{0, 10},
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

func TestChecker_IsKiriban_Option(t *testing.T) {
	nc := func(opts ...OptionFunc) *Checker {
		c, _ := NewChecker(opts...)
		return c
	}
	exs := func(vals ...int) []ExceptionalKiriban {
		eks := make([]ExceptionalKiriban, 0, len(vals))
		for _, v := range vals {
			eks = append(eks, ExceptionalKiriban{v, "sample reason"})
		}
		return eks
	}
	type input struct {
		checker *Checker
		val     int
	}

	tests := []struct {
		name string
		in   input
		out  bool
	}{
		{"87654 is kiriban when set min consecutive digits to 5", input{nc(SetMinConsecutiveDigits(5)), 87654}, true},
		{"87654 is not kiriban when set min consecutive digits to 4", input{nc(SetMinConsecutiveDigits(6)), 87654}, false},
		{"101 is kiriban when set as an exceptional kiriban", input{nc(ExceptionalKiribanOption(exs(101, 103))), 101}, true},
		{"102 is kiriban when not set as an exceptional kiriban", input{nc(ExceptionalKiribanOption(exs(101, 103))), 102}, false},
		{"0 is not kiriban when enabled digit-based round determination", input{nc(EnableDigitBasedRoundDetermination()), 0}, false},
		{"20 is kiriban when enabled digit-based round determination", input{nc(EnableDigitBasedRoundDetermination()), 20}, true},
		{"200 is kiriban when enabled digit-based round determination", input{nc(EnableDigitBasedRoundDetermination()), 200}, true},
		{"230 is not kiriban when enabled digit-based round determination", input{nc(EnableDigitBasedRoundDetermination()), 230}, false},
		{"2000 is kiriban when enabled digit-based round determination", input{nc(EnableDigitBasedRoundDetermination()), 2000}, true},
		{"2300 is kiriban when enabled digit-based round determination", input{nc(EnableDigitBasedRoundDetermination()), 2300}, true},
		{"23000 is kiriban when enabled digit-based round determination", input{nc(EnableDigitBasedRoundDetermination()), 23000}, true},
		{"23400 is not kiriban when enabled digit-based round determination", input{nc(EnableDigitBasedRoundDetermination()), 23400}, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := test.in.checker
			assert.Equal(t, test.out, c.IsKiriban(test.in.val))
		})
	}
}
