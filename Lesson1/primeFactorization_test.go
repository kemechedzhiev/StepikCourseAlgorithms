package Lesson1

import "testing"

func TestToPrimeMultipliers(t *testing.T) {
	var number uint = 75
	result := ToPrimeMultipliers(number)
	if result != "3 5 5 " {
		t.Error("Wrong answer. Expected 3 5 5, got ", result)
	}
}
