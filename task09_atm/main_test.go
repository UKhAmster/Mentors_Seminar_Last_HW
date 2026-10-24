package main

import (
	"errors"
	"testing"
)

func TestDeposit(t *testing.T) {
	tests := []struct {
		name            string
		balance, amount float64
		want            float64
		wantErr         error
	}{
		{"обычное пополнение", 100, 50, 150, nil},
		{"нулевая сумма", 100, 0, 100, errNonPositiveAmount},
		{"отрицательная сумма", 100, -10, 100, errNonPositiveAmount},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := deposit(tt.balance, tt.amount)
			if !errors.Is(err, tt.wantErr) || got != tt.want {
				t.Errorf("deposit(%g, %g) = (%g, %v), ожидалось (%g, %v)", tt.balance, tt.amount, got, err, tt.want, tt.wantErr)
			}
		})
	}
}

func TestWithdraw(t *testing.T) {
	tests := []struct {
		name            string
		balance, amount float64
		want            float64
		wantErr         error
	}{
		{"обычное снятие", 100, 30, 70, nil},
		{"снять всё", 100, 100, 0, nil},
		{"недостаточно средств", 100, 150, 100, errInsufficientFunds},
		{"нулевая сумма", 100, 0, 100, errNonPositiveAmount},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := withdraw(tt.balance, tt.amount)
			if !errors.Is(err, tt.wantErr) || got != tt.want {
				t.Errorf("withdraw(%g, %g) = (%g, %v), ожидалось (%g, %v)", tt.balance, tt.amount, got, err, tt.want, tt.wantErr)
			}
		})
	}
}
