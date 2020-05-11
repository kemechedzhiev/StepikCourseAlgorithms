package main

import (
	"StepikCourseAlgorithms/Lesson1"
	"StepikCourseAlgorithms/Lesson2"
	"fmt"
)

func main() {
	fmt.Println("Started")
	var number uint
	//var base float64
	//fmt.Scan(&base)
	fmt.Scan(&number)
	//fmt.Println(Lesson1.FibonacciNonRecursive(number))
	//fmt.Println(Lesson1.IsPrime(number))
	//fmt.Println(Lesson1.Power(base, number))
	fmt.Println(Lesson1.ToPrimeMultipliers(number))
	fmt.Println(Lesson2.ArrayLinearSearch())
}
