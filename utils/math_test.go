package utils

import (
	"testing"
)

func TestGCD(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{48, 18, 6},
		{18, 48, 6},
		{1, 1, 1},
		{0, 5, 5},
		{5, 0, 5},
		{17, 13, 1}, // 两个质数
		{100, 25, 25},
	}

	for _, test := range tests {
		result := GCD(test.a, test.b)
		if result != test.expected {
			t.Errorf("GCD(%d, %d) = %d; expected %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestLCM(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{48, 18, 144},
		{18, 48, 144},
		{1, 1, 1},
		{0, 5, 0},
		{5, 0, 0},
		{17, 13, 221}, // 两个质数
		{10, 15, 30},
	}

	for _, test := range tests {
		result := LCM(test.a, test.b)
		if result != test.expected {
			t.Errorf("LCM(%d, %d) = %d; expected %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		n        int
		expected bool
	}{
		{0, false},
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{17, true},
		{25, false},
		{29, true},
		{100, false},
	}

	for _, test := range tests {
		result := IsPrime(test.n)
		if result != test.expected {
			t.Errorf("IsPrime(%d) = %t; expected %t", test.n, result, test.expected)
		}
	}
}

func TestFibonacci(t *testing.T) {
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
		result := Fibonacci(test.n)
		if result != test.expected {
			t.Errorf("Fibonacci(%d) = %d; expected %d", test.n, result, test.expected)
		}
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
		{6, 720},
	}

	for _, test := range tests {
		result := Factorial(test.n)
		if result != test.expected {
			t.Errorf("Factorial(%d) = %d; expected %d", test.n, result, test.expected)
		}
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		base     int
		exp      int
		expected int
	}{
		{2, 0, 1},
		{5, 1, 5},
		{2, 3, 8},
		{3, 4, 81},
		{10, 2, 100},
	}

	for _, test := range tests {
		result := Power(test.base, test.exp)
		if result != test.expected {
			t.Errorf("Power(%d, %d) = %d; expected %d", test.base, test.exp, result, test.expected)
		}
	}
}

func TestSqrt(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{4, 2},
		{9, 3},
		{16, 4},
		{25, 5},
		{36, 6},
		{49, 7},
		{64, 8},
		{81, 9},
		{100, 10},
	}

	for _, test := range tests {
		result := Sqrt(test.n)
		if result != test.expected {
			t.Errorf("Sqrt(%d) = %d; expected %d", test.n, result, test.expected)
		}
	}
}

func TestCombination(t *testing.T) {
	tests := []struct {
		n, k     int
		expected int
	}{
		{5, 0, 1},
		{5, 1, 5},
		{5, 2, 10},
		{5, 3, 10},
		{5, 4, 5},
		{5, 5, 1},
		{10, 2, 45},
		{0, 0, 1},
		{5, 6, 0}, // k > n
		{5, -1, 0}, // negative k
	}

	for _, test := range tests {
		result := Combination(test.n, test.k)
		if result != test.expected {
			t.Errorf("Combination(%d, %d) = %d; expected %d", test.n, test.k, result, test.expected)
		}
	}
}

func TestPermutation(t *testing.T) {
	tests := []struct {
		n, k     int
		expected int
	}{
		{5, 0, 1},
		{5, 1, 5},
		{5, 2, 20},
		{5, 3, 60},
		{5, 4, 120},
		{5, 5, 120},
		{10, 2, 90},
		{0, 0, 1},
		{5, 6, 0}, // k > n
		{5, -1, 0}, // negative k
	}

	for _, test := range tests {
		result := Permutation(test.n, test.k)
		if result != test.expected {
			t.Errorf("Permutation(%d, %d) = %d; expected %d", test.n, test.k, result, test.expected)
		}
	}
}