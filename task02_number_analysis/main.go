// Задание 2. Анализ числа.
package main

import "fmt"

// signDescription сообщает, положительное число, отрицательное или ноль.
func signDescription(n int) string {
	if n > 0 {
		return "положительное"
	} else if n < 0 {
		return "отрицательное"
	}
	return "ноль (ни положительное, ни отрицательное)"
}

// parityDescription сообщает, чётное число или нечётное.
func parityDescription(n int) string {
	if n%2 == 0 {
		return "чётное"
	}
	return "нечётное"
}

// divisibilityDescription сообщает, делится ли число на 3 и на 5.
func divisibilityDescription(n int) string {
	by3 := n%3 == 0
	by5 := n%5 == 0
	if by3 && by5 {
		return "делится и на 3, и на 5"
	} else if by3 {
		return "делится на 3, но не делится на 5"
	} else if by5 {
		return "делится на 5, но не делится на 3"
	}
	return "не делится ни на 3, ни на 5"
}

func main() {
	var n int
	fmt.Print("Введите целое число N: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Ошибка: нужно ввести целое число")
		return
	}

	fmt.Println("Знак:", signDescription(n))
	fmt.Println("Чётность:", parityDescription(n))
	fmt.Println("Делимость:", divisibilityDescription(n))
}
