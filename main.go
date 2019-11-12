package main

import (
	"StepikCourseAlgorithms/Lesson1"
	"fmt"
)

func main() {
	fmt.Println("Started")
	var number uint
	fmt.Scan(&number)
	fmt.Println(Lesson1.ToPrimeMultipliers(number))
}
