package main

import (
	"errors"
	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		name    string
		a, b    float64
		op      string
		want    float64
		wantErr error
	}{
		{"сложение", 2, 3, "+", 5, nil},
		{"вычитание", 2, 3, "-", -1, nil},
		{"умножение", 2.5, 4, "*", 10, nil},
		{"деление", 7, 2, "/", 3.5, nil},
		{"деление на ноль", 7, 0, "/", 0, errDivisionByZero},
		{"неизвестная операция", 1, 1, "%", 0, errUnknownOperator},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculate(tt.a, tt.b, tt.op)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("ошибка = %v, ожидалась %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("calculate(%g, %g, %q) = %g, ожидалось %g", tt.a, tt.b, tt.op, got, tt.want)
			}
		})
	}
}
