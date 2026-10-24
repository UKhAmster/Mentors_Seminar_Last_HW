// Задание 5. Максимум из N чисел.
package main

import "fmt"

// minMax возвращает минимальное и максимальное значения непустого среза.
func minMax(numbers []float64) (float64, float64) {
	minVal, maxVal := numbers[0], numbers[0]
	for _, x := range numbers[1:] {
		if x < minVal {
			minVal = x
		}
		if x > maxVal {
			maxVal = x
		}
	}
	return minVal, maxVal
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

	minVal, maxVal := minMax(numbers)
	fmt.Printf("Максимальное значение: %g\n", maxVal)
	fmt.Printf("Минимальное значение: %g\n", minVal)
}
