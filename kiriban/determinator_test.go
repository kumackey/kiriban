package kiriban

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeterminator_IsKiriban(t *testing.T) {
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

	d, _ := NewDeterminator()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.out, d.IsKiriban(test.in))
		})
	}
}

func TestDeterminator_KiribanKinds(t *testing.T) {
	tests := []struct {
		in        int
		kind      Kind
		isKiriban bool
	}{
		{100, KindRound{}, true},
		{2000, KindRound{}, true},
		{110000000, nil, false},
		{101, nil, false},
		{123, KindConsecutive{}, true},
		{124, nil, false},
		{321, KindConsecutive{}, true},
		{111, KindRepdigit{}, true},
		{110, nil, false},
	}

	d, _ := NewDeterminator()
	for _, test := range tests {
		t.Run(strconv.Itoa(test.in), func(t *testing.T) {
			kind, ok := d.KiribanKind(test.in)
			assert.Equal(t, test.kind, kind)
			assert.Equal(t, test.isKiriban, ok)
		})
	}
}

func TestDeterminator_Next(t *testing.T) {
	t.Parallel()

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

	d, _ := NewDeterminator()
	for _, test := range tests {
		t.Run(strconv.Itoa(test.in)+"->"+strconv.Itoa(test.out), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.out, d.Next(test.in))
		})
	}
}

func TestDeterminator_Previous(t *testing.T) {
	t.Parallel()

	type out struct {
		prev int
		err  error
	}
	tests := []struct {
		in  int
		out out
	}{
		{100, out{99, nil}},
		{101, out{100, nil}},
		{102, out{100, nil}},
		{123, out{111, nil}},
		{200, out{123, nil}},
		{10, out{0, ErrorNoPreviousKiriban}},
		{0, out{0, ErrorNoPreviousKiriban}},
		{200000000, out{123456789, nil}}, // too late...
	}

	d, _ := NewDeterminator()
	for _, test := range tests {
		name := strconv.Itoa(test.in) + "->" + strconv.Itoa(test.out.prev)
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			prev, err := d.Previous(test.in)
			assert.ErrorIs(t, err, test.out.err)
			assert.Equal(t, test.out.prev, prev)
		})
	}
}

func TestDeterminator_IsKiribanWithOptions(t *testing.T) {
	d := func(opts ...OptionFunc) *Determinator {
		d, _ := NewDeterminator(opts...)
		return d
	}
	exs := func(vals ...int) []ExceptionalKiriban {
		eks := make([]ExceptionalKiriban, 0, len(vals))
		for _, v := range vals {
			eks = append(eks, ExceptionalKiriban{v, "sample reason"})
		}
		return eks
	}
	type input struct {
		determinator *Determinator
		val          int
	}

	tests := []struct {
		name string
		in   input
		out  bool
	}{
		{"87654 is kiriban when set min consecutive digits to 5", input{d(SetMinConsecutiveDigits(5)), 87654}, true},
		{"87654 is not kiriban when set min consecutive digits to 4", input{d(SetMinConsecutiveDigits(6)), 87654}, false},
		{"101 is kiriban when set as an exceptional kiriban", input{d(ExceptionalKiribanOption(exs(101, 103))), 101}, true},
		{"102 is kiriban when not set as an exceptional kiriban", input{d(ExceptionalKiribanOption(exs(101, 103))), 102}, false},
		{"0 is not kiriban when enabled digit-based round determination", input{d(EnableDigitBasedRoundDetermination()), 0}, false},
		{"20 is kiriban when enabled digit-based round determination", input{d(EnableDigitBasedRoundDetermination()), 20}, true},
		{"200 is kiriban when enabled digit-based round determination", input{d(EnableDigitBasedRoundDetermination()), 200}, true},
		{"230 is not kiriban when enabled digit-based round determination", input{d(EnableDigitBasedRoundDetermination()), 230}, false},
		{"2000 is kiriban when enabled digit-based round determination", input{d(EnableDigitBasedRoundDetermination()), 2000}, true},
		{"2300 is kiriban when enabled digit-based round determination", input{d(EnableDigitBasedRoundDetermination()), 2300}, true},
		{"23000 is kiriban when enabled digit-based round determination", input{d(EnableDigitBasedRoundDetermination()), 23000}, true},
		{"23400 is not kiriban when enabled digit-based round determination", input{d(EnableDigitBasedRoundDetermination()), 23400}, false},
		{"2222 is kiriban when set min repdigit digits to 4", input{d(SetMinRepDigitDigits(4)), 2222}, true},
		{"2222 is not kiriban when set min repdigit digits to 5", input{d(SetMinRepDigitDigits(5)), 2222}, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			d := test.in.determinator
			assert.Equal(t, test.out, d.IsKiriban(test.in.val))
		})
	}
}
