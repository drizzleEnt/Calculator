package main

import (
	"bufio"
	"calculator/internal"
	"calculator/internal/cheker"
	"calculator/internal/writter"
	"log"
	"os"
	"strings"
)

func main() {
	var firstNumber string
	var secondNumber string
	var operationSymbol string
	var fNumber int
	var sNumber int
	var ok bool
	var isArabic bool

	var checkFunc func(string) (int, bool)

	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	s := strings.Split(text, " ")

	if len(s) > 3 {
		log.Fatal("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}

	firstNumber = s[0]
	operationSymbol = s[1]
	secondNumber = strings.TrimSpace(s[2])

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
