package link

//相交链表
/*
 理解上的盲区
 1. 相交是只相交一个节点吗？不是的，每个相交点只有一个next指针，所以后面的后续节点都是一样的
 2. 什么时候需要一个游动指针去代替头结点,head,cur都是指向的链表首结点的地址，直接用head，迭代之后导致head节点被覆盖

*/

type ListNode struct{
	Val int
	Next *ListNode
}
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	len1,len2:=0,0
	l1:=headA
	l2:=headB
	for l1!=nil{
		len1++
		l1=l1.Next
	}
	for l2!=nil{
		len2++
		l2=l2.Next
	}

	gap:=len1-len2
	if gap>0{
		//易错点 两个指针都指向了nil 要全部赋值为head
		l1=headA
		l2=headB
		for i:=gap;i>0;i--{
			l1=l1.Next
		}
		for len2>0{
			if l1==l2{
				return l1
			}
			l1=l1.Next
			l2=l2.Next
			len2--
		}

	}else{
		l1=headA
		l2=headB
		for i:=gap;i<0;i++{
			l2=l2.Next
		}
		for len1>0{
			if l1==l2{
				return l1
			}
			l1=l1.Next
			l2=l2.Next
			len1--
		}


	}
	return nil

	
	
}

