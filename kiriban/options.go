package kiriban

import "errors"

var (
	// ErrorInvalidMinValue is returned when the min value is invalid.
	ErrorInvalidMinValue = errors.New("invalid min value")
)

type options struct {
	exceptionalKiribans          []ExceptionalKiriban
	minConsecutiveDigits         int
	digitBasedRoundDetermination bool
}

// OptionFunc is a function to set options.
type OptionFunc func(*options) error

// SetMinConsecutiveDigits sets the minimum consecutive digits.
func SetMinConsecutiveDigits(v int) OptionFunc {
	return func(o *options) error {
		o.minConsecutiveDigits = v
		return nil
	}
}

// EnableDigitBasedRoundDetermination enables digit-based round determination.
func EnableDigitBasedRoundDetermination() OptionFunc {
	return func(o *options) error {
		o.digitBasedRoundDetermination = true
		return nil
	}
}

type ExceptionalKiriban struct {
	Value  int
	Reason string
}

// ExceptionalKiribanOption sets exceptional kiribans.
// Users can set their own kiribans.
func ExceptionalKiribanOption(eks []ExceptionalKiriban) OptionFunc {
	return func(o *options) error {
		o.exceptionalKiribans = eks
		return nil
	}
}
