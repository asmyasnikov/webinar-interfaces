package potato

func New() Potato {
	return Potato{}
}

type Potato struct{}

func (Potato) Name() string {
	return "potato"
}
