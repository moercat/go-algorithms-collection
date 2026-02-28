package lc

import "fmt"

func canJump(nums []int) bool {
	var ans = make([]bool, len(nums))
	ans[0] = true
	for i := 0; i < len(nums); i++ {
		fmt.Println(i, nums[i], ans[i])
		for j := 0; j < nums[i]; j++ {
			if i+j+1 >= len(nums) {
				break
			}
			ans[i+j+1] = ans[i]
		}
	}
	return ans[len(nums)-1]
}
