package main

import "testing"

func TestSignDescription(t *testing.T) {
	cases := map[int]string{
		7:  "положительное",
		-4: "отрицательное",
		0:  "ноль (ни положительное, ни отрицательное)",
	}
	for n, want := range cases {
		if got := signDescription(n); got != want {
			t.Errorf("signDescription(%d) = %q, ожидалось %q", n, got, want)
		}
	}
}

func TestParityDescription(t *testing.T) {
	cases := map[int]string{4: "чётное", 7: "нечётное", -3: "нечётное", 0: "чётное"}
	for n, want := range cases {
		if got := parityDescription(n); got != want {
			t.Errorf("parityDescription(%d) = %q, ожидалось %q", n, got, want)
		}
	}
}

func TestDivisibilityDescription(t *testing.T) {
	cases := map[int]string{
		15: "делится и на 3, и на 5",
		9:  "делится на 3, но не делится на 5",
		10: "делится на 5, но не делится на 3",
		7:  "не делится ни на 3, ни на 5",
	}
	for n, want := range cases {
		if got := divisibilityDescription(n); got != want {
			t.Errorf("divisibilityDescription(%d) = %q, ожидалось %q", n, got, want)
		}
	}
}
