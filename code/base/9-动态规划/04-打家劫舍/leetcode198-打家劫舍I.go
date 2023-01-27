package dp

func rob(nums []int) int {
    /*
    dp[i] 第i天获取的最大金额
    i-1天偷了
    dp[i]=dp[i-1]
    i-1天没偷
    dp[i]=dp[i-1]+nums[i]
    */
    dp:=make([]int,len(nums))
    dp[0]=nums[0]
    if len(nums)<2{
        return nums[0]
    }
    for i:=0;i<len(nums);i++{
        dp[0]=nums[0]
        dp[1]=max(nums[0],nums[1])
        if i>=2{
            dp[i]=max(dp[i-2]+nums[i],dp[i-1])
        }
    }
    return dp[len(nums)-1]

}

func max(a,b int)int{
    if a<b{
        return b
    }
    return a
}