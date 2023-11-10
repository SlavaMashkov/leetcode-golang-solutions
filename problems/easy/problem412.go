package easy

import (
	"strconv"
)

// 412. Fizz Buzz
// https://leetcode.com/problems/fizz-buzz/

func FizzBuzz(n int) []string {
	result := make([]string, n)

	for i := 1; i <= n; i++ {
		divisibleBy3 := i%3 == 0
		divisibleBy5 := i%5 == 0

		if divisibleBy3 && divisibleBy5 {
			result[i-1] = "FizzBuzz"
		} else if divisibleBy3 {
			result[i-1] = "Fizz"
		} else if divisibleBy5 {
			result[i-1] = "Buzz"
		} else {
			result[i-1] = strconv.Itoa(i)
		}
	}

	return result
}
