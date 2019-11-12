package Lesson1

import "strconv"

// TODO Get the prime multipliers of number n. They should be ranged increasingly and separated by whitespaces.
// T(n) = O(n). M(n) = O(1)
func ToPrimeMultipliers(number uint) string {
	var counter uint
	var answer string
	counter = 2
	for counter <= number {
		if number%counter == 0 {
			answer += strconv.Itoa(int(counter))
			answer += " "
			number /= counter
			counter = 1
		}
		counter++
	}
	return answer
}
