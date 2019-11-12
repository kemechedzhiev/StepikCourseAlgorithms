package Lesson1

import "testing"

func TestPower(t *testing.T) {
	var base float64
	var number int
	number = 3
	base = 2.0
	result := Power(base, number)
	if result != 8.0 {
		t.Error("Wrong answer! Expected 8.0, got ", result)
	}
}
