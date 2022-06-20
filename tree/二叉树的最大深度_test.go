package tree

import (
	"fmt"
	"testing"
)

func TestTraversal(t *testing.T) {
	node := TreeNode{value: 1,
		left: &TreeNode{value: 2,
			left:  &TreeNode{3, nil, nil},
			right: &TreeNode{4, nil, nil}},
		right: &TreeNode{value: 2,
			left:  &TreeNode{3, nil, nil},
			right: &TreeNode{4, nil, nil}},
	}
	fmt.Println(traversal(&node))

	node2 := TreeNode{value: 3,
		left: &TreeNode{value: 9,
			left:  nil,
			right: nil,
		},
		right: &TreeNode{value: 20,
			left:  &TreeNode{15, nil, nil},
			right: &TreeNode{7, nil, nil}},
	}
	fmt.Println(traversal(&node2))
}
