// Задание 6. Количество положительных и отрицательных чисел.
package main

import "fmt"

// signCounts хранит количество положительных, отрицательных чисел и нулей.
type signCounts struct {
	positive int
	negative int
	zero     int
}

// countSigns подсчитывает положительные, отрицательные числа и нули.
func countSigns(numbers []float64) signCounts {
	var c signCounts
	for _, x := range numbers {
		if x > 0 {
			c.positive++
		} else if x < 0 {
			c.negative++
		} else {
			c.zero++
		}
	}
	return c
}

func main() {
	var n int
	fmt.Print("Введите количество чисел N: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Ошибка: нужно ввести целое число")
		return
	}
	if n < 1 {
		fmt.Println("Ошибка: N должно быть не меньше 1")
		return
	}

	numbers := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Введите число %d: ", i+1)
		if _, err := fmt.Scan(&numbers[i]); err != nil {
			fmt.Println("Ошибка: нужно ввести число")
			return
		}
	}

	c := countSigns(numbers)
	fmt.Println("Положительных чисел:", c.positive)
	fmt.Println("Отрицательных чисел:", c.negative)
	fmt.Println("Нулей:", c.zero)
}
