package kiriban

import (
	"math"
	"strconv"
	"strings"
)

const (
	minKiribanValue = 100

	zeroToNine = "0123456789"
	NineToZero = "9876543210"
)

type Checker struct {
	opt *options
}

func (c *Checker) IsKiriban(v int) bool {
	kinds := c.JudgeKinds(v)

	return 0 < len(kinds)
}

func (c *Checker) JudgeKinds(v int) []Kind {
	if v < c.opt.minValue {
		return nil
	}

	var kinds []Kind

	if c.isConsecutive(v) {
		kinds = append(kinds, KindConsecutive{})
	}

	if c.isTrailingZeros(v) {
		kinds = append(kinds, KindTrailingZeros{})
	}

	if c.isRepDigit(v) {
		kinds = append(kinds, KindRepdigit{})
	}

	if ok, ex := isExceptionalKiriban(v, c.opt.exceptionalKiribans); ok {
		kinds = append(kinds, KindExceptionalKiriban{ex})
	}

	return kinds
}

func (c *Checker) isConsecutive(v int) bool {
	str := strconv.Itoa(v)
	return strings.Contains(zeroToNine, str) || strings.Contains(NineToZero, str)
}

func (c *Checker) isTrailingZeros(num int) bool {
	if num == 0 {
		return false
	}

	for num%10 == 0 {
		num /= 10
	}

	return 0 < num && num < 10
}

func (c *Checker) isRepDigit(v int) bool {
	digits := len(strconv.Itoa(v))
	for i := 0; i < 10; i++ {
		if (int(math.Pow10(digits))-1)/9*i == v {
			return true
		}
	}

	return false
}

func isExceptionalKiriban(v int, exs []ExceptionalKiriban) (bool, *ExceptionalKiriban) {
	for _, e := range exs {
		if v == e.Value {
			return true, &e
		}
	}

	return false, nil
}

func (c *Checker) Next(v int) int {
	// TODO: It takes a long time when the next number is too far away.
	for {
		v++
		if c.IsKiriban(v) {
			return v
		}
	}
}

func NewChecker(optFuncs ...OptionFunc) (*Checker, error) {
	// default options
	opts := &options{
		minValue: minKiribanValue,
	}

	for _, opt := range optFuncs {
		err := opt(opts)
		if err != nil {
			return nil, err
		}
	}

	return &Checker{opt: opts}, nil
}
