package btree

import "math"

//自顶部向下
func isBalance(root *TreeNode)bool{
	if root==nil{
		return true
	}
	left:=height(root.Left)
	right:=height(root.Right)
	return math.Abs(left-right)<=1 && isBalance(root.Left)&& isBalance(root.Right)

}

func height(root *TreeNode)int{
	if root==nil{
		return 0
	}
	left:=height(root.Left)
	right:=height(root.Right)
	return max(left,right)+1
}


//自底而上
const UNBLANCED=-1
func isBalanced(root *TreeNode)bool{
	return helper(root)!=UNBLANCED
}

func helper(root *TreeNode)int {
	if root==nil{
		return 0
	}
	left:=helper(root.Left)
	if left==UNBLANCED{
		return UNBLANCED
	}
	right:=helper(root.Right)
	if right==UNBLANCED{
		return UNBLANCED
	}
	if abs(left-right)>1{
		return UNBLANCED
	}
	return max(left,right)+1
}
func max(a,b int)int{
    if a<b{return b}
    return a
}
func abs(a int)int{
    if a>0 {
        return a
    }else{
        return -a
    }
}