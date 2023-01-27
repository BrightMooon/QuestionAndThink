package bp

func countSubstrings(s string) int {

	//左闭右闭
	//区间范围[i,j]的子串是否是回文子串
		dp:=make([][]bool,len(s))
		count:=0
		size:=len(s)
		for index:=range dp{
			dp[index]=make([]bool,len(s))
		}
		for i:=size-1;i>=0;i--{
			for j:=i;j<size;j++{
				if s[i]==s[j]{
					if j-i<=2{
					   dp[i][j]=true
					}else{
						dp[i][j]=dp[i+1][j-1]
					}	
				}else{
					continue
				}
				if dp[i][j]{
					count++
				}
			}
		}
		return count
	
	
	}