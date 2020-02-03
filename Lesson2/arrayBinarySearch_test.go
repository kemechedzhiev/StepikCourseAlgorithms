package Lesson2

import "testing"

func TestArrayBinarySearch(t *testing.T) {
	var size1, size2 uint
	size1, size2 = 3, 3
	arr1, arr2 := make([]int, size1), make([]int, size2)
	arr1[0] = 10
	arr1[1] = 20
	arr1[2] = 30
	arr2[0] = 9
	arr2[1] = 15
	arr2[2] = 35
	result := ArrayBinarySearch(size1, size2, &arr1, &arr2)
	if (*result)[0] != 0 || (*result)[1] != 0 || (*result)[2] != 2 {
		t.Error("Wrong answer! Expected 0 0 2, got ", result)
	}
}
