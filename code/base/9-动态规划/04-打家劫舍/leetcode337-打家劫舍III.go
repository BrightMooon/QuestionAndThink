package dp

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
 }
 /*
 树递归的模型
 1.  空节点返回
 2.  一个节点
 3.  三层的一个树模型去考虑递归的过程
 */
 func rob(root *TreeNode) int {
	if root==nil {
		return 0
	}
	if root.Left==nil && root.Right==nil{
		return root.Val
	}
	rootVal:=root.Val
	if root.Left!=nil {
		rootVal+=rob(root.Left.Left)+rob(root.Left.Right)
	}
	if root.Right!=nil {
		rootVal+=rob(root.Right.Left)+rob(root.Right.Right)
	}
	subVal:=rob(root.Left)+rob(root.Right)
	return rootVal+subVal
}

/*
树形dp
1. 
要求一个节点 偷与不偷的两个状态所得到的金钱，那么返回值就是一个长度为2的数组
dp[0]  不偷该节点所得到最大金钱
dp[1]  偷该节点所得到的最大金钱
在递归的过程中，递归栈会保存每一层的参数
2. 确定终止条件
   root==nil return []int{0,0}
3. 确定遍历顺序
+ 使用后序遍历，因为通过递归的返回值来做下一步计算
+ 先递归左节点，得到左节点偷与不投的金钱
+ 递归右节点 得到右节点偷与不偷的金钱

4. 确定单层递归的逻辑
+ 如果偷当前节点，那么左右孩子就不能偷
val1 =cur.val+left[0]+right[0]
+ 如果不偷当前节点，可以偷左右孩子节点
+ 偷或者不偷选个最大的 
val2=max(left[0],left[1]), max(right[0],right[1])

不偷当前节点的最大金额是val1
偷当前节点的最大金额是val2

*/

/*
树形dp
1 递归和二叉树如何结合，如果是后序，左右递归返回左右子树要求的最终结果，操作当前的cur
2 两个选择，偷 cur+left[0]+right[0],不偷  当前，可以选择子节点偷不偷，选个最大的max（left）+max(right)
3 注意返回数组的赋值顺序，index=0是不偷，index=1是偷，要统一
*/

func robDp(root *TreeNode) int {
	//1. 后序遍历如何与递归结合
	//2. 左右根，左右递归遍历，根节点就是当前节点，取出当前值去做逻辑处理
	//3. []int{}
		dp:=help(root)
		return max(dp[0],dp[1])
	}
	
	func help(root *TreeNode)[]int{
		if root==nil{
			return []int{0,0}
		}
		//盲区，请问这里递归结果返回的是什么，是左子树的遍历结果，即左节点头或者不偷的结果
		left:=help(root.Left)
		right:=help(root.Right)
		//盲区
		valT:=root.Val+left[0]+right[0]
		//不偷cur，但是可以偷或者不偷左右节点，比较最大值的情况
		valF:=max(left[0],left[1])+max(right[0],right[1])
		//盲区:注意这里的顺序，前面是index=0是不偷，后面index=1是偷
		return []int{valF,valT}
	}
	
	func max(a,b int)int{
		if a<b{
			return b
		}
		return a
	}

	