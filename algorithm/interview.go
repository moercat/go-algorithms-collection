package algorithm

import (
	"math"
	"sort"
)

// SlidingWindowMaximum 滑动窗口最大值
func SlidingWindowMaximum(nums []int, k int) []int {
	if len(nums) == 0 || k <= 0 {
		return []int{}
	}
	
	result := make([]int, 0, len(nums)-k+1)
	// 使用双端队列存储数组索引
	deque := make([]int, 0)
	
	for i := 0; i < len(nums); i++ {
		// 移除超出窗口范围的元素
		for len(deque) > 0 && deque[0] < i-k+1 {
			deque = deque[1:]
		}
		
		// 移除所有小于当前元素的索引
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}
		
		deque = append(deque, i)
		
		// 当窗口大小达到k时，开始记录结果
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}
	
	return result
}

// NextPermutation 下一个排列
func NextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	
	// 从右往左找第一个nums[i] < nums[i+1]的位置
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	
	if i >= 0 {
		// 从右往左找第一个大于nums[i]的数
		j := len(nums) - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		
		// 交换nums[i]和nums[j]
		nums[i], nums[j] = nums[j], nums[i]
	}
	
	// 反转i+1到末尾的部分
	reverse(nums[i+1:])
}

// reverse 反转数组
func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// Permute 全排列
func Permute(nums []int) [][]int {
	result := [][]int{}
	permuteHelper(nums, 0, &result)
	return result
}

// permuteHelper 全排列辅助函数
func permuteHelper(nums []int, start int, result *[][]int) {
	if start == len(nums) {
		// 创建当前排列的副本并添加到结果中
		temp := make([]int, len(nums))
		copy(temp, nums)
		*result = append(*result, temp)
		return
	}
	
	for i := start; i < len(nums); i++ {
		// 交换
		nums[start], nums[i] = nums[i], nums[start]
		// 递归处理剩下的元素
		permuteHelper(nums, start+1, result)
		// 回溯，恢复原始顺序
		nums[start], nums[i] = nums[i], nums[start]
	}
}

// CombinationSum 组合总和
func CombinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	current := []int{}
	combinationSumHelper(candidates, target, 0, current, &result)
	return result
}

// combinationSumHelper 组合总和辅助函数
func combinationSumHelper(candidates []int, target, start int, current []int, result *[][]int) {
	if target == 0 {
		// 找到一个有效组合，添加到结果中
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}
	
	for i := start; i < len(candidates); i++ {
		if candidates[i] <= target {
			// 选择当前数字
			current = append(current, candidates[i])
			// 递归处理剩余目标值
			combinationSumHelper(candidates, target-candidates[i], i, current, result)
			// 回溯，移除当前数字
			current = current[:len(current)-1]
		}
	}
}

// MergeIntervals 合并区间
func MergeIntervals(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	
	// 按区间起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	
	result := [][]int{intervals[0]}
	
	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]
		current := intervals[i]
		
		if current[0] <= last[1] { // 有重叠
			// 合并区间，更新结束位置为两个区间结束位置的最大值
			last[1] = max(last[1], current[1])
		} else { // 无重叠
			// 添加新区间
			result = append(result, current)
		}
	}
	
	return result
}

// max 辅助函数
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// min 辅助函数
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// FindMedianSortedArrays 在两个排序数组中找到中位数
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 确保nums1是较短的数组
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	
	x, y := len(nums1), len(nums2)
	low, high := 0, x
	
	for low <= high {
		partitionX := (low + high) / 2
		partitionY := (x+y+1)/2 - partitionX
		
		// 如果partitionX为0，意味着左边没有元素，使用负无穷
		maxLeftX := math.MinInt32
		if partitionX != 0 {
			maxLeftX = nums1[partitionX-1]
		}
		
		// 如果partitionX等于数组长度，意味着右边没有元素，使用正无穷
		minRightX := math.MaxInt32
		if partitionX != x {
			minRightX = nums1[partitionX]
		}
		
		// 对nums2做同样处理
		maxLeftY := math.MinInt32
		if partitionY != 0 {
			maxLeftY = nums2[partitionY-1]
		}
		
		minRightY := math.MaxInt32
		if partitionY != y {
			minRightY = nums2[partitionY]
		}
		
		if maxLeftX <= minRightY && maxLeftY <= minRightX {
			// 找到正确的分割点
			if (x+y)%2 == 0 {
				// 偶数个元素
				return float64(max(maxLeftX, maxLeftY)+min(minRightX, minRightY)) / 2.0
			} else {
				// 奇数个元素
				return float64(max(maxLeftX, maxLeftY))
			}
		} else if maxLeftX > minRightY {
			// partitionX太大，需要向左移动
			high = partitionX - 1
		} else {
			// partitionX太小，需要向右移动
			low = partitionX + 1
		}
	}
	
	// 数组未排序，返回错误值
	return 0
}

// WordBreak 单词拆分
func WordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}
	
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true // 空字符串可以被拆分
	
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	
	return dp[n]
}

