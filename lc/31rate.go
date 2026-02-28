package lc

func nextPermutation(nums []int) {
	n := len(nums)

	k := n - 2
	for {
		if k >= 0 && nums[k] < nums[k+1] {
			break
		}
		k--
		if k < 0 {
			break
		}
	}

	if k < 0 {
		reverse(nums, 0, n-1)
		return
	}

	j := n - 1
	for {
		if nums[k] < nums[j] {
			nums[k], nums[j] = nums[j], nums[k]
			break
		}
		j--
	}

	reverse(nums, k+1, n-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
