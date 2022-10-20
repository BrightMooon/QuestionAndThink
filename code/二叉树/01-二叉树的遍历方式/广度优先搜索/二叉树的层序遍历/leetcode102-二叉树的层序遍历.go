package btree

type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}
func levelOrder(root *TreeNode) [][]int {
	queue:=make([]*TreeNode,0)
	result:=make([][]int,0)
	queue=append(queue, root)
	//1 当queue不为空开始循环遍历 queue!=nil
	for len(queue)>0{
		//2 获取queue的长度 size
		size:=len(queue)
		//3 新建保存当前层元素的array
		currentLvel:=make([]int,0)
	//4 根据size遍历取出queue的元素,这个size控制了，单层的元素
		for i:=0;i<size;i++{
			//5 处理出队的元素，保存到结果里
			node:=queue[0]
			queue=queue[1:]
			currentLvel=append(currentLvel,node.Val)
			//6 将当前节点的元素放入到队列中
			if node.Left!=nil{
				queue=append(queue, node.Left)
			}
			if node.Right!=nil{
				queue=append(queue, node.Right)
			}
		}
		result=append(result, currentLvel)
	}
	return result
}