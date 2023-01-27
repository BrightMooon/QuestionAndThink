package dp
/*
给定两个单词 word1 和 word2 ，返回使得 word1 和  word2 相同所需的最小步数。
每步 可以删除任意一个字符串中的一个字符。

示例 1：

输入: word1 = "sea", word2 = "eat"
输出: 2
解释: 第一步将 "sea" 变为 "ea" ，第二步将 "eat "变为 "ea"
示例  2:

输入：word1 = "leetcode", word2 = "etco"
输出：4
 
提示：

1 <= word1.length, word2.length <= 500
word1 和 word2 只包含小写英文字母

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/delete-operation-for-two-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
前置的重要知识点
1. 我们使用的dp是从0到i-1和0到j-1,dp的下标1对应的word1 word2的下标0
2. 在二维表格中dp[i][j]对应的同行同列的元素是word1[i-1],word2[j-1]
3. 在二维表格中当前元素的左上角对应的是dp[i-1][j-1],同行左边元素是dp[i][j-1]，
同列上边元素是dp[i-1][j]
4. 关于初始化，其中一个元素为空，另一个要把所有的元素都删除了才能形成相同的
*/
func minDistance(word1 string, word2 string) int {
    dp:=make([][]int,len(word1)+1)
    for index:=range dp{
        dp[index]=make([]int,len(word2)+1)
    }
    for i:=0;i<=len(word1);i++{
        dp[i][0]=i
    }
    for j:=0;j<=len(word2);j++{
        dp[0][j]=j
    }
    for i:=1;i<=len(word1);i++{
        for j:=1;j<=len(word2);j++{
            //dp递推公式
            if word1[i-1]==word2[j-1]{
                dp[i][j]=dp[i-1][j-1]
            }else{
                dp[i][j]=min((dp[i-1][j]+1),(dp[i][j-1]+1))
            }
        }
    }
    return dp[len(word1)][len(word2)]
}
func min(a,b int)int{
    if a<b{
        return a
    }
    return b
}