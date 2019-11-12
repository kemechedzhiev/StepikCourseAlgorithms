package Lesson1

import "testing"

func TestFibonacciRecursive(t *testing.T) {
	var number, result uint
	number = 6
	result = FibonacciRecursive(number)
	if result != 8 {
		t.Error("Wrong answer! Expected 8, got ", result)
	}
}

func TestFibonacciNonRecursive(t *testing.T) {
	var number, result uint
	number = 6
	result = FibonacciNonRecursive(number)
	if result != 8 {
		t.Error("Wrong answer! Expected 8, got ", result)
	}
}
