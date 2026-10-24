package main

import "testing"

func TestMinMax(t *testing.T) {
	tests := []struct {
		name             string
		numbers          []float64
		wantMin, wantMax float64
	}{
		{"одно число", []float64{5}, 5, 5},
		{"разные числа", []float64{3, -2, 8.5, 0}, -2, 8.5},
		{"все отрицательные", []float64{-7, -1, -3}, -7, -1},
		{"одинаковые", []float64{4, 4, 4}, 4, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMin, gotMax := minMax(tt.numbers)
			if gotMin != tt.wantMin || gotMax != tt.wantMax {
				t.Errorf("minMax(%v) = (%g, %g), ожидалось (%g, %g)", tt.numbers, gotMin, gotMax, tt.wantMin, tt.wantMax)
			}
		})
	}
}
