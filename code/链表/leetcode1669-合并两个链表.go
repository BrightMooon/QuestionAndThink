
package link

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
    afterB:=&ListNode{}
    beforeA:=&ListNode{}
    cur:=list1
    cur2:=list2
    index:=0
    for cur!=nil{
        if index==a-1{
            beforeA=cur
        }
        if index==b+1{
            afterB=cur
        }
        cur=cur.Next
        index++
    }
    //错误点，cur.next !=nil（cur指向最后一个节点） vs  cur!=nil（最后复制的是null） 区别
    for cur2.Next!=nil{
        cur2=cur2.Next
    }
    cur2.Next=afterB
    beforeA.Next=list2
   return list1
}