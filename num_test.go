package psl

import "testing"

func TestNewNum(t *testing.T) {
	n := NewNum(1, 5, 35)
	if n.cislo != 1 {
		t.Fatalf("Expected: (1), Have: (%d)", n.cislo)
	}
}

func TestCislo(t *testing.T) {
	N := NewNum(23, 4, 10)
	if N.Cislo() != 23 {
		t.Errorf("Excepted (23), Got: (%d)", N.Cislo())
	}
}
