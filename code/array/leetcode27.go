package array

/**
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。

不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/remove-element
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
一 暴力解法双重循环+整体前移+两个变量自减
 */
func removeElement(nums []int, val int) int {
	var size = len(nums)
	for i := 0; i < size; i++ {
		if nums[i] == val {
			//i后面的元素整体前移,所以j赋初值i+1
			for j := i; j < size; j++ {
				nums[j-1]=nums[j]
			}
			//i后面的元素整体前移，如果不减1 会漏掉刚覆盖来的新元素； 循环++移动一位后要是新来的元素，所以要提前--
			i--;
			//数组的长度减1
			size--;

		}
	}
	return size
}


