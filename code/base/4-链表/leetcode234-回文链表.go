package  link

func isPalindrome(head *ListNode) bool {
    //链表一份为2的中间节点处理
    slow := searchMidNode(head)
    back:=slow.Next
    reverseBack:=reverse(back)
    for reverseBack!=nil{
        if reverseBack.val!=head.val{
            return false
        }
        reverseBack=reverseBack.Next
        head=head.Next
    }
    return true
}

// searchMidNode 求中间的节点
func searchMidNode(head *ListNode) *ListNode {
    fast, slow := head, head
    //核心错误点 fast.next.next 大于三个节点
    //这里如果是两个节点相当于直接返回的头结点
    for fast.Next != nil && fast.Next.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    return slow
}

func reverse(head *ListNode)*ListNode{
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