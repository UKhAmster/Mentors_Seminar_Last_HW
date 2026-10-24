package main

import "testing"

func TestSumEvenDivisibleBy3(t *testing.T) {
	cases := map[int]int{
		1:  0,
		5:  0,
		6:  6,
		20: 6 + 12 + 18,
		30: 6 + 12 + 18 + 24 + 30,
	}
	for n, want := range cases {
		if got := sumEvenDivisibleBy3(n); got != want {
			t.Errorf("sumEvenDivisibleBy3(%d) = %d, ожидалось %d", n, got, want)
		}
	}
}
