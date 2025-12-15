package utils

// ReverseArray 反转数组
func ReverseArray(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// RemoveDuplicates 移除数组中的重复元素
func RemoveDuplicates(arr []int) []int {
	seen := make(map[int]bool)
	result := []int{}
	
	for _, item := range arr {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}
	
	return result
}

// Contains 检查数组是否包含指定元素
func Contains(arr []int, target int) bool {
	for _, item := range arr {
		if item == target {
			return true
		}
	}
	return false
}

// Max 返回数组中的最大值
func Max(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

// Min 返回数组中的最小值
func Min(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	
	min := arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}

// Sum 计算数组元素的和
func Sum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

// Filter 根据条件过滤数组
func Filter(arr []int, predicate func(int) bool) []int {
	result := []int{}
	for _, item := range arr {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}