package easy

// 383. Ransom Note
// https://leetcode.com/problems/ransom-note/

func CanConstruct(ransomNote string, magazine string) bool {
	m := make(map[rune]int)

	for _, char := range magazine {
		_, contains := m[char]

		if contains {
			m[char] = m[char] + 1
		} else {
			m[char] = 1
		}
	}

	for _, char := range ransomNote {
		m[char] = m[char] - 1

		if m[char] < 0 {
			return false
		}
	}

	return true
}
