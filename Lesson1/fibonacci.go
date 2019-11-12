package Lesson1

// It's an obviously non-effective algorithm implemented just for the comparison.
func FibonacciRecursive(number uint) uint {
	if number == 0 || number == 1 {
		return number
	}
	return FibonacciRecursive(number-1) + FibonacciRecursive(number-2)
}

// Much more effective algorithm, in fact, O(n).
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
