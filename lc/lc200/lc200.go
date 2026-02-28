package lc200

import "fmt"

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	var res int

	dfs := func(col, row int) {}
	dfs = func(col, row int) {
		if col < 0 || row < 0 || col >= len(grid) || row >= len(grid[0]) || grid[col][row] != '1' {
			return
		}
		grid[col][row] = '2'
		dfs(col-1, row)
		dfs(col, row-1)
		dfs(col+1, row)
		dfs(col, row+1)
	}

	for col := 0; col < len(grid); col++ {
		for row := 0; row < len(grid[col]); row++ {
			if grid[col][row] == '1' {
				res++
				dfs(col, row)
			}
			fmt.Println(grid)
		}
	}

	return res
}
