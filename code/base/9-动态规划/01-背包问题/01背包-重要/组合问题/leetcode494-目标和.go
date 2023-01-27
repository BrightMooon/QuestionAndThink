package dp

import "fmt"

func findTargetSumWays(nums []int, target int) int {
    sum := 0
    for _, v := range nums {
        sum += v
    }
    diff := sum - target
    if diff < 0 || diff%2 == 1 {
        return 0
    }
    n, neg := len(nums), diff/2
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, neg+1)
    }
    dp[0][0] = 1
    for i:=1;i<=len(nums);i++ {
        for j := 0; j <= neg; j++ {
            dp[i][j] = dp[i-1][j]
            if j >= nums[i-1] {
                dp[i][j] += dp[i-1][j-nums[i-1]]
            }
			fmt.Printf("dp[%v][%v]=%v \n",i,j,dp[i][j])
			fmt.Printf("total=%v \n",dp)
        }
    }
	fmt.Printf("end=%v \n",dp[n][neg])
    return dp[n][neg]
}

/**一维
// func findTargetSumWays(nums []int, target int) int {
//     /*
//     正数和P
//     负数和q
//     p-q=target
//     p+q=sum
//     p=(target+sum)/2
//     */
//     sum,s:=0,0
//     for _,v:=range nums{
//         sum+=v
//     }
//     s=(target+sum)/2
//     if (target+sum)%2!=0{
//         return 0
//     }
//     if math.Abs(float64(target)) > math.Abs(float64(sum)){
//         return 0
//     } 

// //从nums中选择数字，可以组成容量为j的表达式数量 为dp[j]
//     dp:=make([]int,s+1)
//     //为什么是1
//     dp[0]=1
//     for i:=0;i<len(nums);i++{
//         //j的起始为什么是nums[i], 循环有没有等号
//         for j:=s;j>=nums[i];j--{
//             dp[j]+=dp[j-nums[i]]
//         }
//     }
//     return dp[s]