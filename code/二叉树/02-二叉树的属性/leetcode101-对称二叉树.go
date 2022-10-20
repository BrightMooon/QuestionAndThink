package btree

/*
对称二叉树关键点：
1. 左右子树是相互反转的，也就是左子树的外侧左节点==右子树的外侧右节点，左子树的内侧右节点等于右子树的内侧左节点
2. 要采用什么遍历方式，后序遍历，后序可以先拿到左右节点做比较，然后回到根节点
*/
func isSymmetric(root *TreeNode) bool {
	return true
}

