package easy

// 1342. Number of Steps to Reduce a Number to Zero
// https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/

func NumberOfSteps(num int) int {
	counter := 0

	for num != 0 {
		if num%2 == 0 {
			num /= 2
			counter++
		} else {
			num -= 1
			counter++
		}
	}

	return counter
}
