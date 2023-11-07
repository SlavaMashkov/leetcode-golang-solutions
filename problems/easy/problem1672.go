package easy

import (
	"fmt"
)

// 1672. Richest Customer Wealth
// https://leetcode.com/problems/richest-customer-wealth/

func MaximumWealth(accounts [][]int) int {
	maxSum := 0

	for i, account := range accounts {
		isum := 0

		for j, wealth := range account {
			fmt.Printf("accounts[%v][%v]: %v\n", i, j, accounts[i][j])
			isum += wealth
		}

		if isum > maxSum {
			maxSum = isum
		}
	}

	return maxSum
}
