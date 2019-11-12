package main

import (
	"StepikCourseAlgorithms/Lesson1"
	"fmt"
)

func main() {
	fmt.Println("Started")
	var number int
	fmt.Scan(&number)
	//fmt.Println(Lesson1.FibonacciNonRecursive(number))
	fmt.Println(Lesson1.IsPrime(number))
}
