package main

import (
	"fmt"
)



func main() {
	// 查缺补漏 二维数组的初始化
	// bool int?
	dp:=make([][]bool,5)
	for item:=range dp{
		dp[item]=make([]bool, 2)
		dp[item][0]=true
	}
	fmt.Println(dp)
	
}


