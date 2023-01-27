package dp
/*
1. 定义状态
dp[i][j]表示下标从i到j组成子串中的最长回文序列的长度
2. 转移方程
dp[i][j]=dp[i+1][j-1]+2
dp[i][j]=max(dp[i][j-1],dp[i+1][j])
3. 初始化
dp[0][0]=1
dp[i][i]=1
4. 遍历顺序
左下方，所以是i从后往前遍历，为了保证覆盖每个子问题，j=i+1开始
*/

func longestPalindromeSubseq(s string) int {
	
}