package apple

func New() Apple {
	return Apple{}
}

type Apple struct{}

func (Apple) Name() string {
	return "apple"
}
