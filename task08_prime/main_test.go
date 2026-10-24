package main

import "testing"

func TestIsPrime(t *testing.T) {
	primes := []int{2, 3, 5, 7, 11, 13, 97, 7919}
	notPrimes := []int{-7, 0, 1, 4, 9, 15, 25, 49, 100, 7917}

	for _, n := range primes {
		if !isPrime(n) {
			t.Errorf("isPrime(%d) = false, ожидалось true", n)
		}
	}
	for _, n := range notPrimes {
		if isPrime(n) {
			t.Errorf("isPrime(%d) = true, ожидалось false", n)
		}
	}
}
