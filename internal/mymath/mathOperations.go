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

func AddRim(a, b int) int {
	res := a + b
	if res > 10 {
		log.Fatal("too big numer")
	}
	return res
}

func SubtractRim(a, b int) int {
	res := a - b
	if res < 0 {
		log.Fatal("rim nums cant be negative")
	}
	return res
}

func MultiplyRim(a, b int) int {
	res := a * b
	if res > 10 {
		log.Fatal("too big numer")
	}
	return res
}

func DivideRim(a, b int) int {
	return a / b
}
