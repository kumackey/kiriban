package kiriban

import "fmt"

type Kind interface {
	String() string
}

type KindConsecutive struct{}

func (_ KindConsecutive) String() string {
	return "Consecutive"
}

type KindTrailingZeros struct{}

func (_ KindTrailingZeros) String() string {
	return "Trailing zeros"
}

type KindRepdigit struct{}

func (_ KindRepdigit) String() string {
	return "Repdigit"
}

type KindExceptionalKiriban struct {
	*ExceptionalKiriban
}

func (k KindExceptionalKiriban) String() string {
	return fmt.Sprintf("Exceptional kiriban: %d, %s", k.Value, k.Reason)
}
