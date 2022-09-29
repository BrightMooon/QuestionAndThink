package array

/*
双指针从两边往中间靠拢
*/
func sortedSquares(nums []int) []int {
	size:=len(nums)
	//1 用make初始化指定长度的数组
	result:=make([]int,size)
	//2 如何从后往前给数组赋值
	k:=size-1
	i:=0;
	j:=size-1
	for i<=j{
		if nums[i]*nums[i]<nums[j]*nums[j]{
				result[k]=nums[j]*nums[j]
			j--
		}else{
			result[k]=nums[i]*nums[i]
			i++
		}
		k--
	}
	return  result
}