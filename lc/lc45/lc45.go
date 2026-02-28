package lc45

func jump(nums []int) int {
	var dp = make([]int, len(nums))
	dp[0] = 0
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < nums[i]; j++ {
			if i+j+1 >= len(nums) {
				break
			}
			if dp[i+j+1] == 0 {
				dp[i+j+1] = dp[i] + 1
				continue
			}
			dp[i+j+1] = min(dp[i+j+1], dp[i]+1)
		}
	}
	return dp[len(nums)-1]
}
