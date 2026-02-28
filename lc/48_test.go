package lc

import (
	"testing"
)

func Test_rotate(t *testing.T) {
	rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	rotate([][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}})
}
