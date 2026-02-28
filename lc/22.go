package lc

func generateParenthesis(n int) []string {
	var enum = "()"
	if n == 1 {
		return []string{enum}
	}
	var resMap = make(map[string]bool)
	var res = []string{enum}
	for i := 1; i < n; i++ {
		var oldRes []string
		for j := 0; j < len(res); j++ {
			s := res[j]
			for k := 0; k < len(s); k++ {
				if s[k] == ')' {
					tmp := s[:k] + enum + s[k:]
					if !resMap[tmp] {
						resMap[tmp] = true
						oldRes = append(oldRes, tmp)
					}
				}
			}
			tmp := s + enum
			if !resMap[tmp] {
				resMap[tmp] = true
				oldRes = append(oldRes, tmp)
			}
			tmp = enum + s
			if !resMap[tmp] {
				resMap[tmp] = true
				oldRes = append(oldRes, tmp)
			}
		}
		res = oldRes
	}

	return res
}
