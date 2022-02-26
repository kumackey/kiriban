package kiriban

type Kind interface {
	Kind() string
}

type KindConsecutive struct{}

func (k KindConsecutive) Kind() string {
	return "consecutive"
}

type KindSeriesOfZero struct{}

func (k KindSeriesOfZero) Kind() string {
	return "series of zero"
}

type KindRepdigit struct{}

func (k KindRepdigit) Kind() string {
	return "repdigit"
}
