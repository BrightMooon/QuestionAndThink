package link
type ListNode struct{
	Val int
	Next *ListNode
}

/*
1. 思路
快慢指针，快的先走n步，慢的再走，当fast==nil的时候，slow等于倒数第K个节点
2. 难点
+ 要拿到前一个节点，才能删除倒数第K个节点，fast先走n+1步
+ dummy节点操作
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy:=&ListNode{}
	dummy.Next=head
	fast,slow:=dummy,dummy
	i:=0
	for fast!=nil {
		if i==n+1{
			slow=slow.Next
		}
		fast=fast.Next
	}
	after:=slow.Next.Next
	slow.Next=after
	return dummy.Next
}