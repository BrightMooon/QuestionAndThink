package link

/*移除链表元素
给你一个链表的头节点 head 和一个整数 val
请你删除链表中所有满足 Node.val == val 的节点，
并返回 新的头节点
*/

/*
key-1:特殊情况，如果要删除的第头头结点怎么处理
key-2:要删除一个元素，必须要知道他的上一个元素

*/
func RemoveElements(head *ListNode, val int) *ListNode {
	//这里容易犯错，是&ListNode{}；如果没有指针类型会导致未删除任何元素
	dummyHead := ListNode{}
	cur := dummyHead
	dummyHead.Next = head
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = *cur.Next
		}
	}
	return dummyHead.Next
}
