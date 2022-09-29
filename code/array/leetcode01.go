package array

/**
两数之和
hashkey 方法
*/
func twoSum1(nums []int, target int) []int {
	mapNum := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		//2 判断差值作为hash key 是否存在，而不是值
		index, ok := mapNum[target-nums[i]]
		if ok {
			return []int{index, i}
		}
		//1 遍历存入map key 是nums[i],val 是i
		mapNum[nums[i]] = i
	}
	return []int{0, 0}

}

/**
两数之和
暴力枚举
最容易想到的方法是枚举数组中的每一个数 x，
寻找数组中是否存在 target - x。
当我们使用遍历整个数组的方式寻找 target - x 时，
需要注意到每一个位于 x 之前的元素都已经和 x 匹配过，因此不需要再进行匹配。
而每一个元素不能被使用两次，
！！！所以我们只需要在 x 后面的元素中寻找 target - x。！！！
*/
func twoSum2(nums []int, target int) []int {
	for i, val := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[j]+val == target {
				return []int{i, j}
			}
		}
	}
	return nil

}
