package mix

import (
	"strings"
)

type component interface {
	Name() string
}

type Mix []component

func (m Mix) Name() string {
	b := strings.Builder{}
	b.WriteString("mixed(")
	for i, c := range m {
		if i != 0 {
			b.WriteString(",")
		}
		b.WriteString(c.Name())
	}
	b.WriteString(")")
	return b.String()
}

func New(components ...component) Mix {
	return components
}
