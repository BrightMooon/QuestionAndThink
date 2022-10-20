package link

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy:=&ListNode{-1,nil}
    cur:=dummy
    cur1:=list1
    cur2:=list2
    var smallNode *ListNode
    if list1==nil{
        return list2
    }
    if list2==nil{
        return list1
    }
  
    //错误点1？如果同时遍历两个链表，并获取其中的节点操作
    //错误点2  cur1!=nil vs cur1.Next!=nil 终止的时候，cur1指针指向null,最后一个节点之前已经操作完成了，如果使用cur1.Next来判断，导致合并的时候丢掉最后这个节点
    //错误点3 是||还是&& 或的关系，一个截止了另一个还要继续
    for cur1!=nil || cur2!=nil{
        if cur1==nil && cur2!=nil{
            cur.Next=cur2
            break
        }
        if cur1!=nil && cur2==nil{
            cur.Next=cur1
            break
        }
        if cur1.Val<cur2.Val{
            smallNode=cur1
            cur1=cur1.Next
        }else{
            smallNode=cur2
            cur2=cur2.Next
        }
        cur.Next=smallNode
        //错误点三：注意这个创建新的节点，移动指针的过程
        cur=smallNode
    }
    return dummy.Next
}