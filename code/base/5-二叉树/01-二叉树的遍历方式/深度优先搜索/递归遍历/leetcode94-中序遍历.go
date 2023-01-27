package btree

func midOrderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	traversal(root, &res)
	return res
}

func traversal(root *TreeNode, rec *[]int) {
	if root == nil {
		return
	}
	traversal(root.Left, rec)
	*rec = append(*rec, root.Val)
	traversal(root.Right, rec)

}
