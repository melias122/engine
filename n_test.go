package psl

import "testing"

func TestNew(t *testing.T) {
	n := New(1, 5, 35)
	if n.cislo != 1 {
		t.Fatalf("Expected: (1), Have: (%d)", n.cislo)
	}
}

func TestCislo(t *testing.T) {
	N := New(23, 4, 10)
	if N.Cislo() != 23 {
		t.Errorf("Excepted (23), Got: (%d)", N.Cislo())
	}
}
