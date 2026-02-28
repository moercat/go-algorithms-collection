package lc153

func findMin(nums []int) int {
	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}
	left := 1
	right := len(nums) - 1
	minNum := nums[0]
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid-1] && nums[mid] > minNum {
			left = mid + 1
		} else if nums[mid] > nums[mid-1] && nums[mid] < minNum {
			right = mid - 1
			minNum = nums[mid]
		} else {
			return nums[mid]
		}
	}
	return minNum
}
