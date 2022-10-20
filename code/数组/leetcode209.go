package array

//滑动窗口，同向双指针

//长度最小的子数组
func minSubArrayLen(target int, nums []int) int {
	left := 0
	sum := 0
	result := len(nums) + 1
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum > target {
			subLen := right - left + 1
			if result > subLen {
				result = subLen
			}
			sum -= nums[left]
			left++
		}
	}
	return result
}
