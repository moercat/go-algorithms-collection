package lc75

func sortColors(nums []int) {
	left := 0
	right := len(nums) - 1
	for i := 0; i < len(nums); i++ {
		if i > right {
			break
		}
		if nums[i] == 0 {
			nums[left] = 0
			left++
		} else if nums[i] == 2 {
			nums[i], nums[right] = nums[right], nums[i]
			right--
			i--
		}
	}
	for i := left; i <= right; i++ {
		nums[i] = 1
	}
}
