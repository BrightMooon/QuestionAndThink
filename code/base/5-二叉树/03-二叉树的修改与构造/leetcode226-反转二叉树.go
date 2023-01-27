package btree

import "fmt"
type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}
func invertTree(root *TreeNode) *TreeNode {
    if root==nil{
        return root
    }
    //1 先交换左右节点，然反转左子树，再反转右子树
    swap(root)
    invertTree(root.Left)
   	invertTree(root.Right)
	fmt.Printf("btree:%v \n",root)
    return root
}


func swap(node *TreeNode){
    node.Left,node.Right=node.Right,node.Left
}