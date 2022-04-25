package list

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInsertListNode(t *testing.T) {
	cases := []struct {
		expected *ListNode
		actual   func(list *Guard) error
	}{
		{expected: &ListNode{2,
			&ListNode{3,
				&ListNode{4,
					&ListNode{5,
						nil}}}},
			actual: func(list *Guard) error {
				list.HeadInterpolation(&ListNode{2, nil})
				list.TailInterpolation(&ListNode{5, nil})
				err := list.MiddleInterpolation(&ListNode{3, nil}, 1)
				if err != nil {
					return err
				}
				err = list.MiddleInterpolation(&ListNode{4, nil}, 2)
				if err != nil {
					return err
				}
				return nil
			},
		},
	}
	for _, c := range cases {
		list := NewNilListNode()
		err := c.actual(list)
		if err != nil {
			t.Fatal()
		}
		require.Equal(t, c.expected, list.head)
		list.Print()
		t.Logf("%+v", list.Format())
	}
}

func TestDeleteListNode(t *testing.T) {
	cases := []struct {
		expected *ListNode
		actual   func() (*Guard, error)
	}{
		{expected: &ListNode{3,
			&ListNode{5,
				&ListNode{6, nil}}},
			actual: func() (*Guard, error) {
				m := &ListNode{2,
					&ListNode{3,
						&ListNode{4,
							&ListNode{5,
								&ListNode{6,
									&ListNode{7, nil}}}}}}
				list := NewListNode(m)
				err := list.DeleteNode(1)
				if err != nil {
					t.Fatal(err)
				}
				err = list.DeleteNode(5)
				if err != nil {
					t.Fatal(err)
				}
				err = list.DeleteNode(2)
				if err != nil {
					t.Fatal(err)
				}
				return list, nil
			},
		},

		{
			expected: nil,
			actual: func() (*Guard, error) {
				list := NewListNode(&ListNode{1, nil})
				err := list.DeleteNode(1)
				if err != nil {
					t.Fatal(err)
				}
				return list, nil
			},
		},
	}
	for _, c := range cases {
		list, _ := c.actual()
		require.Equal(t, c.expected, list.head)
	}
}
