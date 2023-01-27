package dp
/*
给定一个字符串 s 和一个字符串 t ，计算在 s 的子序列中 t 出现的个数。

字符串的一个 子序列 是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。（例如，"ACE" 是 "ABCDE" 的一个子序列，而 "AEC" 不是）

题目数据保证答案符合 32 位带符号整数范围。

 
示例 1：

输入：s = "rabbbit", t = "rabbit"
输出：3
解释：
如下图所示, 有 3 种可以从 s 中得到 "rabbit" 的方案。
rabbbit
rabbbit
rabbbit
示例 2：

输入：s = "babgbag", t = "bag"
输出：5
解释：
如下图所示, 有 5 种可以从 s 中得到 "bag" 的方案。 
babgbag
babgbag
babgbag
babgbag
babgbag
 

提示：

0 <= s.length, t.length <= 1000
s 和 t 由英文字母组成


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/distinct-subsequences
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
1.每一个dp单元格都是当前状态下的结果值，也是下一个状态可以以来的结果
2. 为什么是dp[i-1][j] 应为要使用的是j的全部，是不是i的子序列
3. 如果使用s[i-1]匹配,则直接使用左上角的dp值，因为多家一个字符，对单元格中的子序列数目是没变化的

*/
func numDistinct(s string, t string) int {
    dp:=make([][]int,len(s)+1)
    for index:=range dp{
        dp[index]=make([]int,len(t)+1)
    }

     for i:=0;i<=len(s);i++{
        dp[i][0]=1
    }
    for j:=1;j<=len(t);j++{
        dp[0][j]=0
    }
    
   
    for i:=1;i<=len(s);i++{
        for j:=1;j<=len(t);j++{
            //dp递推公式
            if s[i-1]==t[j-1]{
                //如果使用s[i-1]匹配||如果不使用s[i-1]匹配
                dp[i][j]=dp[i-1][j-1]+dp[i-1][j]
            }else{
                //为什么是dp[i-1][j] 应为要使用的是j的全部，是不是i的子序列
                dp[i][j]=dp[i-1][j]
            }
        }
    }
    return dp[len(s)][len(t)]
}