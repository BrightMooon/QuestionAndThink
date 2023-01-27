package dp

func findLengthOfLCIS(nums []int) int {
    /*
    dp[i] 以nums[i]元素结尾的连续递增子序列中，最长的长度
    if nums[i]>nums[i-1]
    dp[i]=dp[j]+1
    dp[i]=dp[j]
    dp[0]=1

    */
    dp:=make([]int,len(nums))
    dp[0]=1
    res:=1
    for i:=1;i<len(nums);i++{
        dp[i]=1
        for j:=0;j<i;j++{
            if nums[i]>nums[i-1]{
                dp[i]=dp[j]+1
            }
        }
        if dp[i]>res{
            res=dp[i]
        }
    }
    return res
}


//一层循环
func findLengthOfLCIS(nums []int) int {
    /*
    dp[i] 以nums[i]元素结尾的连续递增子序列中，最长的长度
    if nums[i]>nums[i-1]
    dp[i]=dp[j]+1
    dp[i]=dp[j]
    dp[0]=1

    */
    dp:=make([]int,len(nums))
    dp[0]=1
    res:=1
    for i:=1;i<len(nums);i++{
        dp[i]=1
        if nums[i]>nums[i-1]{
            dp[i]=dp[i-1]+1
        }
        if dp[i]>res{
            res=dp[i]
        }
    }
    return res
}