package kiriban

import (
	"math"
	"strconv"
	"strings"
)

const (
	zeroToNine = "0123456789"
	nineToZero = "9876543210"
)

type Determinator struct {
	*options
}

// IsKiriban returns true if the given value is kiriban.
func (c *Determinator) IsKiriban(v int) bool {
	return 0 < len(c.KiribanKinds(v))
}

// KiribanKinds returns kiriban kinds of the given value.
func (c *Determinator) KiribanKinds(v int) []Kind {
	if v < 0 {
		// If the value is negative, convert it to a positive value.
		v = -v
	}

	var kinds []Kind

	if c.isConsecutive(v) {
		kinds = append(kinds, KindConsecutive{})
	}

	if c.isRound(v) {
		kinds = append(kinds, KindRound{})
	}

	if c.isRepDigit(v) {
		kinds = append(kinds, KindRepdigit{})
	}

	if ok, ex := isExceptionalKiriban(v, c.exceptionalKiribans); ok {
		kinds = append(kinds, KindExceptionalKiriban{ex})
	}

	return kinds
}

func (c *Determinator) isConsecutive(v int) bool {
	str := strconv.Itoa(v)
	if len(str) < c.minConsecutiveDigits {
		//　If the number is less than three digits,
		//　it does not appear to be a continuous number and is not determined to be Consecutive.
		// ex) 1, 12, 32
		return false
	}
	return strings.Contains(zeroToNine, str) || strings.Contains(nineToZero, str)
}

func (c *Determinator) isRound(num int) bool {
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

func (c *Determinator) isRepDigit(v int) bool {
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
func (c *Determinator) Next(v int) int {
	// TODO: It takes a long time when the next number is too far away.
	for {
		v++
		if c.IsKiriban(v) {
			return v
		}
	}
}

// NewDeterminator returns a new Determinator.
func NewDeterminator(optFuncs ...OptionFunc) (*Determinator, error) {
	const (
		defaultMinConsecutiveDigits = 3
	)

	// default options
	opts := &options{
		minConsecutiveDigits: defaultMinConsecutiveDigits,
	}

	for _, opt := range optFuncs {
		err := opt(opts)
		if err != nil {
			return nil, err
		}
	}

	return &Determinator{opts}, nil
}
