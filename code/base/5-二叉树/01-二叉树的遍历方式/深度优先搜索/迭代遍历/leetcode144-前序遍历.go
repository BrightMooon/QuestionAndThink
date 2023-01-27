package btree

/*
 1 先把二叉树root放入到栈中
 2 循环遍历条件，栈不为空 出栈
 3 栈顶获取，然后弹出
 4 非空节点，先入右节点，后入左节点
 5 slice是左闭右开的
 */
 func preorderTraversal(root *TreeNode) []int {
	stack:=make([]*TreeNode,0)
	result:=make([]int,0)
    //错误点，控制判断
    if root==nil{
        return result
    }
	stack=append(stack,root)
	//如果栈不为空
	for len(stack)>0{
		top:=stack[len(stack)-1]
		result=append(result, top.Val)
		stack=stack[:len(stack)-1]
		if top.Right!=nil{
			stack=append(stack, top.Right)
		}
		if top.Left!=nil{
			stack=append(stack, top.Left)
		}
	}
	return result
}