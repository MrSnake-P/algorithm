package list

/*
链表问题中需要用到常用的哨兵节点，我把他理解为一个虚拟的头节点；
原链表为a->b->c，则加了哨兵节点的链表即为x->a->b>c；
其作用在于删除链表节点中的时，任何时候`pre.Next=pre.Next.Next`都能成立。
由于如果删除的是头节点，那么其pre为nil，使用pre就会报错。
例1：
给定一个链表: 1->2->3->4->5, n = 2
结果为 1->2->3->5.

进阶：只遍历一遍
*/
/*
解题步骤第一种：（直接法，需要遍历两遍）
1.首先遍历一遍，算出一共有多少节点（N）
2.那么要删除的节点就为N-n+1
*/

type ListNode struct {
	Node int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	guard := &ListNode{} // 哨兵节点，否则需要另外判断是否是删除的是头节点
	guard.Next = head    // 哨兵的下个节点即为头节点
	// 前一个节点
	var pre *ListNode
	// 计算节点数量
	var i int
	cur := guard
	for head != nil {
		head = head.Next
		i++
	}
	for j := 1; j <= i; j++ {
		pre = cur
		cur = cur.Next
		// 移动到要删除的节点前一个位置
		if j == i-n+1 {
			pre.Next = pre.Next.Next
		}
	}
	return guard.Next
}

/*
进阶：只遍历一次
*/
