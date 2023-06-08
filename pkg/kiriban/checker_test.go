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
		{100, []Kind{KindTrailingZeros{}}},
		{2000, []Kind{KindTrailingZeros{}}},
		{110000000, nil},
		{101, nil},
		{123, []Kind{KindConsecutive{}}},
		{124, nil},
		{321, []Kind{KindConsecutive{}}},
		{111, []Kind{KindRepdigit{}}},
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

func BenchmarkChecker_IsKiriban(b *testing.B) {
	c, _ := NewChecker()
	for i := 0; i < b.N; i++ {
		c.IsKiriban(i)
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
		{"50 is kiriban when 50 min", input{nc(MinValueFunc(50)), 50}, true},
		{"50 is not kiriban when 51 min", input{nc(MinValueFunc(51)), 50}, false},
		{"56 is kiriban when 50 min", input{nc(MinValueFunc(50)), 56}, true},
		{"101 is kiriban when set as an exceptional kiriban", input{nc(ExceptionalKiribanFunc(exs(101, 103))), 101}, true},
		{"102 is kiriban when not set as an exceptional kiriban", input{nc(ExceptionalKiribanFunc(exs(101, 103))), 102}, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := test.in.checker
			assert.Equal(t, test.out, c.IsKiriban(test.in.val))
		})
	}
}
