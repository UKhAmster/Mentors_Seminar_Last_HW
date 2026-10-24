// Задание 9. Мини-банкомат.
package main

import (
	"errors"
	"fmt"
)

var (
	errNonPositiveAmount = errors.New("сумма должна быть больше нуля")
	errInsufficientFunds = errors.New("недостаточно средств на балансе")
)

// deposit возвращает новый баланс после пополнения на amount.
func deposit(balance, amount float64) (float64, error) {
	if amount <= 0 {
		return balance, errNonPositiveAmount
	}
	return balance + amount, nil
}

// withdraw возвращает новый баланс после снятия amount.
// Если денег не хватает, баланс не меняется.
func withdraw(balance, amount float64) (float64, error) {
	if amount <= 0 {
		return balance, errNonPositiveAmount
	}
	if amount > balance {
		return balance, errInsufficientFunds
	}
	return balance - amount, nil
}

// readAmount запрашивает у пользователя сумму операции.
func readAmount(prompt string) (float64, error) {
	var amount float64
	fmt.Print(prompt)
	if _, err := fmt.Scan(&amount); err != nil {
		return 0, fmt.Errorf("нужно ввести число: %w", err)
	}
	return amount, nil
}

func main() {
	var balance float64
	fmt.Print("Введите начальный баланс: ")
	if _, err := fmt.Scan(&balance); err != nil {
		fmt.Println("Ошибка: нужно ввести число")
		return
	}
	if balance < 0 {
		fmt.Println("Ошибка: баланс не может быть отрицательным")
		return
	}

	for {
		fmt.Println()
		fmt.Println("1. Посмотреть баланс")
		fmt.Println("2. Пополнить баланс")
		fmt.Println("3. Снять средства")
		fmt.Println("0. Выход")
		fmt.Print("Выберите действие: ")

		var choice int
		if _, err := fmt.Scan(&choice); err != nil {
			fmt.Println("Ошибка: нужно ввести номер действия")
			return
		}

		switch choice {
		case 1:
			fmt.Printf("Текущий баланс: %.2f\n", balance)
		case 2:
			amount, err := readAmount("Сумма пополнения: ")
			if err != nil {
				fmt.Println("Ошибка:", err)
				return
			}
			newBalance, err := deposit(balance, amount)
			if err != nil {
				fmt.Println("Ошибка:", err)
				continue
			}
			balance = newBalance
			fmt.Printf("Баланс пополнен. Текущий баланс: %.2f\n", balance)
		case 3:
			amount, err := readAmount("Сумма снятия: ")
			if err != nil {
				fmt.Println("Ошибка:", err)
				return
			}
			newBalance, err := withdraw(balance, amount)
			if err != nil {
				fmt.Println("Ошибка:", err)
				continue
			}
			balance = newBalance
			fmt.Printf("Выдано %.2f. Текущий баланс: %.2f\n", amount, balance)
		case 0:
			fmt.Println("До свидания!")
			return
		default:
			fmt.Println("Неизвестное действие, выберите 0–3")
		}
	}
}
