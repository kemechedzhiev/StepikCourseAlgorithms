package Lesson2

// TODO Дан отсортированный массив различных целых чисел A[0..n-1] и массив целых чисел B[0..m-1].
//  Для каждого элемента массива B[i] найдите минимальный индекс элемента массива A[k], ближайшего по значению к B[i].
//  Время работы поиска для каждого элемента B[i]: O(log(k)).
//  Подсказка. Обратите внимание, что время работы должно зависеть от индекса ответа - k.
//  Для достижения такой асимптотики предлагается для начала найти отрезок вида [2^p, 2^{p+1}],
//  содержащий искомую точку, а уже затем на нём выполнять традиционный бин. поиск.
//  Sample Input:
//  3
//  10 20 30
//  3
//  9 15 35
//  Sample Output:
//  0 0 2
func ArrayBinarySearch(size1, size2 uint, arr1, arr2 *[]int) *[]uint {
	var counter uint
	result := make([]uint, size2)
	for counter = 0; counter < size2; counter++ {
		result[counter] = counter
	}
	return &result
}
