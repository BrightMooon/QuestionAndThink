package link

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
1 双指针 pre和cur pre=cur cur=cur.next
2 提前保持next=cur.next
*/
func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	var temp *ListNode
	//错误点 判断cur不为空即可||一次循环改变一个节点的方向
	for cur != nil {
		//step-1 提前保存断开节点，并方向反转
		temp = cur.Next
		cur.Next = pre
		//step-2 双指针 整体后移动，移动的顺序先移动pre后移动cur
		pre = cur
		cur = temp
	}
	return pre
}
