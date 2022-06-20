package list

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMerge2OrderlyListNode(t *testing.T) {
	cases := []struct {
		expected *ListNode
		actual   []*ListNode
	}{
		{
			expected: &ListNode{
				Node: 1,
				Next: &ListNode{
					Node: 1,
					Next: &ListNode{
						Node: 2,
						Next: &ListNode{
							Node: 3,
							Next: &ListNode{
								Node: 4,
								Next: &ListNode{
									Node: 4,
									Next: nil,
								},
							},
						},
					},
				},
			},
			actual: []*ListNode{
				{1,
					&ListNode{2,
						&ListNode{4,
							nil}},
				},
				{
					1,
					&ListNode{3,
						&ListNode{4,
							nil}},
				},
			},
		},
	}
	for _, c := range cases {
		require.Equal(t, c.expected, merge2OrderlyListNode(c.actual[0], c.actual[1]))
	}
}
