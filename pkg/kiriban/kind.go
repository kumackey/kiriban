package kiriban

// enum
const (
	// Consecutive e.g. 123, 234, 345, 456, 567, 678, 789
	Consecutive Kind = iota

	// TrailingZeros e.g. 100, 1000, 10000
	TrailingZeros

	// Repdigit e.g. 111, 222, 333, 444, 555, 666, 777, 888, 999
	Repdigit
)

type Kind int

func (k Kind) String() string {
	switch k {
	case Consecutive:
		return "Consecutive"
	case TrailingZeros:
		return "Trailing zeros"
	case Repdigit:
		return "Repdigit"
	default:
		return "Unknown"
	}
}
