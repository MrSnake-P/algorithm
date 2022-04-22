package string

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReverseString(t *testing.T) {
	cases := []struct {
		actual []byte
		expect []byte
	}{
		{
			actual: []byte("hello"),
			expect: []byte("olleh"),
		},
		{
			actual: []byte("hannah"),
			expect: []byte("hannah"),
		},
	}

	for _, v := range cases {
		reverseString(v.actual)
		require.Equal(t, v.expect, v.actual)
	}
}
