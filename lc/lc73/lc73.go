package lc73

func setZeroes(matrix [][]int) {
	var m = make(map[int]struct{})
	var n = make(map[int]struct{})
	for i := 0; i < len(matrix); i++ {
		item := matrix[i]
		for j := 0; j < len(item); j++ {
			if item[j] == 0 {
				m[i] = struct{}{}
				n[j] = struct{}{}
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		item := matrix[i]
		_, existM := m[i]
		for j := 0; j < len(item); j++ {
			_, existN := n[j]
			if existM || existN {
				matrix[i][j] = 0
			}
		}
	}
}
