package mymath

import "log"

func AddInt(a, b int) int {
	return a + b
}

func SubtractInt(a, b int) int {
	return a - b
}

func MultiplyInt(a, b int) int {
	return a * b
}

func DivideInt(a, b int) int {
	return a / b
}

func SubtractRim(a, b int) int {
	res := a - b
	if res < 0 {
		log.Fatal("rim nums cant be negative")
	}
	return res
}
