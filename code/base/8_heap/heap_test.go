package heap_demo

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestXxx(t  *testing.T){
	hp:=&RecHeap{}
	for i:=2;i<6;i++{
		*hp=append(*hp,Rectangle{i,i})
	}
	fmt.Println("原始slice:",hp)
	heap.Init(hp)
	heap.Push(hp,Rectangle{100,10})
	fmt.Println("top元素",(*hp)[0])
	fmt.Println("删除并返回最后一个:",heap.Pop(hp))
	fmt.Println("最终slice",hp)
}

