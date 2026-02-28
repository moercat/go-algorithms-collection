package lc118

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	if numRows == 1 {
		return [][]int{{1}}
	}
	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}
	var res = make([][]int, numRows)
	res[0] = []int{1}
	res[1] = []int{1, 1}
	for i := 2; i < numRows; i++ {
		for j := 0; j < len(res[i-1]); j++ {
			if j == len(res[i-1])-1 {
				res[i] = append(res[i], 1)
				break
			}
			if j == 0 {
				res[i] = append(res[i], 1)
			}
			res[i] = append(res[i], res[i-1][j]+res[i-1][j+1])
		}
	}
	return res
}
