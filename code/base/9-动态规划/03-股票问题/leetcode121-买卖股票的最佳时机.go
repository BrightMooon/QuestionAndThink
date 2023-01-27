package dp
func maxProfit(prices []int) int {
	/*
	第一次思考
	dp[i] 在第i天获取的最大利润
	第i天卖出
	dp[i]=dp[i-1]+prices[i]
	第i天买入
	dp[i]=dp[i-1]
	dp[0]=0
	1. 盲区 如何表示买入和卖出
	*/
	
	/*
	第二次思考
	+盲区
	第二个维度，持有/不持有
	+ 
	dp[i][0] 第i天持有股票获取的最大利润 
	dp[i][1] 第i天不持有股票获取的最大利润
	+
	如何表示今天持有股票获取的最大利润，确定递推关系
	dp[i][0] 持有股票可以由两个状态推出
	- 第i天买入股票 -prices[i]
	- 第i天之前就持有了股票  dp[i-1][0] (这里会迷惑到底是那天持有的，这个可以一直递推到第1天来获取)
	- dp[i][0]=max(dp[i-1][0],-prices[i])
	+ 
	也由两个状态推出
	- 第i天不持有
	dp[i][1]= prices[i]+dp[i-1][0]
	- 第i天之前就不持有了
	dp[i][1]=dp[i-1][1]
	
	+ 初始化
	
	dp[0][0]=-prices[0]
	dp[0][1]=0
	
	+ 遍历顺序
	for i:=0;i<len;i++{
	
	}
	
	*/
	dp:=make([][]int,len(prices))
	for index:=range dp{
		dp[index]=make([]int,2)
	}
	dp[0][0]=-prices[0]
	dp[0][1]=0
	 for i:=1;i<len(prices);i++{
		 dp[i][0]=max(-prices[i],dp[i-1][0])
		 dp[i][1]=max(dp[i-1][0]+prices[i],dp[i-1][1])
	 }
	 return dp[len(prices)-1][1]
	}
	
	func max(a,b int)int{
		if a<b{return b}
		return a
	}
	
	func min(a,b int)int{
		if a<b{return a}
		return b
	}