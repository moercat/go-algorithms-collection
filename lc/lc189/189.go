package lc189

func rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		newNums[(i+k)%len(nums)] = nums[i]
	}
	for i := 0; i < len(newNums); i++ {
		nums[i] = newNums[i]
	}
}
