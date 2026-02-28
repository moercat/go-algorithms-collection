package lc

import "sort"

func threeSum(nums []int) [][]int {
	// 排序
	sort.Ints(nums)
	res := [][]int{}

	// 定位 k 值在最左边边
	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		if n1 > 0 {
			break
		}

		if i > 0 && n1 == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1

		for left < right {
			n1, n2, n3 := nums[i], nums[left], nums[right]
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})

				for left < right && n2 == nums[left] {
					left++
				}

				for left < right && n3 == nums[right] {
					right--
				}
			} else if n1+n2+n3 > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return res
}
