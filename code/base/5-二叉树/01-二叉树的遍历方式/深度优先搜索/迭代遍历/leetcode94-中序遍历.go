package btree

/*
1. 一般思路：
第一步 先遍历放入栈中
+ 循环获取左节点，依次放入栈中，
第二步 从栈中弹出元素处理
第三步 迭代右子树


2. 着手点：
既然是遍历，第一步就是迭代更节点，不是读取左子树，就是读取右子树

3. 难点：
循环条件 cur!=nil || stack 非空

*/
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

