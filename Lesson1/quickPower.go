package Lesson1

// T(n) = O(logn), because we divide n to 2 while we have the power of 2 which is greater than n itself. M(n) = O(1)
func Power(base float64, number int) float64 {
	var result, baseInDegreeOfTwo float64
	result = 1
	baseInDegreeOfTwo = base
	for number > 0 {
		if number&1 == 1 {
			result *= baseInDegreeOfTwo
		}
		baseInDegreeOfTwo *= baseInDegreeOfTwo
		number = number >> 1
	}
	return result
}
