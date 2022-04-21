package array

import (
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

func TestIntersect(t *testing.T) {
	cases := []struct {
		num1     []int
		num2     []int
		expected []int
	}{
		{
			num1:     []int{1, 2, 2, 1},
			num2:     []int{2, 2},
			expected: []int{2, 2},
		},
		{
			num1:     []int{4, 9, 5},
			num2:     []int{9, 4, 9, 8, 4},
			expected: []int{4, 9},
		},
	}
	for _, v := range cases {
		result := intersect(v.num1, v.num2)
		// 对结果进行一个排序
		sort.Ints(result)
		require.Equal(t, v.expected, result)
	}
}
