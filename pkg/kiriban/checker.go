package kiriban

import (
	"math"
	"strconv"
	"strings"
)

const (
	minKiribanValue = 100

	zeroToNine = "0123456789"
	nineToZero = "9876543210"
)

type Checker struct {
	*options
}

// IsKiriban returns true if the given value is kiriban.
func (c *Checker) IsKiriban(v int) bool {
	return 0 < len(c.JudgeKinds(v))
}

// JudgeKinds returns kiriban kinds of the given value.
func (c *Checker) JudgeKinds(v int) []Kind {
	if v < 0 {
		// If the value is negative, convert it to a positive value.
		v = -v
	}

	if v < c.minValue {
		return nil
	}

	var kinds []Kind

	if c.isConsecutive(v) {
		kinds = append(kinds, KindConsecutive{})
	}

	if c.isRoundNumber(v) {
		kinds = append(kinds, KindRoundNumber{})
	}

	if c.isRepDigit(v) {
		kinds = append(kinds, KindRepdigit{})
	}

	if ok, ex := isExceptionalKiriban(v, c.exceptionalKiribans); ok {
		kinds = append(kinds, KindExceptionalKiriban{ex})
	}

	return kinds
}

func (c *Checker) isConsecutive(v int) bool {
	str := strconv.Itoa(v)
	if len(str) < 3 {
		//　If the number is less than three digits,
		//　it does not appear to be a continuous number and is not determined to be Consecutive.
		// ex) 1, 12, 32
		return false
	}
	return strings.Contains(zeroToNine, str) || strings.Contains(nineToZero, str)
}

func (c *Checker) isRoundNumber(num int) bool {
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
	for i := 1; i < 10; i++ {
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

// Next returns the next kiriban value.
func (c *Checker) Next(v int) int {
	// TODO: It takes a long time when the next number is too far away.
	for {
		v++
		if c.IsKiriban(v) {
			return v
		}
	}
}

// NewChecker returns a new Checker.
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

	return &Checker{opts}, nil
}
