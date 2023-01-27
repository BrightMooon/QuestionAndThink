func pivotIndex(nums []int) int {
    total,sum,result:=0,0,-1
    for _,item:=range nums{
        total+=item
    }
    for i:=0;i<len(nums);i++{
        sum+=nums[i]
        if(sum-nums[i]==total-sum){
            return i
        }
    }
    return result
}