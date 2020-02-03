package main

import (
	"StepikCourseAlgorithms/Lesson2"
	"fmt"
)

func main() {
	fmt.Println("Started")
	var size1, size2 uint
	fmt.Scan(&size1)
	arr1 := make([]int, size1)
	Lesson2.FillArray(&arr1, size1)
	fmt.Scan(&size2)
	arr2 := make([]int, size2)
	Lesson2.FillArray(&arr2, size2)
	fmt.Println(Lesson2.ArrayBinarySearch(size1, size2, &arr1, &arr2))
}
