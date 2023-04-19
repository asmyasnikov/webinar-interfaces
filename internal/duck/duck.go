package duck

func New() Duck {
	return Duck{}
}

type Duck struct{}

func (Duck) Name() string {
	return "duck"
}

func (Duck) Swim() string {
	return "swim"
}
