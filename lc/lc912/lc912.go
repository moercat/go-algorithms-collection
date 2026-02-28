package lc912

import (
	"math/rand"
	"slices"
)

func sortPart(nums []int) int {
	n := len(nums)
	k := rand.Intn(n)
	item := nums[k]
	nums[0], nums[k] = nums[k], nums[0]

	i, j := 1, n-1
	for {
		for i <= j && nums[i] < item {
			i++
		}
		for i <= j && nums[j] > item {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	nums[j], nums[0] = nums[0], nums[j]
	return j
}

func sortArray(nums []int) []int {

	if slices.IsSorted(nums) {
		return nums
	}

	k := sortPart(nums)
	sortArray(nums[:k])
	sortArray(nums[k+1:])

	return nums
}

func sortArray2(nums []int) []int {
	// 优化：如果 nums 已是升序，直接返回
	if slices.IsSorted(nums) {
		return nums
	}

	i := sortPart(nums)    // 划分 nums
	sortArray2(nums[:i])   // 排序在 pivot 左侧的元素
	sortArray2(nums[i+1:]) // 排序在 pivot 右侧的元素
	return nums
}
