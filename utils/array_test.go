package utils

import (
	"testing"
)

func TestReverseArray(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	expected := []int{5, 4, 3, 2, 1}
	ReverseArray(arr)
	
	for i, v := range arr {
		if v != expected[i] {
			t.Errorf("ReverseArray: at index %d, expected %d, got %d", i, expected[i], v)
		}
	}
}

func TestRemoveDuplicates(t *testing.T) {
	arr := []int{1, 2, 2, 3, 4, 4, 5}
	expected := []int{1, 2, 3, 4, 5}
	result := RemoveDuplicates(arr)
	
	if len(result) != len(expected) {
		t.Errorf("RemoveDuplicates: expected length %d, got %d", len(expected), len(result))
		return
	}
	
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("RemoveDuplicates: at index %d, expected %d, got %d", i, expected[i], v)
		}
	}
}

func TestContains(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	
	if !Contains(arr, 3) {
		t.Error("Contains should return true for element 3")
	}
	
	if Contains(arr, 6) {
		t.Error("Contains should return false for element 6")
	}
}

func TestMax(t *testing.T) {
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6}
	expected := 9
	
	if result := Max(arr); result != expected {
		t.Errorf("Max: expected %d, got %d", expected, result)
	}
}

func TestMin(t *testing.T) {
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6}
	expected := 1
	
	if result := Min(arr); result != expected {
		t.Errorf("Min: expected %d, got %d", expected, result)
	}
}

func TestSum(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	expected := 15
	
	if result := Sum(arr); result != expected {
		t.Errorf("Sum: expected %d, got %d", expected, result)
	}
}

func TestFilter(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	evenPredicate := func(x int) bool { return x%2 == 0 }
	expected := []int{2, 4, 6}
	
	result := Filter(arr, evenPredicate)
	
	if len(result) != len(expected) {
		t.Errorf("Filter: expected length %d, got %d", len(expected), len(result))
		return
	}
	
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Filter: at index %d, expected %d, got %d", i, expected[i], v)
		}
	}
}