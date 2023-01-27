package dp


func lengthOfLIS(nums []int) int {
    //dp[i] 以nums[i]结尾的情况下 nums[0...i]中，长度最长的子序列长度
    dp:=make([]int,len(nums))
    res:=0
    for i:=0;i<len(nums);i++{
        dp[i]=1
        for j:=0;j<i;j++{
            if nums[i]>nums[j]{
                dp[i]=max(dp[i],dp[j]+1)
            } 
        }
        if dp[i]  > res{
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