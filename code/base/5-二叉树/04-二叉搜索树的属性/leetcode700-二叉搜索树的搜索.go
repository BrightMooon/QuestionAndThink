package btree
type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}
/*
二叉搜索树如何遍历
1. result:=bstSearch(root.Left,val)
*/
func searchBST(root *TreeNode, val int) *TreeNode {
	var result *TreeNode
	if root==nil || root.Val==val{
		return root
	}
	if root.Val<val{
		result=searchBST(root.Right,val)
	}
	if root.Val>val{
		result=searchBST(root.Left,val)
	}

	return result
}

//迭代解法
func searchBSTLoop(root *TreeNode, val int) *TreeNode {
	for root!=nil{
        if root.Val>val{
            root=root.Left
        }else if root.Val<val{
            root=root.Right
        }else{
            return root
        }
    }
	return nil
}