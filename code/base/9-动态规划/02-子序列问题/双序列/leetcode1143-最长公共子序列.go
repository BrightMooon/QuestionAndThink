package dp 

func longestCommonSubsequence(text1 string, text2 string) int {
    //i-1 j-1
    //if 相等 dp[i][j]=dp[i-1][j-1]+1
    // else max(dp[i-1][j],dp[i][j-1])
    dp:=make([][]int,len(text1)+1)
    for index:=range dp{
        dp[index]=make([]int,len(text2)+1)
    }
    for i:=1;i<=len(text1);i++{
        for j:=1;j<=len(text2);j++{
            if text1[i-1]==text2[j-1]{
                dp[i][j]=dp[i-1][j-1]+1
            }else{
                dp[i][j]=max(dp[i-1][j],dp[i][j-1])
            }
        }
    }
    return dp[len(text1)][len(text2)]
}
func max(a,b int)int {
    if a<b{
        return b
    }
    return a
}