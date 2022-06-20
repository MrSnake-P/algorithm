package list

/*
将两个有序链表合并为一个新的有序链表并返回。
新链表是通过拼接给定的两个链表的所有节点组成的。
例：
输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/
/*
解题步骤：
1.设置哨兵节点 guard
2.每次 guard 的 next 指向两个链表中小的元素
3.直至有一个链表 next 为 null
4.如果另一个链表还没有遍历完，则直接拼接在后面
*/

func merge2OrderlyListNode(l1, l2 *ListNode) *ListNode {
	// 哨兵节点
	guard := &ListNode{}
	cur := guard
	for l1 != nil && l2 != nil {
		if l1.Node < l2.Node {
			guard.Next = l1
			l1 = l1.Next
		} else {
			guard.Next = l2
			l2 = l2.Next
		}
		guard = guard.Next
	}

	// 指向未拼接完的链表
	if l1 != nil {
		guard.Next = l1
	}

	if l2 != nil {
		guard.Next = l2
	}

	return cur.Next
}
