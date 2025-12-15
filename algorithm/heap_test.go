package algorithm

import (
	"container/heap"
	"testing"
)

func TestIntHeap(t *testing.T) {
	h := &IntHeap{}
	heap.Init(h)

	// 测试Push
	values := []int{3, 1, 4, 1, 5, 9, 2, 6}
	for _, v := range values {
		heap.Push(h, v)
	}

	// 测试Pop，应该是升序排列
	expected := []int{1, 1, 2, 3, 4, 5, 6, 9}
	for i, exp := range expected {
		if h.Len() == 0 {
			t.Errorf("Heap is empty but expected %d at position %d", exp, i)
			break
		}
		actual := heap.Pop(h).(int)
		if actual != exp {
			t.Errorf("Expected %d at position %d, got %d", exp, i, actual)
		}
	}
}

func TestMaxHeap(t *testing.T) {
	h := &MaxHeap{}
	heap.Init(h)

	// 测试Push
	values := []int{3, 1, 4, 1, 5, 9, 2, 6}
	for _, v := range values {
		heap.Push(h, v)
	}

	// 测试Pop，应该是降序排列
	expected := []int{9, 6, 5, 4, 3, 2, 1, 1}
	for i, exp := range expected {
		if h.Len() == 0 {
			t.Errorf("Heap is empty but expected %d at position %d", exp, i)
			break
		}
		actual := heap.Pop(h).(int)
		if actual != exp {
			t.Errorf("Expected %d at position %d, got %d", exp, i, actual)
		}
	}
}

func TestHeapSort(t *testing.T) {
	input := []int{3, 1, 4, 1, 5, 9, 2, 6}
	expected := []int{1, 1, 2, 3, 4, 5, 6, 9}
	result := HeapSort(input)

	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("At index %d: expected %d, got %d", i, expected[i], v)
		}
	}
}