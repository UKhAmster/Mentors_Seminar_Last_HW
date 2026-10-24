// Задание 7. Сумма только подходящих чисел.
package main

import "fmt"

// sumEvenDivisibleBy3 возвращает сумму чисел от 1 до n, которые одновременно
// чётные и делятся на 3.
func sumEvenDivisibleBy3(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		if i%2 == 0 && i%3 == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	var n int
	fmt.Print("Введите натуральное число N: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Ошибка: нужно ввести целое число")
		return
	}
	if n < 1 {
		fmt.Println("Ошибка: N должно быть не меньше 1")
		return
	}

	fmt.Printf("Сумма чётных чисел от 1 до %d, делящихся на 3: %d\n", n, sumEvenDivisibleBy3(n))
}
