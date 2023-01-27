package top

//基础模型其实就是：
//在一维数组中对每一个数找到第一个比自己小的元素。
//这类“在一维数组中找第一个满足某种条件的数”的场景就是典型的单调栈应用场景



func largestRectangleArea(heights []int) int {
	res:=0
	n:=len(heights)
	stack:=make([]int,0)
	hArray:=make([]int,n+2)
	hArray[0],hArray[n+1]=0,0
	for i:=0;i<n;i++{
		hArray[i+1]=heights[i]
	}

	for i:=0;i<len(hArray);i++{
		for len(stack)>0 && hArray[i] < hArray[stack[len(stack)-1]]{
			cur:=stack[len(stack)-1]
			stack=stack[:len(stack)-1]
			curHeight:=hArray[cur]
			//栈顶元素已经弹出，新的栈顶
			left:=stack[len(stack)-1]
			right:=i
			curWidth:=right-left-1
			res=max(res,curWidth*curHeight)
		}
		stack=append(stack, i)
	}
	return res
}

func max(a,b int)int{
	if a<b {return b}
	return a
}