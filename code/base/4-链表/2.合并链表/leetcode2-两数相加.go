package link

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 /*
1. 易错点 
 1.1 忘记最后的carry==1 时的进位
 1.2 dummy节点，就是申明两个节点，dummy,cur 操作节点cur,最后返回dummy
 1.3 链表迭代的条件，当前是cur!=nil,才会去cur.Next
 1.4 carry要在sum变化之前获取


2. 难点
 不相等的链表就理解为用0补齐

3. 着手点
    3.1 为新的链表新建一个dummy节点，
    3.2 loop条件是||
    3.3 % / 计算sum和carry
4. 大体思路
迭代 cur=cur.next l1=l1.next l2=l2.next
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy:=&ListNode{Val:0}
    cur:=dummy
    carry:=0
    sum:=0
    for l1!=nil || l2!=nil {
        v1:=0
        v2:=0
        if l1!=nil{
           v1=l1.Val 
        }
        if l2!=nil{
           v2=l2.Val 
        }
        sum=v1+v2+carry
        //错误点 先去进位，不然sum被覆盖了
        carry=sum/10
        sum=sum%10
        cur.Next=&ListNode{Val:sum}
        cur=cur.Next
        //错误点：注意这里是l1!=nil ,因为要取next
        if l1!=nil{
            l1=l1.Next
        }
        if l2!=nil{
            l2=l2.Next
        }
        if carry==1{
            cur.Next=&ListNode{Val:1}
        }
    }
    return dummy.Next
}