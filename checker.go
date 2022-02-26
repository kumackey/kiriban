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

type Checker struct{}

func (c Checker) IsKiriban(v int) (bool, []Kind) {

	if v < minKiribanValue {
		return false, nil
	}

	kinds := c.judgeKinds(v)

	return 0 < len(kinds), kinds
}

func (c Checker) judgeKinds(v int) []Kind {
	var kinds []Kind
	str := strconv.Itoa(v)

	if strings.Contains(zeroToNine, str) || strings.Contains(NineToZero, str) {
		kinds = append(kinds, KindConsecutive{})
	}

	if seriesOfZero.MatchString(str) {
		kinds = append(kinds, KindSeriesOfZero{})
	}

	if repdigit.MatchString(str) {
		kinds = append(kinds, KindRepdigit{})
	}

	return kinds
}
