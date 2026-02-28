package lc

func longestValidParentheses(s string) int {

	var m = make(map[int]struct{})
	leftList := []int{}
	lessLeft := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			lessLeft++
			leftList = append(leftList, i)
		} else if lessLeft > 0 {
			m[i] = struct{}{}
			m[leftList[len(leftList)-1]] = struct{}{}
			lessLeft--
			leftList = leftList[:len(leftList)-1]
		}
	}
	maxNum, num := 0, 0
	for i := 0; i < len(s); i++ {
		if _, ok := m[i]; ok {
			num++
		} else {
			maxNum = max(maxNum, num)
			num = 0
		}
	}
	return max(maxNum, num)
}
