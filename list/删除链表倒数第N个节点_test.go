package list

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	cases := []struct {
		actual   *ListNode
		n        int
		expected *ListNode
	}{
		{actual: &ListNode{1,
			&ListNode{2,
				&ListNode{3,
					&ListNode{4,
						&ListNode{5,
							nil}}}}},
			n: 2,
			expected: &ListNode{1,
				&ListNode{2,
					&ListNode{3,
						&ListNode{5,
							nil}}}},
		},
		{actual: &ListNode{1,
			&ListNode{2,
				&ListNode{3,
					&ListNode{4,
						&ListNode{5,
							nil}}}}},
			n: 5,
			expected: &ListNode{2,
				&ListNode{3,
					&ListNode{4,
						&ListNode{5,
							nil}}}},
		},
	}
	for _, c := range cases {
		require.Equal(t, c.expected, removeNthFromEnd(c.actual, c.n))
	}
}
