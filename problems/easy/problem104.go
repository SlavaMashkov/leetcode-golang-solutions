package easy

import "leetcode/problems/structures"

// 104. Maximum Depth of Binary Tree
// https://leetcode.com/problems/maximum-depth-of-binary-tree/

func MaxDepth(root *structures.TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(MaxDepth(root.Left), MaxDepth(root.Right))
}
