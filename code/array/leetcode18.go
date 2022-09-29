package array

import "sort"

/**
四数和
难点是去重和剪枝
*/
func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	result := [][]int{}
	if n < 4 {
		return nil
	}
	sort.Ints(nums)
	//注意这个边界
	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			//这里不是大于0
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l := j + 1
			r := n - 1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				} else if sum < target {
					l++
				} else {
					r--
				}

			}

		}

	}

	return result
}
