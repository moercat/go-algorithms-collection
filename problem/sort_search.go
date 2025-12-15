package problem

// QuickSort 快速排序
func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	
	pivotIndex := partition(arr, 0, len(arr)-1)
	QuickSort(arr[:pivotIndex])
	QuickSort(arr[pivotIndex+1:])
}

// partition 是快速排序的分区函数
func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// MergeSort 归并排序
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	
	return merge(left, right)
}

// merge 是归并排序的合并函数
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	
	// 添加剩余元素
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	
	return result
}

// BinarySearch 二分搜索
func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	
	return -1
}

// BinarySearchFirst 找到目标元素第一次出现的位置
func BinarySearchFirst(arr []int, target int) int {
	left, right := 0, len(arr)-1
	result := -1
	
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			result = mid
			right = mid - 1 // 继续向左搜索
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	
	return result
}

// BinarySearchLast 找到目标元素最后一次出现的位置
func BinarySearchLast(arr []int, target int) int {
	left, right := 0, len(arr)-1
	result := -1
	
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			result = mid
			left = mid + 1 // 继续向右搜索
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	
	return result
}

// SearchInRotatedSortedArray 在旋转排序数组中搜索
func SearchInRotatedSortedArray(nums []int, target int) int {
	left, right := 0, len(nums)-1
	
	for left <= right {
		mid := left + (right-left)/2
		
		if nums[mid] == target {
			return mid
		}
		
		// 判断哪一半是有序的
		if nums[left] <= nums[mid] { // 左半部分有序
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // 右半部分有序
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	
	return -1
}

// FindKthLargest 找到数组中第K大的元素
func FindKthLargest(nums []int, k int) int {
	// 使用最小堆的思想来优化
	// 这里使用快速选择算法，平均时间复杂度O(n)
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

// quickSelect 是快速选择算法
func quickSelect(nums []int, left, right, k int) int {
	if left == right {
		return nums[left]
	}
	
	pivotIndex := partition(nums, left, right)
	
	if k == pivotIndex {
		return nums[k]
	} else if k < pivotIndex {
		return quickSelect(nums, left, pivotIndex-1, k)
	} else {
		return quickSelect(nums, pivotIndex+1, right, k)
	}
}