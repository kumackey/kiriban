package kiriban

import (
	"iter"
)

func Iter(optFuncs ...OptionFunc) iter.Seq[int] {
	checker, err := NewChecker(optFuncs...)
	if err != nil {
		return func(yield func(int) bool) {
			// オプションが間違えてるときは何も返さない
			return
		}
	}

	return func(yield func(int) bool) {
		for i := 0; ; i++ {
			if checker.IsKiriban(i) {
				if !yield(i) {
					return
				}
			}
		}
	}
}
