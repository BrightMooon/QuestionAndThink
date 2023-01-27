package dp

import "fmt"

/* 初始化的细节
根源 dp[i][j]=dp[i-1][j-1]
1.  关于dp数组开多大的思考
    如果都是赋值数组的大小，判断条件获取的是num[i-1]会导致i=len最后一个num[i]无法进入，循环就退出了，导致结果少了一位；
    多开一位，dp数组整体后移一位，即dp[1] 和numsp[i-1](nums[0])对齐

2. 关于遍历结束的索引位置
判断条件获取的是num[i-1]会导致i=len最后一个num[i]无法进入，循环就退出了，导致结果少了一位
3. 关于遍历开始的索引位置
如果从0开始，会导致i-1为负出现空指针异常

*/

func findLength(nums1 []int, nums2 []int) int {
    //step1 dp[i][j] 以nums1[i]结尾，nums2[j]结尾构成的最长公共子数组长度
    dp:=make([][]int,len(nums1)+1)
    for index:=range dp{
        dp[index]=make([]int,len(nums2)+1)
    }
    result:=0
    //step2 确定递推关系，试一试
	//dp[i][j]=dp[i-1][j-1]+1
    
    for i:=1;i<=len(nums1);i++{
        for j:=1;j<=len(nums2);j++{
            if nums1[i-1]==nums2[j-1]{
                dp[i][j]=dp[i-1][j-1]+1
            }
            if dp[i][j]>result{
                result=dp[i][j]
            }
        }
    }
    return result
}
/*
压缩为1维
https://programmercarl.com/0718.%E6%9C%80%E9%95%BF%E9%87%8D%E5%A4%8D%E5%AD%90%E6%95%B0%E7%BB%84.html#%E6%80%9D%E8%B7%AF

1维为什么要倒叙，相加的上一个状态在左上角，不符合条件的时候，会设置为0；倒叙的时候不会影响前面，但是顺序的时候前面会被初始化为0

*/


func FindLength1D(nums1 []int, nums2 []int) int {
    dp:=make([]int,len(nums2)+1)
    result:=0
    for i:=1;i<=len(nums1);i++{
        for j:=len(nums2);j>=1;j--{
            if nums1[i-1]==nums2[j-1]{
                dp[j]=dp[j-1]+1
            }else{
                dp[j]=0
            }
            if dp[j]>result{
                result=dp[j]
            }
        }
		fmt.Printf("i=%d,%v",i,dp)
    }
    return result
}