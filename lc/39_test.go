package lc

import (
	"fmt"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
	fmt.Println(combinationSum([]int{2}, 1))
	fmt.Println(combinationSum([]int{1}, 1))
	fmt.Println(combinationSum([]int{1}, 2))
	fmt.Println(combinationSum([]int{1, 2, 3}, 4))
	fmt.Println(combinationSum([]int{2, 3, 5}, 11))
	fmt.Println(combinationSum([]int{2, 3, 5}, 9))
	fmt.Println(combinationSum([]int{2, 3, 5}, 10))
	fmt.Println(combinationSum([]int{2, 3, 5}, 12))
	fmt.Println(combinationSum([]int{2, 3, 5}, 13))
}