// LadderLength 最短转换序列长度 (BFS)
func LadderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}
	
	if !wordSet[endWord] {
		return 0
	}
	
	queue := []string{beginWord}
	visited := make(map[string]bool)
	visited[beginWord] = true
	
	level := 1
	
	for len(queue) > 0 {
		size := len(queue)
		
		for i := 0; i < size; i++ {
			currentWord := queue[0]
			queue = queue[1:]
			
			if currentWord == endWord {
				return level
			}
			
			// 尝试改变currentWord的每个字符
			for j := 0; j < len(currentWord); j++ {
				for c := 'a'; c <= 'z'; c++ {
					newWord := currentWord[:j] + string(c) + currentWord[j+1:]
					
					if wordSet[newWord] && !visited[newWord] {
						visited[newWord] = true
						queue = append(queue, newWord)
					}
				}
			}
		}
		
		level++
	}
	
	return 0
}

// NumIslands 岛屿数量 (DFS)
func NumIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	
	count := 0
	m, n := len(grid), len(grid[0])
	
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfsIsland(grid, i, j, m, n)
				count++
			}
		}
	}
	
	return count
}

// dfsIsland 深度优先搜索标记岛屿
func dfsIsland(grid [][]byte, i, j, m, n int) {
	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != '1' {
		return
	}
	
	grid[i][j] = '0' // 标记已访问
	
	// 搜索四个方向
	dfsIsland(grid, i-1, j, m, n) // 上
	dfsIsland(grid, i+1, j, m, n) // 下
	dfsIsland(grid, i, j-1, m, n) // 左
	dfsIsland(grid, i, j+1, m, n) // 右
}

// CanJump 跳跃游戏
func CanJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	
	farthest := 0
	
	for i := 0; i < len(nums); i++ {
		// 如果当前位置无法到达，返回false
		if i > farthest {
			return false
		}
		
		// 更新最远可达位置
		farthest = max(farthest, i+nums[i])
		
		// 如果能到达末尾，返回true
		if farthest >= len(nums)-1 {
			return true
		}
	}
	
	return false
}

// ThreeSum 三数之和
func ThreeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	
	// 先排序
	sort.Ints(nums)
	result := [][]int{}
	
	for i := 0; i < len(nums)-2; i++ {
		// 跳过重复元素
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		
		left, right := i+1, len(nums)-1
		
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				
				// 跳过重复元素
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	
	return result
}

// LongestPalindrome 最长回文子串 (中心扩展法)
func LongestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	
	start, maxLen := 0, 0
	
	for i := 0; i < len(s); i++ {
		// 以i为中心的奇数长度回文
		len1 := expandAroundCenter(s, i, i)
		// 以i和i+1为中心的偶数长度回文
		len2 := expandAroundCenter(s, i, i+1)
		
		currentMax := max(len1, len2)
		if currentMax > maxLen {
			maxLen = currentMax
			start = i - (currentMax-1)/2
		}
	}
	
	return s[start : start+maxLen]
}

// expandAroundCenter 从中心向外扩展
func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

// ValidParentheses 有效的括号
func ValidParentheses(s string) bool {
	stack := []rune{}
	mapping := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	
	for _, char := range s {
		if opening, exists := mapping[char]; exists {
			var topElement rune
			if len(stack) == 0 {
				topElement = '#'
			} else {
				topElement = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
			
			if topElement != opening {
				return false
			}
		} else {
			stack = append(stack, char)
		}
	}
	
	return len(stack) == 0
}

// Trap 接雨水
func Trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	result := 0
	
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				result += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				result += rightMax - height[right]
			}
			right--
		}
	}
	
	return result
}

// FindDuplicateArray 找出数组中重复的数字（Floyd环检测）
func FindDuplicateArray(nums []int) int {
	// 使用Floyd环检测算法
	// 将数组视为链表，nums[i]指向索引nums[i]
	
	slow, fast := nums[0], nums[0]
	
	// 第一阶段：找到环中的相遇点
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	
	// 第二阶段：找到环的入口（重复的数字）
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	
	return slow
}

// RobHouse 打家劫舍（房屋围成一圈）
func RobHouse(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	
	// 分两种情况：包含第一个房屋或不包含第一个房屋
	return max(robRange(nums, 0, len(nums)-2), robRange(nums, 1, len(nums)-1))
}

// robRange 计算指定范围内的最大抢劫金额
func robRange(nums []int, start, end int) int {
	if start == end {
		return nums[start]
	}
	
	dp := make([]int, len(nums))
	dp[start] = nums[start]
	dp[start+1] = max(nums[start], nums[start+1])
	
	for i := start + 2; i <= end; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	
	return dp[end]
}

// IsPalindromeNumber 判断是否是回文数
func IsPalindromeNumber(x int) bool {
	// 负数不是回文数
	if x < 0 {
		return false
	}
	
	// 计算反转后的数字
	original := x
	reversed := 0
	
	for x != 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}
	
	return original == reversed
}

// ReverseLinkedList 反转链表（递归）
func ReverseLinkedList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	
	newHead := ReverseLinkedList(head.Next)
	head.Next.Next = head
	head.Next = nil
	
	return newHead
}

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// HasCycleLinkedList 检测链表是否有环
func HasCycleLinkedList(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	
	slow, fast := head, head
	
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		
		if slow == fast {
			return true
		}
	}
	
	return false
}