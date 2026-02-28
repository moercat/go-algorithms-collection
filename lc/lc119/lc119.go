package lc119

func getRow(rowIndex int) []int {
	numRows := rowIndex + 1
	if numRows == 1 {
		return []int{1}
	}
	if numRows == 2 {
		return []int{1, 1}
	}
	lastRes := []int{1, 1}
	newRes := []int{}
	for i := 2; i < numRows; i++ {
		for j := 0; j < len(lastRes); j++ {
			if j == len(lastRes)-1 {
				newRes = append(newRes, 1)
				break
			}
			if j == 0 {
				newRes = append(newRes, 1)
			}
			newRes = append(newRes, lastRes[j]+lastRes[j+1])
		}
		lastRes = newRes
		newRes = []int{}
	}
	return lastRes
}
