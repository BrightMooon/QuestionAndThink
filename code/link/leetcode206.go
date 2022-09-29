package link

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
key:
1 反转 当前节点的下一个 next指向pre
2 更新双指针 pre=cur cur=cur.next
*/
func reverseList(head *ListNode) *ListNode {
	var temp *ListNode
	var pre *ListNode
	cur := head
	//错误点 判断cur不为空即可
	for cur.Next != nil {
		temp = cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}
