package lc

import (
	"fmt"
	"testing"
)

func Test_search(t *testing.T) {
	fmt.Println(search([]int{1, 3}, 0))
}

func Test_search2(t *testing.T) {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
}
