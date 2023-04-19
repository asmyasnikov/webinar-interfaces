package mix

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type mockComponent string

func (mock mockComponent) Name() string {
	return string(mock)
}

func TestMixName(t *testing.T) {
	for _, tt := range []struct {
		components []component
		name       string
	}{
		{
			components: []component{
				mockComponent("a"),
				mockComponent("b"),
				mockComponent("c"),
			},
			name: "mixed(a,b,c)",
		},
		{
			components: []component{},
			name:       "mixed()",
		},
	} {
		t.Run("", func(t *testing.T) {
			require.Equal(t, tt.name, New(tt.components...).Name())
		})
	}
}
