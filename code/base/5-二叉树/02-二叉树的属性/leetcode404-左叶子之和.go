package btree
type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode

}
/*
解题思路及伪代码：
1. 一般二叉树，递归过程中还有回溯的过程，例如走一个左方向的分支走到头了，那么要调头，在走右分支。
*/
func sumOfLeftLeaves(root *TreeNode) int {
	return 0

}