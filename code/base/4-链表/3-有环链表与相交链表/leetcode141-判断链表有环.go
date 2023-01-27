package link

/*
方法一
需要记忆的有环的结论：
1. 快慢指针，第一次相遇后，快指针指向head节点
2. 在第二次相遇时，相同的节点，就是相交的节点
*/
func hasCycle(head *ListNode)bool{
	//0 边界为空的判断
	//1 (这一步非常重要 )快慢指针的初始化 slow=head fast=head
	//2 loop的结束的条件 快指针到达链表的结尾
	//3 快指针每次走两步，慢指针每次走一步
	if head==nil ||head.Next==nil{
		return false
	}
	fast,slow:=head,head
	for fast!=nil && fast.Next!=nil{
		fast=fast.Next.Next
		slow=slow.Next
		if fast==slow{
			return true
		}

	}
	return false

}

/*
方法二
如果不要求时间复杂度是O(1),可以用hashSet，遍历的时候添加到hashSet里，后序存在说明有
*/

func hasCycle2(head *ListNode) bool {
	exists:=make(map[*ListNode]int)
	for head!=nil{
		if _,ok:=exists[head];ok{
			return true
		}else{
			exists[head]=1
			head=head.Next
		}
	}
    return false
}