package dp

func rob(nums []int) int {
    if len(nums)==1{
        return nums[0]
    }
    return max(subRob(nums[0:len(nums)-1]),subRob(nums[1:len(nums)]))
}

func subRob(nums []int)int{

    dp:=make([]int,len(nums))
    size:=len(nums)
    if size==0{
        return 0
    }
    if size==1{
        return nums[0]
    }
    if size==2{
        return max(nums[0],nums[1])
    }
    dp[0]=nums[0]
    dp[1]=max(nums[0],nums[1])
    for i:=2;i<len(nums);i++{
        dp[i]=max(dp[i-2]+nums[i],dp[i-1])
    }
    return dp[size-1]
}
func max(a,b int )int{
    if a<b{
        return b
    }
    return a
}