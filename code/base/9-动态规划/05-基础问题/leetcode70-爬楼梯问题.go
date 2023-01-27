package dp
func climbStairs(n int) int {
    // 定义dp数组下标的含义 dp[i] 到达i阶的方法个数
    //错误点 数组只能开常量长度，不能开变量长度；开n+1，从0到n有n+1个元素
     dp:=make([]int,n+1)
    // 初始化 
    //错误点：初始化的时候包含1,
    if n<=1 {return n}
    dp[1]=1
    dp[2]=2
    for i:=3;i<=n;i++{
        dp[i]=dp[i-1]+dp[i-2]
    }
    return dp[n]
}
