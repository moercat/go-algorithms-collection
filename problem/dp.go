package problem

// FibonacciDP 使用动态规划计算斐波那契数列
func FibonacciDP(n int) int {
	if n <= 1 {
		return n
	}
	
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	
	return dp[n]
}

// MaxSubArray 最大子数组和（Kadane算法）
func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	maxSoFar := nums[0]
	maxEndingHere := nums[0]
	
	for i := 1; i < len(nums); i++ {
		maxEndingHere = max(nums[i], maxEndingHere+nums[i])
		maxSoFar = max(maxSoFar, maxEndingHere)
	}
	
	return maxSoFar
}

// max 是辅助函数，返回两个数中的较大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// min 是辅助函数，返回两个数中的较小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// ClimbingStairs 爬楼梯问题
func ClimbingStairs(n int) int {
	if n <= 1 {
		return 1
	}
	
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	
	return dp[n]
}

// UniquePaths 不同路径问题
func UniquePaths(m, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	
	// 初始化第一行和第一列
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	
	// 填充dp表
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	
	return dp[m-1][n-1]
}

// CoinChange 零钱兑换问题
func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	
	// 初始化为一个比amount大的值
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// LongestCommonSubsequence 最长公共子序列
func LongestCommonSubsequence(text1, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	
	return dp[m][n]
}

// Rob 打家劫舍问题
func Rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	
	return dp[len(nums)-1]
}