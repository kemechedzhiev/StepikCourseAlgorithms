package Lesson1

// TODO Написать реализацию рекурсивного алгоритма
func FibonacciRecursive(number uint) uint {
	if number == 0 || number == 1 {
		return number
	}
	return FibonacciRecursive(number-1) + FibonacciRecursive(number-2)
}

// TODO Написать реализацию более эффективного алгоритма, O(n)
func FibonacciNonRecursive(number uint) uint {
	if number == 0 {
		return 1
	}
	var prev, current, counter uint
	prev = 1
	current = 1
	for counter = 2; counter < number; counter++ {
		var temp uint = current
		current += prev
		prev = temp
	}
	return current
}

// TODO Даны целые числа 1≤n≤10^18 и 2≤m≤10^5, необходимо
//  найти остаток от деления n-го числа Фибоначчи на m.
