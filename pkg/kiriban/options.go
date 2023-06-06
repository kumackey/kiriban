package kiriban

import "errors"

type options struct {
	minValue int
}

var (
	ErrorInvalidMinValue = errors.New("invalid min value")
)

type OptionFunc func(*options) error

func MinValue(v int) OptionFunc {
	return func(o *options) error {
		if v <= 0 {
			return ErrorInvalidMinValue
		}
		o.minValue = v
		return nil
	}
}
