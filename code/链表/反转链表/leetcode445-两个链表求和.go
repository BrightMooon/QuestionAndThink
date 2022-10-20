package link


package link

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//反转两个链表
	reverseL1:=reverse(l1)
	reverseL2:=reverse(l2)
	//反转求和结果链表
	sumListNode:=sumListNode(reverseL1,reverseL2)
	return reverse(sumListNode)
}

func sumListNode(l1 *ListNode ,l2 *ListNode )*ListNode{
//用dummy节点创建一个新的链表
dummy:=&ListNode{}
//新建一个指针指向dummy节点
sumNode:=dummy
//定义一个进位保持变量，如果进位了是新加点的值
carry,sum,temp1,temp2:=0,0,0,0
//循环迭代条件 链表1不为空或者链表2不为空
for l1!=nil||l2!=nil {
	//求和 sum=node1.val+node2.val+carry,如果为空则为0，需要加上前两位的进位树
	if l1==nil {
		temp1=0
	}else{
		temp1=l1.Val
	}
	if l2==nil {
		temp2=0
	}else{
		temp2=l2.Val
	}
	//通过sum的值判断进位carray的值，如果sum大于10需要减去10保留剩下
	sum=temp1+temp2+carry
	if sum>=10{
		sum-=10
		carry=1
	}else{
		carry=0
	}
	//构建新的节点
	newNode:=&ListNode{sum,nil}
	sumNode.Next=newNode
	//迭代新的节点
	sumNode=sumNode.Next
	//错误~~判断进位构建节点
	//正常迭代
	if l1!=nil{
		 l1=l1.Next
	}

	if l2!=nil{
		l2=l2.Next
	}
	
}
//为什么 ，循环迭代之后，判断进位构建新节点
//全部迭代之后会产生一个新的节点
if carry>0{
	sumNode.Next=&ListNode{1,nil}
}
//***错误点 这里的返回值是dummy.Next,如果返回sumNode这个指针只会返回最后一个节点
return dummy.Next
}

func reverse(head *ListNode) *ListNode{
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