package utils

// GCD 计算两个数的最大公约数
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM 计算两个数的最小公倍数
func LCM(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return (a * b) / GCD(a, b)
}

// IsPrime 检查一个数是否是质数
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	
	return true
}

// Fibonacci 计算斐波那契数列的第n项
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// Factorial 计算n的阶乘
func Factorial(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 || n == 1 {
		return 1
	}
	
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// Power 计算base的exp次幂
func Power(base, exp int) int {
	if exp == 0 {
		return 1
	}
	
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

// Sqrt 计算整数的平方根（牛顿法）
func Sqrt(n int) int {
	if n <= 1 {
		return n
	}
	
	x := n
	for {
		y := (x + n/x) / 2
		if y >= x {
			return x
		}
		x = y
	}
}

// Combination 计算组合 C(n, k)
func Combination(n, k int) int {
	if k > n || k < 0 {
		return 0
	}
	if k == 0 || k == n {
		return 1
	}
	
	// C(n, k) = C(n, n-k) 优化
	if k > n-k {
		k = n - k
	}
	
	result := 1
	for i := 0; i < k; i++ {
		result = result * (n - i) / (i + 1)
	}
	
	return result
}

// Permutation 计算排列 P(n, k)
func Permutation(n, k int) int {
	if k > n || k < 0 {
		return 0
	}
	
	return Factorial(n) / Factorial(n-k)
}