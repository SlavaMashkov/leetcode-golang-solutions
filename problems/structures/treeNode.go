package structures

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (treeNode *TreeNode) String() string {
	return fmt.Sprintf("%v", treeNode.Val)
}
