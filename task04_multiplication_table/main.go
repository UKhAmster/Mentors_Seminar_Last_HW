// Задание 4. Таблица умножения через функцию.
package main

import (
	"fmt"
	"io"
	"os"
)

// writeTable пишет в w таблицу умножения числа n на числа от 1 до 10.
func writeTable(w io.Writer, n int) {
	for i := 1; i <= 10; i++ {
		fmt.Fprintf(w, "%d x %2d = %d\n", n, i, n*i)
	}
}

// printTable выводит таблицу умножения числа n на экран.
func printTable(n int) {
	writeTable(os.Stdout, n)
}

func main() {
	var n int
	fmt.Print("Введите число: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Ошибка: нужно ввести целое число")
		return
	}
	printTable(n)
}
