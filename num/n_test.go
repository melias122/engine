package num

import "testing"

// func TestZero(t *testing.T) {
// 	n := Zero(5, 35)
// 	if n.x != 0 {
// 		t.Fatalf("Expected: (0), Have: (%d)", n.x)
// 	}
// 	e := "0 0 0 0 0 0 0 0 0 0"
// 	if n.C().String() != e {
// 		t.Fatalf("Expected: (%s), Have: (%s)", e, n.C())
// 	}
// }

func TestNew(t *testing.T) {
	n := New(1, 5, 35)
	if n.cislo != 1 {
		t.Fatalf("Expected: (1), Have: (%d)", n.cislo)
	}
}

// func TestMakeCopy(t *testing.T) {
// 	n := New(1, 5, 35)
// 	n.Inc(1)
// 	copy := n.MakeCopy()
// 	if !reflect.DeepEqual(n, copy) {
// 		t.Error("Copy excepted to be equal")
// 	}
// }

func TestCislo(t *testing.T) {
	N := New(23, 4, 10)
	if N.Cislo() != 23 {
		t.Errorf("Excepted (23), Got: (%d)", N.Cislo())
	}
}
