package lc

import (
	"fmt"
	"testing"
)

func Test_searchRange(t *testing.T) {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{5, 6, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{6, 6, 6, 6, 6, 6, 10}, 6))
	fmt.Println(searchRange([]int{1, 4}, 4))
}
