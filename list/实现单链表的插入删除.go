package list

import "fmt"

/*
结构：
哨兵节点 -> [1, &] -> [2, &] -> [3,&]
		| 					   ^
		|----------------------|
其中哨兵节点的结构包
head 指向头节点
tail 指向尾节点
len 链表的长度
*/

type Guard struct {
	head *ListNode
	tail *ListNode
	len  int
}

func NewListNode() *Guard {
	// 初始化头尾都指向空节点
	return &Guard{
		nil,
		nil,
		0,
	}
}

// 头插法
func (g *Guard) HeadInterpolation(m *ListNode) {
	// m的下一个节点指向头节点
	m.Next = g.head
	// 哨兵的head指向新的m头节点，我把head理解为指向头节点的虚拟头节点
	g.head = m
	g.len++
	// 边界限制
	if g.len == 1 {
		// 当插入的是第一个节点时，那么尾节点也是它
		g.tail = m
	}
}

// 尾插法
func (g *Guard) TailInterpolation(m *ListNode) {
	// 当没有节点时，插入时后头尾节点都是它
	if g.len == 0 {
		g.head = m
		g.tail = m
		g.len++
		return
	}

	m.Next = nil
	// 原先tail的下一个节点需要先指向m节点
	g.tail.Next = m
	// tail指向m节点
	g.tail = m
	g.len++
}

// 中间插入法
func (g *Guard) MiddleInterpolation(m *ListNode, position int) error {
	if position > g.len {
		return fmt.Errorf("out of index")
	} else if position == 0 {
		g.HeadInterpolation(m)
		return nil
	} else if position == g.len {
		g.TailInterpolation(m)
		return nil
	} else {
		cur := g.head
		// 先找到要插入的位置
		for i := 1; i < position; i++ {
			cur = cur.Next
		}
		m.Next = cur.Next
		cur.Next = m
		g.len++
		return nil
	}
}

func (g *Guard) Print() {
	cur := g.head
	for cur != nil {
		fmt.Println(cur)
		cur = cur.Next
	}
}

func (g *Guard) Format() (list []int) {
	cur := g.head
	for cur != nil {
		list = append(list, cur.Node)
		cur = cur.Next
	}
	return list
}
