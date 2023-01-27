package link

 func reorderList(head *ListNode)  {
    dummy:=&ListNode{}
    dummy.Next=head
    //1 寻找链表中点 
    fast,slow:=dummy,dummy
    for fast!=nil && fast.Next!=nil{
        fast=fast.Next.Next
        slow=slow.Next
    }
    //错误点 要找到slow的上一个节点，并把next值为空,这种取法会导致，后面的链表比前面的少一个节点
    l2:=slow.Next
    slow.Next=nil
    //2 链表逆序 
    reverseL2:=reverse(l2)
    //3 合并链表
    l1:=head
    //错误点 这里为什么要用&&，只要与一个为空之前，内部已经都处理结束额
    for l1!=nil && reverseL2!=nil{
        //错误点二 如何解决一个长一个端的问题（两个链表相差一个节点）
        /*
        1 提前保持两个链表next节点
        2. l1->l2; l1—>temp1;  l2->l1  ;l2->temp2
        */
        temp1:=l1.Next
        temp2:=reverseL2.Next
        
        l1.Next=reverseL2
        l1=temp1
        reverseL2.Next=l1
        reverseL2=temp2
    }
}
func reverse(head *ListNode) *ListNode{
    //错误 pre:=&ListNode{} 会导致该节点为0
    var pre *ListNode
    cur:=head
    for cur!=nil{
        next:=cur.Next
        cur.Next=pre
        pre=cur
        cur=next
    }
    return pre
}