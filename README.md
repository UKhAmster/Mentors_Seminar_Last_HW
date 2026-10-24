# Основы Go: 10 заданий

Каждое задание — отдельная программа в своей папке. Логика вынесена в функции, которые покрыты табличными тестами.

| № | Папка | Задание |
|---|-------|---------|
| 1 | [`task01_calculator`](task01_calculator/main.go) | Калькулятор с проверкой деления на ноль |
| 2 | [`task02_number_analysis`](task02_number_analysis/main.go) | Анализ числа: знак, чётность, делимость на 3 и 5 |
| 3 | [`task03_sum_average`](task03_sum_average/main.go) | Сумма и среднее чисел от 1 до N |
| 4 | [`task04_multiplication_table`](task04_multiplication_table/main.go) | Таблица умножения через функцию `printTable(n)` |
| 5 | [`task05_min_max`](task05_min_max/main.go) | Максимум и минимум из N чисел |
| 6 | [`task06_sign_count`](task06_sign_count/main.go) | Количество положительных, отрицательных чисел и нулей |
| 7 | [`task07_filtered_sum`](task07_filtered_sum/main.go) | Сумма чётных чисел от 1 до N, делящихся на 3 |
| 8 | [`task08_prime`](task08_prime/main.go) | Проверка простого числа `isPrime(n int) bool` |
| 9 | [`task09_atm`](task09_atm/main.go) | Мини-банкомат: баланс, пополнение, снятие |
| 10 | [`task10_sequence_stats`](task10_sequence_stats/main.go) | Статистика последовательности |

## Запуск

Нужен Go 1.22+.

```bash
go run ./task01_calculator
go run ./task08_prime
```

Ввод можно передать и через пайп:

```bash
echo "7 * 6" | go run ./task01_calculator
echo "5 4 -3 10 7 0" | go run ./task10_sequence_stats
```

## Тесты

```bash
go test ./...
```
