package lc322

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32 / 2
	}
	dp[0] = 0
	if amount <= 0 {
		return dp[amount]
	}
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i < coins[j] {
				continue
			}
			dp[i] = min(dp[i], dp[i-coins[j]]+1)
		}
	}
	if dp[amount] >= math.MaxInt32/2 {
		return -1
	}
	return dp[amount]
}
