package lc322

import (
	"fmt"
	"testing"
)

func Test_coinChange(t *testing.T) {
	fmt.Println(coinChange([]int{1, 2}, 3))
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 3))
	fmt.Println(coinChange([]int{1}, 0))
	fmt.Println(coinChange([]int{1}, 1))
	fmt.Println(coinChange([]int{1, 2, 5}, 19))
}
