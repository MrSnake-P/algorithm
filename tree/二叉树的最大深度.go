package tree

/*
树的问题很多涉及到广度优先和深度优先。
例：
给定一个二叉树，找出其最大深度。二叉树的深度为根节点到最远叶子节点的最长路径上的节点数
说明: 叶子节点是指没有子节点的节点。

给定二叉树 [3,9,20,null,null,15,7]，
    3
   / \
  9  20
    /  \
   15   7
*/

/*
解题步骤：(递归)
每个节点的深度与它左右子树的深度有关，且等于其左右子树最大深度值加上
转换成公式，maxDepth(root) = max(maxDepth(root.left), maxDepth(root.right)) + 1
*/

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func MaxDepth(t *TreeNode) int {
	if t == nil {
		return 0
	}

	return max(MaxDepth(t.left), MaxDepth(t.right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
进阶：非递归的深度遍历优先
递归会导致栈溢出，99%的递归转非递归，都可以通过栈来进行实现。
*/

// todo 如何通过栈求出二叉树最大深度

type Stack struct {
	arr []*TreeNode
	top int
}

func (s *Stack) Push(t *TreeNode) {
	s.arr = append(s.arr, t)
	s.top++
}

func (s *Stack) Pop() {
	s.arr = s.arr[:s.top-1]
	s.top--
}

func (s *Stack) peek() *TreeNode {
	return s.arr[s.top-1]
}

func traversal(t *TreeNode) int {
	if t == nil {
		return 0
	}
	s := &Stack{arr: make([]*TreeNode, 0, 10), top: 0}
	s.Push(t)
	res1 := 1
	res2 := 1
	for s.top != 0 {
		node := s.peek()
		s.Pop()
		if node.right != nil {
			s.Push(node.right)
			res1++
		}
		if node.left != nil {
			s.Push(node.left)
			res2++
		}
	}
	return max(res1, res2)
}
