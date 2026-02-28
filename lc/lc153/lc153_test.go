package lc153

import (
	"fmt"
	"testing"
)

func Test_findMin(t *testing.T) {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	fmt.Println(findMin([]int{2, 3, 1}))
	fmt.Println(findMin([]int{1}))
	fmt.Println(findMin([]int{3, 1, 2}))
	fmt.Println(findMin([]int{1, 2, 3}))
	fmt.Println(findMin([]int{4, 5, 3}))
	fmt.Println(findMin([]int{3, 1}))
	fmt.Println(findMin([]int{8, 9, 10, 2, 3}))
	fmt.Println(findMin([]int{3, 4, 5, 6, 2}))
}
