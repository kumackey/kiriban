package kiriban

type options struct {
	minValue int
}

type OptionFunc func(*options) error

func MinValue(v int) OptionFunc {
	return func(o *options) error {
		o.minValue = v
		return nil
	}
}
