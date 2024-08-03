package easy

import (
	"fmt"
	"leetcode/problems/structures"
)

// 94. Binary Tree Inorder Traversal
// https://leetcode.com/problems/binary-tree-inorder-traversal/

func InorderTraversal(root *structures.TreeNode) []int {
	var stack []structures.TreeNode
	var result []int

	current := root

	for current != nil || len(stack) != 0 {
		for current != nil {
			stack = append(stack, *current)
			current = current.Left
		}

		current = &stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, current.Val)

		current = current.Right
	}

	fmt.Printf("result: %v\n", result)

	return result
}
