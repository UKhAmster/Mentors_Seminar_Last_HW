package main

import "testing"

func TestCalcStats(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		want    stats
	}{
		{
			"одно число",
			[]int{5},
			stats{sum: 5, average: 5, max: 5, min: 5, even: 0, odd: 1},
		},
		{
			"смешанные",
			[]int{4, -3, 10, 7, 0},
			stats{sum: 18, average: 3.6, max: 10, min: -3, even: 3, odd: 2},
		},
		{
			"все отрицательные",
			[]int{-2, -8, -5},
			stats{sum: -15, average: -5, max: -2, min: -8, even: 2, odd: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcStats(tt.numbers); got != tt.want {
				t.Errorf("calcStats(%v) = %+v, ожидалось %+v", tt.numbers, got, tt.want)
			}
		})
	}
}
