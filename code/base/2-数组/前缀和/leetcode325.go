package array

func MaxSubArrayLen(nums []int, k int) int {
	val2CountMap := make(map[int]int)
	val2CountMap[0] = -1
	count, sum := 0, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		_, ok := val2CountMap[sum]
		if !ok {
			val2CountMap[sum] = i
		}
		kLen, okk := val2CountMap[sum-k]
		if okk {
			if i-kLen > count {
				count = -i - kLen
			}
		}
	}
	return count
}
