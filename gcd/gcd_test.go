package main

import "testing"

func TestGCD_normal_1(t *testing.T) {
	got := Gcd(12, 21)
	if got != 3 {
		t.Errorf("Gcd(12, 21) = %d; want 3", got)
	}
}

func TestGCD_normal_2(t *testing.T) {
	got := Gcd(3465, 1323)
	if got != 63 {
		t.Errorf("Gcd(3465, 1323) = %d; want 63", got)
	}
}
