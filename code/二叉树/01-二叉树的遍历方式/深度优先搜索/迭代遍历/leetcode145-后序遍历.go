package  btree
type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}
func postorderTraversal(root *TreeNode) []int {
    stack:=make([]*TreeNode,0)
    result:=make([]int,0)
    if root==nil{
        return result
    }
    stack=append(stack,root)
    for len(stack)>0{
        top:=stack[len(stack)-1]
        result=append(result,top.Val)
        stack=stack[:len(stack)-1]
        if top.Left!=nil{
            stack=append(stack,top.Left)
        }
         if top.Right!=nil{
             stack=append(stack,top.Right)
        }
    }
    reverse(result)
    return result 
}

func reverse(array []int){
    l:=0
    r:=len(array)-1
    for l<r{
        array[l],array[r]=array[r],array[l]
        l++
        r--
    }
}



