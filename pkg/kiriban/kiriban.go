package kiriban

import (
	"regexp"
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

func IsKiriban(v int) bool {
	kinds := JudgeKinds(v)

	return 0 < len(kinds)
}
