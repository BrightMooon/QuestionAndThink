package array

/**
快慢指针 删除数组的中的重复元素
优质解答：https://leetcode.cn/problems/remove-duplicates-from-sorted-array/solution/kuai-man-zhi-zhen-26-shan-chu-pai-xu-shu-xf3t/
*/
func removeDuplicates(nums []int) int {
	/*
	伪代码
	1定义快慢指针fast,slow
	2[0,slow] 为不重复的元素
	3 fast遍历从第二个位置
	fast=1
	for fast<len
	4 if nums[fast]!=nums[slow]
	5 slow指针位置+1，fast指向的不重复的元素赋值给slow

	*/
	slow:=0
	for  fast:=1;fast<len(nums);fast++{
		if nums[fast]!=nums[slow]{
			slow++
			nums[slow]=nums[fast]
		}
	}
	return slow+1
}
