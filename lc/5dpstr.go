package lc

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return s
	}
	maxLeft, maxRight := 1, 1
	ants := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		if ants[i] == nil {
			ants[i] = make([]bool, len(s))
		}
		ants[i][i] = true
		if i < len(s)-1 {
			ants[i][i+1] = s[i] == s[i+1]
			if ants[i][i+1] {
				maxLeft, maxRight = i, i+1
			}
		}
	}

	for step := 2; step < len(s); step++ {
		for i := 0; i < len(s); i++ {
			if i+step >= len(s) {
				break
			}
			ants[i][i+step] = ants[i+1][i+step-1] && s[i] == s[i+step]
			if ants[i][i+step] {
				maxLeft, maxRight = i, i+step
			}
		}
	}

	return s[maxLeft : maxRight+1]
}
