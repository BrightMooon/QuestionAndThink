package btree

func preorderTraversal(root *TreeNode) []int {
	result:=make([]int,0)
	traversal(root,&result)
	return result
}

func traversal (root *TreeNode,result *[]int){
	if root ==nil{
		return
	}
	//关于slice易错点，指针才能修改外面的值；直接传递就是值传递
	*result=append(*result,root.Val)
	traversal(root.Left,result)
	traversal(root.Right,result)
}