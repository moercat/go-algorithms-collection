package problem

import (
	"testing"
)

func TestFibonacciDP(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
	}

	for _, test := range tests {
		result := FibonacciDP(test.n)
		if result != test.expected {
			t.Errorf("FibonacciDP(%d) = %d; expected %d", test.n, result, test.expected)
		}
	}
}

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6}, // [4, -1, 2, 1]
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
		{[]int{-1}, -1},
		{[]int{-2, -1}, -1},
	}

	for _, test := range tests {
		result := MaxSubArray(test.nums)
		if result != test.expected {
			t.Errorf("MaxSubArray(%v) = %d; expected %d", test.nums, result, test.expected)
		}
	}
}

func TestClimbingStairs(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{6, 13},
	}

	for _, test := range tests {
		result := ClimbingStairs(test.n)
		if result != test.expected {
			t.Errorf("ClimbingStairs(%d) = %d; expected %d", test.n, result, test.expected)
		}
	}
}

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		m, n     int
		expected int
	}{
		{3, 2, 3},
		{3, 7, 28},
		{7, 3, 28},
		{3, 3, 6},
	}

	for _, test := range tests {
		result := UniquePaths(test.m, test.n)
		if result != test.expected {
			t.Errorf("UniquePaths(%d, %d) = %d; expected %d", test.m, test.n, result, test.expected)
		}
	}
}

func TestCoinChange(t *testing.T) {
	tests := []struct {
		coins    []int
		amount   int
		expected int
	}{
		{[]int{1, 3, 4}, 6, 2}, // 3 + 3 = 6
		{[]int{2}, 3, -1},      // 无法组成
		{[]int{1}, 0, 0},       // 0元需要0枚硬币
		{[]int{1, 2, 5}, 11, 3}, // 5 + 5 + 1 = 11
	}

	for _, test := range tests {
		result := CoinChange(test.coins, test.amount)
		if result != test.expected {
			t.Errorf("CoinChange(%v, %d) = %d; expected %d", test.coins, test.amount, result, test.expected)
		}
	}
}

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []struct {
		text1    string
		text2    string
		expected int
	}{
		{"abcde", "ace", 3},     // "ace"
		{"abc", "abc", 3},       // "abc"
		{"abc", "def", 0},       // 无公共子序列
		{"", "abc", 0},          // 空字符串
		{"abc", "", 0},          // 空字符串
		{"", "", 0},             // 两个空字符串
		{"bl", "yby", 1},        // "b"
	}

	for _, test := range tests {
		result := LongestCommonSubsequence(test.text1, test.text2)
		if result != test.expected {
			t.Errorf("LongestCommonSubsequence(%q, %q) = %d; expected %d", test.text1, test.text2, result, test.expected)
		}
	}
}

func TestRob(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3, 1}, 4},   // 偷第0和第2间房 (1+3=4)
		{[]int{2, 7, 9, 3, 1}, 12}, // 偷第0,2,4间房 (2+9+1=12)
		{[]int{2, 1, 1, 2}, 4},   // 偷第0和第3间房 (2+2=4)
		{[]int{1}, 1},            // 只有一间房
		{[]int{}, 0},             // 没有房
	}

	for _, test := range tests {
		result := Rob(test.nums)
		if result != test.expected {
			t.Errorf("Rob(%v) = %d; expected %d", test.nums, result, test.expected)
		}
	}
}