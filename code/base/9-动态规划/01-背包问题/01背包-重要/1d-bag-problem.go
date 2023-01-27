package dp

import "fmt"

func OneDBagProblem() int{
	weight:=[]int{1,3,4}
	value:=[]int{15,20,30}
	bagWeight:=4
	//错误点：初始化  这里 是weight 还是bagWeigeht
	dp:=make([]int,bagWeight+1)
	for i:=0;i<len(weight);i++{
		for j:=bagWeight;j>=weight[i];j--{
			dp[j]=max(dp[j],dp[j-weight[i]]+value[i])
			fmt.Println(dp[j])
		}
	}
	return dp[bagWeight]

}