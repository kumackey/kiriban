package kiriban

import "errors"

var (
	ErrorInvalidMinValue = errors.New("invalid min value")
)

type options struct {
	minValue            int
	exceptionalKiribans []ExceptionalKiriban
}

type OptionFunc func(*options) error

func MinValueOption(v int) OptionFunc {
	return func(o *options) error {
		if v <= 0 {
			return ErrorInvalidMinValue
		}
		o.minValue = v
		return nil
	}
}

type ExceptionalKiriban struct {
	Value  int
	Reason string
}

func ExceptionalKiribanOption(eks []ExceptionalKiriban) OptionFunc {
	return func(o *options) error {
		o.exceptionalKiribans = eks
		return nil
	}
}
