package btree

func inorderTraversal(root *TreeNode) []int {
    stack:=make([]*TreeNode,0)
    result:=make([]int,0)
    cur:=root
    for cur!=nil || len(stack)>0 {
        if cur!=nil{
            stack=append(stack,cur)
            cur=cur.Left
        }else{
            pop:=stack[len(stack)-1]
            stack=stack[:len(stack)-1]
            result=append(result,pop.Val)
            //错误点，要用从栈取出来的node 是pop 不是cur,别搞混了，不要为了格式和弄过了过程
            cur=pop.Right
        }
    }
    return result
}

