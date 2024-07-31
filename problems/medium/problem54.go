package medium

import "fmt"

// 54. Spiral Matrix
// https://leetcode.com/problems/spiral-matrix

func SpiralOrder(matrix [][]int) []int {
	fmt.Println("width =", len(matrix[0]))
	fmt.Println("heigth =", len(matrix))

	if matrix == nil || len(matrix) == 0 {
		return []int{}
	}

	result := make([]int, 0)

	left := 0
	right := len(matrix[0]) - 1

	top := 0
	bottom := len(matrix) - 1

	for {
		// Top row (from left to right)
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++
		if top > bottom {
			break
		}

		// Right col (from top to bottom)
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--
		if left > right {
			break
		}

		// Bottom row (from left to right)
		for i := right; i >= left; i-- {
			result = append(result, matrix[bottom][i])
		}
		bottom--
		if top > bottom {
			break
		}

		// Left col (from bottom to top)
		for i := bottom; i >= top; i-- {
			result = append(result, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}

	return result
}
