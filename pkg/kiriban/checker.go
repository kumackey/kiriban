package kiriban

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	minKiribanValue = 100

	zeroToNine = "0123456789"
	NineToZero = "9876543210"
)

var (
	seriesOfZero = regexp.MustCompile(`^[1-9]0+$`)
	repdigit     = regexp.MustCompile(`^(1+|2+|3+|4+|5+|6+|7+|8+|9+)$`)
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
	str := strconv.Itoa(v)

	if strings.Contains(zeroToNine, str) || strings.Contains(NineToZero, str) {
		kinds = append(kinds, Consecutive)
	}

	if seriesOfZero.MatchString(str) {
		kinds = append(kinds, SeriesOfZero)
	}

	if repdigit.MatchString(str) {
		kinds = append(kinds, Repdigit)
	}

	return kinds
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

func NewChecker(opts ...OptionFunc) (*Checker, error) {
	defaultOpt := &options{
		minValue: minKiribanValue,
	}

	for _, opt := range opts {
		err := opt(defaultOpt)
		if err != nil {
			return nil, err
		}
	}

	return &Checker{opt: defaultOpt}, nil
}
