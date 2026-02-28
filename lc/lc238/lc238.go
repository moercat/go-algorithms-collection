package lc238

func productExceptSelf(nums []int) []int {
	var left = make([]int, len(nums))
	var right = make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		rightIdx := len(nums) - i - 1
		if i == 0 {
			left[i] = nums[i]
			right[rightIdx] = nums[rightIdx]
			continue
		}
		left[i] = nums[i] * left[i-1]
		right[rightIdx] = nums[rightIdx] * right[rightIdx+1]
	}
	var result = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			result[i] = right[i+1]
			continue
		}
		if i == len(nums)-1 {
			result[i] = left[i-1]
			continue
		}
		result[i] = left[i-1] * right[i+1]
	}
	return result
}
