package  btree
type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}
func BackorderTraversal(root *TreeNode)[]int{
	result:=make([]int,0)
	traversal(root,&result)
	return result
}

func traversal(root *TreeNode ,result *[]int){
	if root==nil{
		return
	}
	traversal(root.Left,result)
	traversal(root.Right,result)
	*result = append(*result, root.Val)
}



