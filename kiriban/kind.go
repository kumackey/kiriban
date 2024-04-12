package kiriban

import "fmt"

// Kind is a type of kiriban.
type Kind interface {
	String() string
}

// KindConsecutive is a kind of kiriban, which is consecutive numbers.
// ex) 1234, 2345, 3456
type KindConsecutive struct{}

func (_ KindConsecutive) String() string {
	return "Consecutive"
}

// KindRound is a kind of kiriban, which is round numbers.
// ex) 10, 300, 2000
type KindRound struct{}

func (_ KindRound) String() string {
	return "Round"
}

// KindRepdigit is a kind of kiriban, which is repdigit.
// ex) 11, 222, 4444
type KindRepdigit struct{}

func (_ KindRepdigit) String() string {
	return "Repdigit"
}

// KindExceptionalKiriban is a kind of kiriban, which is exceptional kiriban.
type KindExceptionalKiriban struct {
	*ExceptionalKiriban
}

func (k KindExceptionalKiriban) String() string {
	return fmt.Sprintf("Exceptional kiriban: %d, %s", k.Value, k.Reason)
}
