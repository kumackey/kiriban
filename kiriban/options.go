package kiriban

import "errors"

var (
	// ErrorInvalidMinValue is returned when the min value is invalid.
	ErrorInvalidMinValue = errors.New("invalid min value")
)

type options struct {
	minValue             int
	exceptionalKiribans  []ExceptionalKiriban
	minConsecutiveDigits int
}

// OptionFunc is a function to set options.
type OptionFunc func(*options) error

// SetMinValue sets the minimum value.
func SetMinValue(v int) OptionFunc {
	return func(o *options) error {
		if v <= 0 {
			return ErrorInvalidMinValue
		}
		o.minValue = v
		return nil
	}
}

// SetMinConsecutiveDigits sets the minimum consecutive digits.
func SetMinConsecutiveDigits(v int) OptionFunc {
	return func(o *options) error {
		o.minConsecutiveDigits = v
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
