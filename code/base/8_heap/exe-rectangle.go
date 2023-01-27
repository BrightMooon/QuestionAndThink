package heap_demo

type Rectangle struct{
	width int
	height int
}

func (rec *Rectangle) Area() int{
	return rec.width * rec.height
}

type RecHeap [] Rectangle

func (recHeap RecHeap) Len() int{
	return len(recHeap)
}

func (recHeap RecHeap)Swap(i,j int){
	recHeap[i],recHeap[j]=recHeap[j],recHeap[i]
}

func (recHeap RecHeap)Less(i,j int)bool{
	return recHeap[i].Area()<recHeap[j].Area()
}

func (recHeap *RecHeap)Push(h interface{}){
	*recHeap=append(*recHeap,h.(Rectangle))
}

func (recHeap  *RecHeap)Pop()(x interface{}){
	n:=len(*recHeap)
	x=(*recHeap)[n-1]
	*recHeap=(*recHeap)[:n-1]
	return x
}

