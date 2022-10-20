package link

func middleNode(head *ListNode) *ListNode {
    fast,slow:=head,head
    //快指针走两步，慢指针走一步，当快指针走到末尾的时候，慢指针在中间位置
    //为什么边界条件是这两个，一个的话有什么问题
    //后面要用到next.next
    for fast!=nil &&fast.Next!=nil{
        fast=fast.Next.Next
        slow=slow.Next

    }
    return slow
}