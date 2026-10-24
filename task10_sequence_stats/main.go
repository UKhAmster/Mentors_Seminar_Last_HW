// Задание 10. Статистика последовательности.
package main

import "fmt"

// stats содержит итоговую статистику по последовательности чисел.
type stats struct {
	sum     int
	average float64
	max     int
	min     int
	even    int
	odd     int
}

// calcStats считает статистику по непустому срезу целых чисел.
func calcStats(numbers []int) stats {
	s := stats{max: numbers[0], min: numbers[0]}
	for _, x := range numbers {
		s.sum += x
		if x > s.max {
			s.max = x
		}
		if x < s.min {
			s.min = x
		}
		if x%2 == 0 {
			s.even++
		} else {
			s.odd++
		}
	}
	s.average = float64(s.sum) / float64(len(numbers))
	return s
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

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Введите целое число %d: ", i+1)
		if _, err := fmt.Scan(&numbers[i]); err != nil {
			fmt.Println("Ошибка: нужно ввести целое число")
			return
		}
	}

	s := calcStats(numbers)
	fmt.Println("Сумма:", s.sum)
	fmt.Printf("Среднее значение: %.2f\n", s.average)
	fmt.Println("Максимальное значение:", s.max)
	fmt.Println("Минимальное значение:", s.min)
	fmt.Println("Чётных чисел:", s.even)
	fmt.Println("Нечётных чисел:", s.odd)
}
