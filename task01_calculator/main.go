// Задание 1. Калькулятор с проверкой.
package main

import (
	"errors"
	"fmt"
)

var (
	errDivisionByZero  = errors.New("деление на ноль невозможно")
	errUnknownOperator = errors.New("неизвестная операция, допустимы: + - * /")
)

// calculate выполняет арифметическую операцию op над a и b.
func calculate(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errDivisionByZero
		}
		return a / b, nil
	default:
		return 0, errUnknownOperator
	}
}

func main() {
	var a, b float64
	var op string

	fmt.Print("Введите первое число: ")
	if _, err := fmt.Scan(&a); err != nil {
		fmt.Println("Ошибка: нужно ввести число")
		return
	}
	fmt.Print("Введите операцию (+, -, *, /): ")
	if _, err := fmt.Scan(&op); err != nil {
		fmt.Println("Ошибка: нужно ввести операцию")
		return
	}
	fmt.Print("Введите второе число: ")
	if _, err := fmt.Scan(&b); err != nil {
		fmt.Println("Ошибка: нужно ввести число")
		return
	}

	result, err := calculate(a, b, op)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Printf("%g %s %g = %g\n", a, op, b, result)
}
