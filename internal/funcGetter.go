package internal

import (
	"calculator/internal/mymath"
	"log"
)

func GetOperationFunc(operationSymbol string, isArabic bool) func(a int, b int) int {
	var f func(a, b int) int
	switch operationSymbol {
	case "+":
		f = mymath.AddInt
	case "-":
		if isArabic {
			f = mymath.SubtractInt
		} else {
			f = mymath.SubtractRim
		}
	case "/":
		f = mymath.DivideInt
	case "*":
		f = mymath.MultiplyInt
	default:
		log.Fatalf("unknown operation symbol")
	}
	return f
}
