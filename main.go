package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	firstOperand, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Operands must be numeric.")
		return
	}

	secondOperand, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Operands must be numeric.")
		return
	}

	operation := os.Args[2]
	if operation == "+" {
		fmt.Println(firstOperand + secondOperand)
	} else if operation == "x" {
		fmt.Println(firstOperand * secondOperand)
	} else {
		fmt.Println("Only addition (+) and multiplication (x) are supported.")
	}
}
