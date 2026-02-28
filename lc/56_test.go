package lc

import (
	"fmt"
	"testing"
)

func Test_merge(t *testing.T) {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}}))
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))
	fmt.Println(merge([][]int{{1, 4}, {2, 3}}))
	fmt.Println(merge([][]int{{1, 4}, {2, 3}, {4, 5}}))
	fmt.Println(merge([][]int{{1, 4}, {2, 3}, {4, 5}, {6, 7}}))
	fmt.Println(merge([][]int{{1, 4}, {2, 3}, {4, 5}, {6, 7}, {8, 9}}))
}
