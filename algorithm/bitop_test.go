package algorithm

import (
	"testing"
)

// 位操作相关的测试
func TestCountOneBits(t *testing.T) {
	tests := []struct {
		input    uint
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{15, 4},
		{255, 8},
	}

	for _, test := range tests {
		result := CountOneBits(test.input)
		if result != test.expected {
			t.Errorf("CountOneBits(%d) = %d; expected %d", test.input, result, test.expected)
		}
	}
}

func TestIsPowerOfTwo(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{0, false},
		{1, true},
		{2, true},
		{3, false},
		{4, true},
		{7, false},
		{8, true},
		{16, true},
		{-1, false},
		{-2, false},
	}

	for _, test := range tests {
		result := IsPowerOfTwo(test.input)
		if result != test.expected {
			t.Errorf("IsPowerOfTwo(%d) = %t; expected %t", test.input, result, test.expected)
		}
	}
}

func TestNextPowerOfTwo(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 4},
		{4, 4},
		{5, 8},
		{7, 8},
		{8, 8},
		{9, 16},
		{15, 16},
		{16, 16},
	}

	for _, test := range tests {
		result := NextPowerOfTwo(test.input)
		if result != test.expected {
			t.Errorf("NextPowerOfTwo(%d) = %d; expected %d", test.input, result, test.expected)
		}
	}
}

func TestGrayCode(t *testing.T) {
	tests := []struct {
		n        int
		expected []int
	}{
		{0, []int{0}},
		{1, []int{0, 1}},
		{2, []int{0, 1, 3, 2}},
		{3, []int{0, 1, 3, 2, 6, 7, 5, 4}},
	}

	for _, test := range tests {
		result := GrayCode(test.n)
		if len(result) != len(test.expected) {
			t.Errorf("GrayCode(%d) length = %d; expected %d", test.n, len(result), len(test.expected))
			continue
		}
		
		for i, v := range result {
			if v != test.expected[i] {
				t.Errorf("GrayCode(%d)[%d] = %d; expected %d", test.n, i, v, test.expected[i])
			}
		}
	}
}

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
		{[]int{1}, 1},
		{[]int{-1, -1, 0}, 0},
	}

	for _, test := range tests {
		result := SingleNumber(test.input)
		if result != test.expected {
			t.Errorf("SingleNumber(%v) = %d; expected %d", test.input, result, test.expected)
		}
	}
}