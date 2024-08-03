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

func DFS(root *structures.TreeNode) []*structures.TreeNode {
	var visited []*structures.TreeNode

	if root == nil {
		return visited
	}

	//return recursiveFunc(root, visited)
	return stackFunc(root, visited)
}

func recursiveFunc(root *structures.TreeNode, visited []*structures.TreeNode) []*structures.TreeNode {
	visited = append(visited, root)

	if root.Left != nil {
		visited = recursiveFunc(root.Left, visited)
	} else {
		//visited = append(visited, nil)
	}

	if root.Right != nil {
		visited = recursiveFunc(root.Right, visited)
	} else {
		//visited = append(visited, nil)
	}

	return visited
}

func stackFunc(root *structures.TreeNode, visited []*structures.TreeNode) []*structures.TreeNode {
	stack := []*structures.TreeNode{root}

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		visited = append(visited, node)
		stack = stack[:len(stack)-1] // lifo

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return visited
}
