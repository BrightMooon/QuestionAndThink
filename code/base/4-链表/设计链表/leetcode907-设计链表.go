package link
//设计链表 
type MyLinkedList struct {
    head *ListNode
    size int
}

func Constructor() MyLinkedList {
    return MyLinkedList{&ListNode{}, 0}
}

//1 查找某个节点
func (l *MyLinkedList) Get(index int) int {
    if index < 0 || index >= l.size {
        return -1
    }
    cur := l.head
    for i := 0; i <= index; i++ {
        cur = cur.Next
    }
    return cur.Val
}
//2 添加新的头结点
func (l *MyLinkedList) AddAtHead(val int) {
    l.AddAtIndex(0, val)
}

//3 添加尾部节点
func (l *MyLinkedList) AddAtTail(val int) {
    l.AddAtIndex(l.size, val)
}

//4 指定位置插入节点
func (l *MyLinkedList) AddAtIndex(index, val int) {
    if index > l.size {
        return
    }
    index = max(index, 0)
    l.size++
    pred := l.head
    for i := 0; i < index; i++ {
        pred = pred.Next
    }
    toAdd := &ListNode{val, pred.Next}
    pred.Next = toAdd
}
//5 删除指定节点
func (l *MyLinkedList) DeleteAtIndex(index int) {
    if index < 0 || index >= l.size {
        return
    }
    pred := l.head
    for i := 0; i < index; i++ {
        pred = pred.Next
    }
    pred.Next = pred.Next.Next
    l.size--
}

func max(a, b int) int {
    if b > a {
        return b
    }
    return a
}
