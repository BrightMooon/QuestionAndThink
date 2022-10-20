package link

//1如何用双指针找到倒数第N个节点
//2 快指针 走N+1步，慢指针开始移动，为了让slow指向要删除节点的前一个节点

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 type ListNode struct {
	Val  int
	Next *ListNode
}


 func removeNthFromEnd(head *ListNode, n int) *ListNode {
    //1 定义快慢指针指向dummpy节点
    dummy:=&ListNode{0,head}
    //2 定义操作指针
    fast,slow,cur:=dummy,dummy,dummy
    //3 让fast指针先移动n-1步
    for i:=0;i<n;i++{
        fast=fast.Next
    }
    /*这是核心点，带入例子测试
    a 当fast.next判断是否nil时，
    b  fast指针已经移动了链表最后一个节点，
    c slow指针也移动到待删除节点的妻哪一个节点*/
    for fast.Next!=nil{
        slow=slow.Next
        fast=fast.Next
    }
    //4 开始删除操作
    slow.Next=slow.Next.Next
    return cur.Next

}
