package lc3

func lengthOfLongestSubstring(s string) int {

	m := make(map[string]int)
	left := -1
	maxLength := 0
	for i := 0; i < len(s); i++ {
		item := string(s[i])

		if v, exist := m[item]; !exist {
			m[item] = i
			maxLength = max(maxLength, i-left)
		} else if v > left {
			left = v
			m[item] = i
		} else {
			m[item] = i
		}
	}
	return maxLength
}
