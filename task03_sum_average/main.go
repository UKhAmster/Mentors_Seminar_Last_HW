// Задание 3. Сумма и среднее от 1 до N.
package main

import "fmt"

// sumAndAverage возвращает сумму чисел от 1 до n и их среднее значение.
func sumAndAverage(n int) (int, float64) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum, float64(sum) / float64(n)
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

	sum, avg := sumAndAverage(n)
	fmt.Printf("Сумма чисел от 1 до %d: %d\n", n, sum)
	fmt.Printf("Среднее значение: %.2f\n", avg)
}
