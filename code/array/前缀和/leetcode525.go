package array

//连续数组
/*
二进制数组，寻找0,1 数量相同最长连续子数组
key-1 hashMap key:前缀和 val:索引
key-2 if sum[j]==sum[i](j<i), sum[j+1....i] =0 len(i-j-1+1)=i-j
key-3 map[0]=-1
*/
func findMaxLength(nums []int) int {
	valueIndexMap:=make(map[int]int)
	valueIndexMap[0]=-1
	result,sum:=0,0
	for i:=0;i<len(nums);i++{
		if nums[i]==0{
			sum--
		}else{
			sum++
		}
		if index,ok:=valueIndexMap[sum];ok{
			if i-index>result{
				result=i-index
			}
		}else{
			valueIndexMap[sum]=i
		}
		
	}
	return result 
}
