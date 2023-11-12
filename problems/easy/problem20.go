package easy

// 20. Valid Parentheses
// https://leetcode.com/problems/valid-parentheses/

func IsValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := make([]string, 0, len(s))

	for _, v := range s {
		if string(v) == "(" || string(v) == "[" || string(v) == "{" {
			stack = append(stack, string(v))
		} else if string(v) == ")" && len(stack) != 0 && stack[len(stack)-1] == "(" {
			stack = stack[:len(stack)-1]
		} else if string(v) == "]" && len(stack) != 0 && stack[len(stack)-1] == "[" {
			stack = stack[:len(stack)-1]
		} else if string(v) == "}" && len(stack) != 0 && stack[len(stack)-1] == "{" {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}
