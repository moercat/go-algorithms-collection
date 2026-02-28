package lc912

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	nums := []int{5, 4, 3, 2, 1}
	sortArray(nums)

	fmt.Println(nums)

	nums2 := []int{5, 4, 3, 2, 1}
	sortArray2(nums2)

	fmt.Println(nums2)

}
