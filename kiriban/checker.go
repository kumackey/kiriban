package kiriban

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

const (
	zeroToNine = "0123456789"
	nineToZero = "9876543210"
)

type Checker struct {
	*options
}

// IsKiriban returns true if the given value is kiriban.
func (c *Checker) IsKiriban(v int) bool {
	_, ok := c.KiribanKind(v)
	return ok
}

// KiribanKind returns kiriban kind of the given value.
func (c *Checker) KiribanKind(v int) (Kind, bool) {
	if v < 0 {
		// If the value is negative, convert it to a positive value.
		v = -v
	}

	if ok, ex := isExceptionalKiriban(v, c.exceptionalKiribans); ok {
		return KindExceptionalKiriban{ex}, true
	}

	if c.isConsecutive(v) {
		return KindConsecutive{}, true
	}

	if c.isRound(v) {
		return KindRound{}, true
	}

	if c.isRepDigit(v) {
		return KindRepdigit{}, true
	}

	return nil, false
}

func (c *Checker) isConsecutive(v int) bool {
	str := strconv.Itoa(v)
	if len(str) < c.minConsecutiveDigits {
		//　If the number is less than three digits,
		//　it does not appear to be a continuous number and is not determined to be Consecutive.
		// ex) 1, 12, 32
		return false
	}
	return strings.Contains(zeroToNine, str) || strings.Contains(nineToZero, str)
}

func (c *Checker) isRound(num int) bool {
	str := strconv.Itoa(num)
	if len(str) == 1 {
		// 0, 1, 2, ...,9 are not round numbers.
		return false
	}

	zeros := func() int {
		if c.digitBasedRoundDetermination {
			return len(str) / 2
		}
		return 1
	}()

	last := str[zeros:]
	return strings.Trim(last, "0") == ""
}

func (c *Checker) isRepDigit(v int) bool {
	digits := len(strconv.Itoa(v))
	if digits < c.minRepDigitDigits {
		return false
	}

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

var ErrorNoPreviousKiriban = errors.New("no previous kiriban")

// Previous returns the previous kiriban value.
func (c *Checker) Previous(v int) (int, error) {
	for {
		if v == 0 {
			return 0, ErrorNoPreviousKiriban
		}

		v--
		if c.IsKiriban(v) {
			return v, nil
		}
	}
}

// NewChecker returns a new Checker.
func NewChecker(optFuncs ...OptionFunc) (*Checker, error) {
	const (
		defaultMinConsecutiveDigits = 3
		defaultMinRepDigitDigits    = 2
	)

	// default options
	opts := &options{
		minConsecutiveDigits: defaultMinConsecutiveDigits,
		minRepDigitDigits:    defaultMinRepDigitDigits,
	}

	for _, opt := range optFuncs {
		err := opt(opts)
		if err != nil {
			return nil, err
		}
	}

	return &Checker{opts}, nil
}
