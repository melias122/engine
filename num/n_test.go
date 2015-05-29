package num

import "testing"

func TestZero(t *testing.T) {
	n := Zero(5, 35)
	if n.x != 0 {
		t.Fatalf("Expected: (0), Have: (%d)", n.x)
	}
	e := "0 0 0 0 0 0 0 0 0 0"
	if n.C().String() != e {
		t.Fatalf("Expected: (%s), Have: (%s)", e, n.C())
	}
}

func TestNew(t *testing.T) {
	n := New(1, 5, 35)
	if n.x != 1 {
		t.Fatalf("Expected: (1), Have: (%d)", n.x)
	}
}

// func TestInc1(t *testing.T) {
// 	n := New(1, 5, 35)
// }

// func TestInc2(t *testing.T) {
// 	n := New(1, 5, 35)
// }

// func TestReset2(t *testing.T) {
// 	n := New(1, 5, 35)
// }
