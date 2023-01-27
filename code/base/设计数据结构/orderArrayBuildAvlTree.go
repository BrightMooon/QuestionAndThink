package design

type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}
func buildAVLTree(nums []int) *TreeNode{
	var build func(nums[] int,left,right int) *TreeNode
	build=func(nums[] int,left,right int) *TreeNode{
		if left>right{
			return nil
		}
		mid:=left+(right-left)>>1
		root:=&TreeNode{Val: nums[mid]}
		root.Left=build(nums,left,mid-1)
		root.Right=build(nums,mid+1,right)
		return root
	}
	return build(nums,0,len(nums)-1)
}