package easy

// 13. Roman to Integer
// https://leetcode.com/problems/roman-to-integer/

func RomanToInt(s string) int {
	symbolValueMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	result := 0

	if len(s) == 1 {
		return symbolValueMap[string(s[0])]
	}

	for i := 0; i < len(s); i++ {
		result += symbolValueMap[string(s[i])]

		if i > 0 && symbolValueMap[string(s[i])] > symbolValueMap[string(s[i-1])] {
			result -= 2 * symbolValueMap[string(s[i-1])]
		}
	}

	return result
}
