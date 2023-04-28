package main

import (
	"fmt"
	"os"
	"strconv"
)

func getFloatInput(prompt string) (float64, error) {
	var input string
	fmt.Print(prompt)
	fmt.Scan(&input)
	return strconv.ParseFloat(input, 64)
}

func getOperatorInput(prompt string) (string, error) {
	var input string
	fmt.Print(prompt)
	fmt.Scan(&input)

	switch input {
	case "+", "-", "*", "/":
		return input, nil
	default:
		return "", fmt.Errorf("Invalid operator: %s", input)
	}
}

func applyOperator(num1, num2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 != 0 {
			return num1 / num2, nil
		}
		return 0, fmt.Errorf("Error: Division by zero!")
	}
	return 0, fmt.Errorf("Invalid operator!")
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	num1, err := getFloatInput("Enter the first number: ")
	checkError(err)

	num2, err := getFloatInput("Enter the second number: ")
	checkError(err)

	operator, err := getOperatorInput("Enter an operator (+, -, *, /): ")
	checkError(err)

	result, err := applyOperator(num1, num2, operator)
	checkError(err)

	fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}
