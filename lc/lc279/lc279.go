package lc279

import "math"

func numSquares(n int) int {
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		item := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			item = min(item, f[i-j*j]+1)
		}
		f[i] = item
	}
	return f[n]
}
