package link
/*
1. 先快速完成快慢指针环的检测
2. 用龟兔赛跑的结论，完成第二次相遇的
*/

func detectCycle(head *ListNode) *ListNode {
    if head ==nil ||head.Next==nil{
        return nil
    }
    fast,slow:=head,head
    for fast!=nil && fast.Next!=nil{
        fast=fast.Next.Next
        slow=slow.Next
        //第一次相遇是 fast每次走两步，slow走一步
        if fast==slow{
            //第二次相遇 fast,slow每次都走一步，fast从头开始走
            fast=head
            for fast!=slow{
                 fast=fast.Next
                 slow=slow.Next
            }
            return fast
        }
    }
    return nil
}