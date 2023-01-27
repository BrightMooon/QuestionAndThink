package btree


  type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }
 
 func balanceBST(root *TreeNode) *TreeNode {
    res:=make([]int,0)
    midTraversal(root,&res)
    var build func(nums[] int,left,right int) *TreeNode
    build =func(nums[] int,left,right int) *TreeNode{
		if left>right{
			return nil
		}
		mid:=left+(right-left)>>1
		root:=&TreeNode{Val: nums[mid]}
		root.Left=build(nums,left,mid-1)
		root.Right=build(nums,mid+1,right)
		return root
	}
    return build(res,0,len(res)-1)
}

func midTraversal(root *TreeNode, res *[]int){
    if root==nil{
        return 
    }
    //盲区：中序遍历不需要接收值
    midTraversal(root.Left,res)
    *res=append(*res,root.Val)
    midTraversal(root.Right,res)
}