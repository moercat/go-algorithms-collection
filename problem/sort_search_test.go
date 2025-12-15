package problem

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	expected := []int{11, 12, 22, 25, 34, 64, 90}

	QuickSort(arr)

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("QuickSort failed: got %v, expected %v", arr, expected)
	}
}

func TestMergeSort(t *testing.T) {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	expected := []int{11, 12, 22, 25, 34, 64, 90}

	result := MergeSort(arr)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MergeSort failed: got %v, expected %v", result, expected)
	}
}

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15}

	// 测试存在的元素
	if index := BinarySearch(arr, 7); index != 3 {
		t.Errorf("BinarySearch(7) = %d; expected 3", index)
	}

	// 测试不存在的元素
	if index := BinarySearch(arr, 6); index != -1 {
		t.Errorf("BinarySearch(6) = %d; expected -1", index)
	}

	// 测试边界情况
	if index := BinarySearch(arr, 1); index != 0 {
		t.Errorf("BinarySearch(1) = %d; expected 0", index)
	}

	if index := BinarySearch(arr, 15); index != 7 {
		t.Errorf("BinarySearch(15) = %d; expected 7", index)
	}
}

func TestBinarySearchFirst(t *testing.T) {
	arr := []int{1, 2, 2, 2, 3, 4, 4, 5}

	if index := BinarySearchFirst(arr, 2); index != 1 {
		t.Errorf("BinarySearchFirst(2) = %d; expected 1", index)
	}

	if index := BinarySearchFirst(arr, 4); index != 5 {
		t.Errorf("BinarySearchFirst(4) = %d; expected 5", index)
	}

	if index := BinarySearchFirst(arr, 6); index != -1 {
		t.Errorf("BinarySearchFirst(6) = %d; expected -1", index)
	}
}

func TestBinarySearchLast(t *testing.T) {
	arr := []int{1, 2, 2, 2, 3, 4, 4, 5}

	if index := BinarySearchLast(arr, 2); index != 3 {
		t.Errorf("BinarySearchLast(2) = %d; expected 3", index)
	}

	if index := BinarySearchLast(arr, 4); index != 6 {
		t.Errorf("BinarySearchLast(4) = %d; expected 6", index)
	}

	if index := BinarySearchLast(arr, 6); index != -1 {
		t.Errorf("BinarySearchLast(6) = %d; expected -1", index)
	}
}

func TestSearchInRotatedSortedArray(t *testing.T) {
	// 旋转数组 [4,5,6,7,0,1,2]
	rotatedArr := []int{4, 5, 6, 7, 0, 1, 2}

	if index := SearchInRotatedSortedArray(rotatedArr, 0); index != 4 {
		t.Errorf("SearchInRotatedSortedArray(0) = %d; expected 4", index)
	}

	if index := SearchInRotatedSortedArray(rotatedArr, 3); index != -1 {
		t.Errorf("SearchInRotatedSortedArray(3) = %d; expected -1", index)
	}

	if index := SearchInRotatedSortedArray(rotatedArr, 7); index != 3 {
		t.Errorf("SearchInRotatedSortedArray(7) = %d; expected 3", index)
	}

	if index := SearchInRotatedSortedArray(rotatedArr, 4); index != 0 {
		t.Errorf("SearchInRotatedSortedArray(4) = %d; expected 0", index)
	}
}

func TestFindKthLargest(t *testing.T) {
	arr := []int{3, 2, 1, 5, 6, 4}

	// 第2大的元素应该是5
	if result := FindKthLargest(arr, 2); result != 5 {
		t.Errorf("FindKthLargest([3,2,1,5,6,4], 2) = %d; expected 5", result)
	}

	// 第1大的元素应该是6
	if result := FindKthLargest(arr, 1); result != 6 {
		t.Errorf("FindKthLargest([3,2,1,5,6,4], 1) = %d; expected 6", result)
	}

	// 第6大的元素应该是1
	if result := FindKthLargest(arr, 6); result != 1 {
		t.Errorf("FindKthLargest([3,2,1,5,6,4], 6) = %d; expected 1", result)
	}
}