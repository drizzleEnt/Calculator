package internal

import (
	"calculator/internal/mymath"
	"log"
)

func GetOperationFunc(operationSymbol string, isArabic bool) func(a int, b int) int {
	var f func(a, b int) int
	switch operationSymbol {
	case "+":
		if isArabic {
			f = mymath.AddInt
		} else {
			f = mymath.AddRim
		}
	case "-":
		if isArabic {
			f = mymath.SubtractInt
		} else {
			f = mymath.SubtractRim
		}
	case "/":
		if isArabic {
			f = mymath.DivideInt
		} else {
			f = mymath.DivideRim
		}
	case "*":
		if isArabic {
			f = mymath.MultiplyInt
		} else {
			f = mymath.MultiplyRim
		}
	default:
		log.Fatalf("unknown operation symbol")
	}
	return f
}
