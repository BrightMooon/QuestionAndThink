package btree
type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}
func rightSideView(root *TreeNode) []int {
    queue:=make([]*TreeNode,0)
	result:=make([]int,0)
	//空判断不要忘了
    if root==nil{
        return result
    }
	//要把根节点先放入数据结构队列里面然后开始处理
    queue=append(queue,root)
	for len(queue)>0{
		size:=len(queue)
		curLevel:=make([]int,0)
		for i:=0;i<size;i++{
			cur:=queue[0]
			queue=queue[1:]
			if i==size-1{
				result=append(result,cur.Val)
			}
			if cur.Left!=nil{
				queue=append(queue, cur.Left)
			}
			if cur.Right!=nil{
				queue=append(queue, cur.Right)
			}
		}
		//记得吧每层的结果放到最后的结果
		result=append(result, curLevel...)
	}
	return result
}