package lc

import (
	"fmt"
	"testing"
)

func Test_subarraySum(t *testing.T) {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
	fmt.Println(subarraySum([]int{1, 2, 3}, 3))
	fmt.Println(subarraySum([]int{1, 2, 1, 2, 1}, 3))
	fmt.Println(subarraySum([]int{1}, 0))
	fmt.Println(subarraySum([]int{-1, -1, 1}, 0))
}
