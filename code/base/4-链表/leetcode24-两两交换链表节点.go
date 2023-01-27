package link
type ListNode struct{
	val int
	Next *ListNode
}
//画图举例子 dummy->1->2->3->4->5->null
//要确定三个节点，两个节点前的一个节点
func swapPairs(head *ListNode) *ListNode {
	dummy:=&ListNode{-1,head}
	cur:=dummy
	for cur.Next!=nil&&cur.Next.Next!=nil{
		temp:=cur.Next
		temp1:=cur.Next.Next.Next
		cur.Next=cur.Next.Next
		cur.Next.Next=temp
		temp.Next=temp1
		cur=cur.Next.Next
	}
	return dummy.Next
}