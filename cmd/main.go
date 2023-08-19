package main

import (
	"calculator/internal"
	"calculator/internal/cheker"
	"calculator/internal/writter"
	"fmt"
	"log"
)

var rim map[string]int

func main() {
	var firstNumber string
	var secondNumber string
	var operationSymbol string
	var fNumber int
	var sNumber int
	var ok bool
	var isArabic bool

	var checkFunc func(string) (int, bool)

	fmt.Scan(&firstNumber, &operationSymbol, &secondNumber)

	if fNumber, ok = cheker.CheckIsArabic(firstNumber); ok {
		checkFunc = cheker.CheckIsArabic
		isArabic = true
	} else {
		if fNumber, ok = cheker.CheckIsRim(firstNumber); ok {
			isArabic = false
			checkFunc = cheker.CheckIsRim
		} else {
			log.Fatal("not number")
		}
	}

	sNumber, ok = checkFunc(secondNumber)

	if !ok {
		log.Fatal("wrong second number")
	}

	f := internal.GetOperationFunc(operationSymbol, isArabic)

	writter.WriteOut(f(fNumber, sNumber), isArabic)

}
