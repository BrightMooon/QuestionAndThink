package array

func MaxSubArrayLen(nums []int, k int) int {
	val2CountMap := make(map[int]int)
	val2CountMap[0] = 1
	count, sum := 0, 0
	for item := range nums {
		sum += item
		count += val2CountMap[sum-k]
		val2CountMap[sum] = val2CountMap[sum] + 1
	}
	return count
}
