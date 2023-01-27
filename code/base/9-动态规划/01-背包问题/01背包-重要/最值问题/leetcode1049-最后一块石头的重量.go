package dp

import "fmt"

func LastStoneWeightII(stones []int) int {
    sum:=0
    for _,v:=range stones{
        sum+=v
    }
    target:=sum/2

    dp:=make([]int,target+1)
    for k:=range dp{
        dp[k]=0
    }

    for i:=0;i<len(stones);i++{
        for j:=target;j>0;j--{
            if j>=stones[i]{
                dp[j]=max(dp[j],dp[j-stones[i]]+stones[i])
				fmt.Printf("第i=%v行,第j=%v列 ,最大 %v \n",i,j,dp)
            }else{
                dp[j]=dp[j]
				fmt.Printf("第i=%v行，第j=%v列，复制上一行 %v \n",i,j,dp)
            }
        }
    }
    return sum-dp[target]-dp[target]

}

