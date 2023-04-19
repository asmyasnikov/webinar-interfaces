package freeze

type component interface {
	Name() string
}

type Freeze struct {
	piece component
}

func (f Freeze) Name() string {
	return "frozen(" + f.piece.Name() + ")"
}

func (f Freeze) Warm() component {
	return f.piece
}

func New(piece component) Freeze {
	return Freeze{piece: piece}
}
