package array

import (
	"sort"
)

/**
三数和
1 如何去除重复解是三数和的难点
2
*/
func ThreeSum(nums []int) [][]int {

	//1 基本边界条件判断
	//长度n 小于3 || nums==nil  直接返回nil

	//2 对数组排序 才能使用双指针从左右两边逼近

	//3 排序后遍历
	//i=0 大于0 循环中断
	//对于重复元素直接跳过 i>0 && nums[i]==nums[i-1]
	//左指针L=i+1  右指针R=n-1,当L<R 开始第二次循环：
	// sum = nums[i]+ums[l]+nums[r]
	// if sum =0
	//添加到结果
	// if nums[l]=nums[l+1] l++
	// if nums[r]=nums[r-1] r--
	// if sum>0 r --
	// if  sum<0 l++
	result := [][]int{}
	n := len(nums)
	if n < 3 {
		return nil
	}
	sort.Ints(nums)
	for i := range nums {
		if nums[i] > 0 {
			break
		}
		//去重逻辑
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := n - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				//去重逻辑
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if sum < 0 {
				l++
			} else {
				r--
			}

		}

	}

	return result
}
