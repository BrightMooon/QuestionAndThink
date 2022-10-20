package btree
import "testing"
func TestXxx(t *testing.T) {
	invertTree(FrontBuildTree())
}


func FrontBuildTree() *TreeNode{
	root:=&TreeNode{4,nil,nil}
	root.Left=&TreeNode{2,nil,nil}
	root.Left.Left=&TreeNode{1,nil,nil}
	root.Left.Right=&TreeNode{3,nil,nil}
	root.Right=&TreeNode{7,nil,nil}
	root.Right.Left=&TreeNode{6,nil,nil}
	root.Right.Right=&TreeNode{9,nil,nil}
	return root
}
