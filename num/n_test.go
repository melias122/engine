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

func TestMax(t *testing.T) {
	tests := []struct {
		x, y, n, m int
		w          string
	}{
		{1, 1, 5, 35, "46376"},
		{1, 2, 5, 35, "0"},
		{5, 1, 5, 35, "27405"},
		{5, 2, 5, 35, "16240"},
		{5, 3, 5, 35, "2610"},
		{5, 4, 5, 35, "120"},
		{5, 5, 5, 35, "1"},
	}
	for _, x := range tests {
		r := Max(x.x, x.y, x.n, x.m).String()
		if r != x.w {
			t.Fatalf("Expected: (%s), Have: (%s)", x.w, r)
		}
	}
}
