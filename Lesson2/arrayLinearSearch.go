package Lesson2

import "fmt"

// TODO Даны два массива целых чисел одинаковой длины A[0..n−1] и B[0..n−1].
//  Необходимо найти первую пару индексов i0 и j0,
//  такую что A[i] + B[j] = max{A[i]+B[j], где 0 ≤ i < n, 0 ≤j <n, i ≤ j}.
//  Время работы – O(n).
//  Ограничения: 1 ≤ n ≤100000, 0 ≤ A[i] ≤100000, 0 ≤ B[i] ≤100000 для любого i.
//  Sample Input:
//  4
//  4 -8 6 0
//  -10 3 1 1
//  Sample Output:
//  0 1
func ArraysLinearSearch(arr1 *[]int, arr2 *[]int, size uint) (uint, uint) {
	var counter, result1, result2, currentMaxIndex uint
	currentMax := (*arr1)[0]
	maxSum := (*arr1)[0] + (*arr2)[0]
	for counter = 1; counter < size; counter++ {
		if currentMax+(*arr2)[counter] > maxSum {
			maxSum = currentMax + (*arr2)[counter]
			result1 = currentMaxIndex
			result2 = counter
		}
		if currentMax < (*arr1)[counter] {
			currentMax = (*arr1)[counter]
			currentMaxIndex = counter
		}
	}
	return result1, result2
}

func FillArray(arr *[]int, size uint) {
	var counter uint
	for counter = 0; counter < size; counter++ {
		fmt.Scan((*arr)[counter])
	}
}
