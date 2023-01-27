package bp



func maxSubArray(nums []int) int {
    dp:=make([]int,len(nums))
    //错误点，初始化
    dp[0]=nums[0]
    //错误点 res不是初始化为0，而是初始化为dp[0]
    res:=dp[0]
    for i:=1;i<len(nums);i++{
        dp[i]=max(nums[i],dp[i-1]+nums[i])
        if dp[i]>res{
            res=dp[i]
        }
    }
    return res
}

func max(a,b int)int{
    if a<b{
        return b
    }
    return a
}