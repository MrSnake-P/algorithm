package dp

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMaxSubArray(t *testing.T) {

	cases := []struct {
		nums     []int
		expected int
	}{
		{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expected: 6,
		},
	}

	for _, c := range cases {
		require.Equal(t, 6, maxSubArray2(c.nums))
	}
}
