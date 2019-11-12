package Lesson1

import "testing"

func TestIsPrime(t *testing.T) {
	var number int
	var result bool
	number = 37
	result = IsPrime(number)
	if !result {
		t.Error("Wrong answer! Expected true, got false")
	}
}
