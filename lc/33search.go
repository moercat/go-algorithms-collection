package lc

func search(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	k := left + (right-left)/2
	for {
		if left >= right && nums[k] != target {
			return -1
		}

		if nums[k] == target {
			return k
		} else if nums[left] == target {
			return left
		} else if nums[right] == target {
			return right
		} else if nums[k] > nums[left] && nums[left] > target {
			left = k + 1
			right--
		} else if nums[k] > nums[left] && nums[left] < target && nums[k] > target {
			right = k - 1
			left++
		} else if nums[k] > nums[left] && nums[k] < target {
			left = k + 1
			right--
		} else if nums[k] < nums[left] && nums[left] < target {
			right = k - 1
			left++
		} else if nums[k] < nums[left] && nums[left] > target && nums[k] < target {
			left = k + 1
			right--
		} else if nums[k] < nums[left] && nums[k] > target {
			right = k - 1
			left++
		} else {
			left++
			right--
		}
		k = left + (right-left)/2
	}
}
