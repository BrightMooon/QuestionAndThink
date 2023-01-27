package array

func numSubarrayProductLessThanK(nums []int, k int) int {
	left:=0
	product:=0
	count:=0
	for right:=0;right<len(nums);right++{
		product*=nums[right]
		for product>=k{
			product/=nums[left]
			left++
		}
		if right>=left{
			count+=right-left+1
		}
	}
	return  count
}