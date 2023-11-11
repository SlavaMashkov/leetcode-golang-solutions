package easy

import "strings"

// 14. Longest Common Prefix
// https://leetcode.com/problems/longest-common-prefix/

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) != 0 {
			prefix = prefix[:len(prefix)-1]

			if prefix == "" {
				return ""
			}
		}
	}

	return prefix
}
