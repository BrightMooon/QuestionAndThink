package dp

func TwoDBagProblem()int{
	weight:=[]int{1,3,4}
	value:=[]int{15,20,30}

	bagWeight:=4
	//二维dp数组初始化
	dp:=make([][]int,len(weight)+1)
	//错误点 这里初始化的是行的长度，即背包的总重量
	for i:=range dp{
		dp[i]=make([]int,bagWeight+1)
	}

	//初始化
	for j:=bagWeight;j>=weight[0];j--{
		dp[0][j]=dp[0][j-weight[0]]+value[0]
	}

	//双重循环

	for i:=1;i<len(weight);i++{
		//等于背包总容量
		for j:=0;j<=bagWeight;j++{
			if j< weight[i]{
				dp[i][j]= dp[i-1][j]
			}else{
				dp[i][j]=max(dp[i-1][j],value[i]+dp[i-1][j-weight[i]])
			}
		}
	}
	return dp[len(weight)-1][bagWeight]
}

func max (a,b int )int {
	if a>b{
		return a 
	}

	return b
}