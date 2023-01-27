package btree
type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}
/*
1. 既在左子树找p和q，又在右子树找p和q。 
2. 只要找到p或q的其中一个就可以返回，
    如果left和right都不为空，说明左右子树各找到p和q当中的一个，那么p和q在root的两侧。
    如果left为空，说明p,q在右子树。 
    如果right为空，说明p,q在左子树。 
    left和right都为空，说明找不到。 
    感谢感谢，第一次看懂了。 个人觉得，这题最后的返回条件有没有写清楚很影响阅读体验。 贴一个详细一点的。
 */

 // 后序遍历（左右中）就是天然的回溯过程，可以根据左右子树的返回值，来处理中节点的逻辑。
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root==nil ||root==q||root==q{
		return root
	}
	left:=lowestCommonAncestor(root.Left,p,q)
	right:=lowestCommonAncestor(root.Right,p,q)
	if left==nil{
		return right
	}
	 if right==nil {
		return left
	}
	if left!=nil && right!=nil{
		return root
	}
	return nil
}