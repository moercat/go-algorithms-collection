package lc

func combinationSum(candidates []int, target int) [][]int {

	var (
		com = make([]int, 0)
		res = make([][]int, 0)
	)
	dfs := func(target int, idx int) {}
	dfs = func(target int, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			res = append(res, append([]int{}, com...))
			return
		}

		if target-candidates[idx] >= 0 {
			com = append(com, candidates[idx])
			dfs(target-candidates[idx], idx)
			com = com[:len(com)-1]
		}

		dfs(target, idx+1)
	}
	dfs(target, 0)
	return res
}
