package dp
func fib(n int) int {
    /*
    1.下标含义
        第i个 斐波那契值的 dp[i]
    2. 什么是递推公式，当前值和前几个值的关系，体现了状态，如何由前面的状态，推到出后面的状态
        dp[i]=dp[i-1]+dp[i-2]
    3. dp数组如何初始化
        dp[0]=0 dp[1]=1
    4. 确定遍历顺序
        dp[i]=dp[i-1]+dp[i-2] 从递推公式看出，dp[i]是依赖 dp[i-1] dp[i-2],那么遍历顺序一定是从前往后遍历
    5. 举例推导递推数组
        当N为10
        dp={0,1,1,2,3,5,8,13,21,34,55}
    */
    if n< 2 {
        return n
    }
    dp:=make([]int,n+1)
    dp[0]=0 
    dp[1]=1

    for i:=2;i<=n;i++{
        dp[i]=dp[i-1]+dp[i-2]
    }
    return dp[n]
       
}