package btree

func levelOrderBottom(root *TreeNode) [][]int {
    queue:=make([]*TreeNode,0)
    result:=make([][]int,0)
    if root==nil{
        return result
    }
    queue=append(queue,root)
    for len(queue)>0{
        size:=len(queue)
        curLevel:=make([]int,0)
        for i:=0;i<size;i++{
            cur:=queue[0]
            queue=queue[1:]
            curLevel=append(curLevel,cur.Val)
            if cur.Left!=nil{
                queue=append(queue,cur.Left)
            }
            if cur.Right!=nil{
                queue=append(queue,cur.Right)
            }
        }
        result=append(result,curLevel)
    }
    reverse(result)
    return result
}

func reverse(array [][]int){
    l:=0
    //注意这里长度，右指针的下标是长度-1
    r:=len(array)-1
    for l<r{
        //注意这里如何交换二位数组和一维是一样的
        array[l],array[r]= array[r],array[l]
        l++
        r--
    }
}