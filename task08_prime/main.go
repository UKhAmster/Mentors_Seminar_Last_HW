// Задание 8. Проверка простого числа.
package main

import "fmt"

// isPrime определяет, является ли n простым числом.
// Делители перебираются только до корня из n.
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n%2 == 0 {
		return n == 2
	}
	for d := 3; d*d <= n; d += 2 {
		if n%d == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Print("Введите целое число: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Ошибка: нужно ввести целое число")
		return
	}

	if isPrime(n) {
		fmt.Printf("%d — простое число\n", n)
	} else {
		fmt.Printf("%d — не простое число\n", n)
	}
}
