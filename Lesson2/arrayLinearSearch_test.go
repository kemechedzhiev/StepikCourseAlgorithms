package Lesson2

import "testing"

func TestArraysLinearSearch2(t *testing.T) {
	var size uint
	size = 4
	arr1, arr2 := make([]int, size), make([]int, size)
	arr1[0] = 4
	arr1[1] = -8
	arr1[2] = 6
	arr1[3] = 0
	arr2[0] = -10
	arr2[1] = 3
	arr2[2] = 1
	arr2[3] = 1
	result1, result2 := ArraysLinearSearch(&arr1, &arr2, size)
	if result1 != 0 && result2 != 1 {
		t.Error("Warning! Wrong answer. Expected 0 & 1, got ", result1, result2)
	}
}
