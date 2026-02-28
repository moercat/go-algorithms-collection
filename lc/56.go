package lc

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var res [][]int
	var check = intervals[0]
	for i := 1; i < len(intervals); i++ {
		if check[1] >= intervals[i][0] {
			check = []int{min(check[0], intervals[i][0]), max(check[1], intervals[i][1])}
		} else if check[1] < intervals[i][0] {
			res = append(res, check)
			check = intervals[i]
		}
	}
	if check != nil {
		res = append(res, check)
	}

	return res
}
