package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type mockGetter struct {
	v   string
	err error
}

func (m mockGetter) Get() (string, error) {
	return m.v, m.err
}

func TestHandleGet(t *testing.T) {
	for _, tt := range []struct {
		s   getter
		err error
	}{
		{
			s: mockGetter{
				v:   "abc",
				err: nil,
			},
			err: nil,
		},
	} {
		t.Run("", func(t *testing.T) {
			require.Equal(t, tt.err, HandleGet(tt.s))
		})
	}
}
