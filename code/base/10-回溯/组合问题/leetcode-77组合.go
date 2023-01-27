package back
func combine(n int, k int) [][]int {
    if n<k {
        return nil
    }
    path:=make([]int,0)
    res:=make([][]int,0)
    backTracking(n,k,1,path,&res)
    return res
}

func backTracking (n,k,startIndex int,path []int, res *[][]int){
    if len(path)==k{
        b:=make([]int,k)
        copy(b,path)
        *res=append(*res,b)
        return
    }
    //错误点，这里为什么要等号
    for i:=startIndex;i<=n-(k-len(path))+1;i++{
        path=append(path,i)
        backTracking(n,k,i+1,path,res)
        path=path[:len(path)-1]
    }
    return
}
