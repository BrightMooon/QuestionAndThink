package heap_demo

func mergeKListsTwo(lists []*ListNode) *ListNode {
	if lists==nil {return nil}
	return merge(lists,0,len(lists)-1)

}

func merge(lists []*ListNode,left,right int)*ListNode{
	if left==right{return lists[left]}
	mid:=left+(right-left)/2
	l1:=merge(lists,left,mid)
	l2:=merge(lists,mid+1,right)
	return mergeTwo(l1,l2)
}
	

func mergeTwo(l1 *ListNode,l2 *ListNode) *ListNode{
	if l1==nil {return l2}
	if l2==nil {return l1}
	if l1.Val<=l2.Val{
		l1.Next=mergeTwo(l1.Next,l2)
		return l1
	}else{
		l2.Next=mergeTwo(l1,l2.Next)
		return l2
	}

}