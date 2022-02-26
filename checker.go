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
	seriesOfZeros = regexp.MustCompile(`^[1-9]0+$`)
	repdigit      = regexp.MustCompile(`^(1+|2+|3+|4+|5+|6+|7+|8+|9+)$`)
)

type Checker struct{}

func (c Checker) IsKiriban(v int) bool {
	if v < minKiribanValue {
		return false
	}

	str := strconv.Itoa(v)

	// consecutive numbers
	if strings.Contains(zeroToNine, str) || strings.Contains(NineToZero, str) {
		return true
	}

	if seriesOfZeros.MatchString(str) {
		return true
	}

	if repdigit.MatchString(str) {
		return true
	}

	return false
}
