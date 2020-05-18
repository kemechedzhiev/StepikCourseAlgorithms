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
func ModNFibToM(n, m uint) uint {
	var counter uint
	// Метод решения - через период Пизано. Последовательность остатков от деления чисел Фибоначчи на число
	// периодична.
	pisanoLength, pisanoArr := findPisano(m)
	for counter = 1; counter >= pisanoLength; {
		counter = n / m
	}
	return (*pisanoArr)[counter]
}

// Служебная функция.
func findPisano(delimiter uint) (uint, *[]uint) {
	var counter, pisanoLength, prevMod, currMod uint
	pisanoArr := make([]uint, delimiter, delimiter*6)
	prevMod = 0
	currMod = 1
	pisanoArr[0] = prevMod
	pisanoArr[1] = currMod
	if delimiter == 2 {
		return 2, &pisanoArr
	}
	pisanoArr[2] = currMod
	for counter = 3; counter < delimiter*6; counter++ {
		currMod = FibonacciNonRecursive(counter) % delimiter
		if currMod == 1 && prevMod == 0 {
			break
		}
		pisanoArr = append(pisanoArr, currMod)
		prevMod = currMod
		pisanoLength = counter
	}
	return pisanoLength, &pisanoArr
}
