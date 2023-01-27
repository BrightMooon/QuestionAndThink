package btree

import "strconv"

  type TreeNode struct {
     Val int
      Left *TreeNode
     Right *TreeNode
 }

 //三毒，困弃躁
 func binaryTreePaths(root *TreeNode) []string {
    res:=make([]string,0)
    path:=""
    dfs(root,path,&res)
    return res
}

//前序遍历
func dfs(root *TreeNode,path string ,res *[]string){
    if root ==nil{
        return 
    }
    if root.Left==nil && root.Right==nil{
        path=path+strconv.Itoa(root.Val)
        *res=append(*res,path)
        //错误点
        return
    }
    //错误点 箭头的位置
    dfs(root.Left,path+strconv.Itoa(root.Val)+"->",res)
    dfs(root.Right,path+strconv.Itoa(root.Val)+"->",res)
}