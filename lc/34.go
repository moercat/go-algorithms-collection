package lc

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0, 0}
		}

		return []int{-1, -1}
	}
	left, right := 0, len(nums)-1
	if nums[left] == target && nums[right] == target {
		return []int{left, right}
	}
	for left < right {
		k := left + (right-left)/2
		if nums[right] == target {
			k = right
		}
		if nums[left] == target {
			k = left
		}
		if nums[k] < target {
			left = k + 1
		} else if nums[k] > target {
			right = k - 1
		} else {
			left = k
			right = k
			for left > 0 && nums[left-1] == nums[k] {
				left--
			}
			for right < len(nums)-1 && nums[right+1] == nums[k] {
				right++
			}
			return []int{left, right}
		}
	}

	return []int{-1, -1}
}
