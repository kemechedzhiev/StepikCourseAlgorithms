package Lesson1

// T(n) = O(sqrt(n)), because we've only got sqrt(n) iterations. M(n) = O(1), because we've got no recursion and only use local variables.
func IsPrime(number int) bool {
	if number == 1 {
		return false
	}
	for i := 2; i*i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
