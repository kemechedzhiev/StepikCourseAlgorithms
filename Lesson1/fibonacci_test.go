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

func TestModNFibToM(t *testing.T) {
	var index, number uint
	index = 10
	number = 2
	result := ModNFibToM(index, number)
	if result != 1 {
		t.Error("Wrong answer! Expected 1, got ", result)
	}
}

func TestFindPisano(t *testing.T) {
	var delimiter uint
	delimiter = 4
	result, pisanoArr := findPisano(delimiter)
	if result != 6 {
		t.Error("Wrong answer! Expected 6, got ", result)
	}
	if (*pisanoArr)[0] != 0 && (*pisanoArr)[1] != 1 && (*pisanoArr)[2] != 1 &&
		(*pisanoArr)[3] != 2 && (*pisanoArr)[4] != 3 && (*pisanoArr)[5] != 1 {
		t.Error("Wrong values in array! Expected 0 1 1 2 3 1, got ", pisanoArr)
	}
}
