package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	operation := os.Args[2]
	firstOperand, _ := strconv.Atoi(os.Args[1])
	secondOperand, _ := strconv.Atoi(os.Args[3])
	if operation == "+" {
		fmt.Println(firstOperand + secondOperand)
	}
	if operation == "x" {
		fmt.Println(firstOperand * secondOperand)
	}
}
