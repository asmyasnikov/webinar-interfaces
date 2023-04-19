package burn

type component interface {
	Name() string
}

type Burn struct {
	piece component
}

func (b Burn) Name() string {
	return "burned(" + b.piece.Name() + ")"
}

func New(piece component) Burn {
	return Burn{piece: piece}
}
