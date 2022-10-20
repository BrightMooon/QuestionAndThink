package dp
func canPartition(nums []int) bool {
    // i为nums里的元素nums[0-i] ，j为这些元素的和，j的范围为0-target, 表格里的格子填的是true/false
    dp:=make([][]bool,len(nums))
	sum,max:=0,0
	for _,v:=range nums{
		sum+=v
        if v>max{
            max=v
        }

	}
    target:=sum/2
    if sum%2 != 0 {
        return false
    }
    if max>target{
        return false
    }
    //初始化第一列
    for index:=range dp{
        dp[index]=make([]bool,target+1)
        dp[index][0]=true
    }
    dp[0][nums[0]]=true
	for i:=1;i<len(nums);i++{
		for j:=0;j<=target;j++{
            dp[i][j]=dp[i-1][j]
            //动态规划里的打印，会导致超时的
			//fmt.Printf("i=%v,j=%v,不放入 dp[i-1][j]=%v \n",i,j,dp[i][j])
			 /*错误点  获取的当前元素的重量
             1.nums[i] len(nums) 
             2.nums[i-1] len(nums)+1
             */
            if  j>=nums[i]{
                //如果放入当前元素，由前一个状态获取当前的状态
                dp[i][j]=dp[i-1][j]||dp[i-1][j-nums[i]]
				//fmt.Printf("i=%v,j=%v,放入 dp[i-1][j-nums[i-1]]=%v \n",i,j,dp[i][j])
            }
		}
	}
	return dp[len(nums)-1][target]
}

