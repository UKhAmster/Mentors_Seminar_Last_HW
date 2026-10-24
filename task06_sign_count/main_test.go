package main

import "testing"

func TestCountSigns(t *testing.T) {
	tests := []struct {
		name    string
		numbers []float64
		want    signCounts
	}{
		{"смешанные", []float64{1, -2, 0, 3.5, -0.1, 0}, signCounts{positive: 2, negative: 2, zero: 2}},
		{"только положительные", []float64{1, 2, 3}, signCounts{positive: 3}},
		{"только нули", []float64{0, 0}, signCounts{zero: 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSigns(tt.numbers); got != tt.want {
				t.Errorf("countSigns(%v) = %+v, ожидалось %+v", tt.numbers, got, tt.want)
			}
		})
	}
}
