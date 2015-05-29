package num

import (
	"fmt"
	"testing"
)

func TestInc(t *testing.T) {
	ph := newph(1, 1, 5, 35)
	for i := 1; i <= 122; i++ {
		ph.inc()
	}

	r := fmt.Sprintf("%d", ph.p)
	w := "122"
	if r != w {
		t.Fatalf("Expected: (%s), Have: (%s)", w, r)
	}

	r = fmt.Sprintf("%.10f", ph.h)
	w = "0.0026306710"
	if r != w {
		t.Fatalf("Expected: (%s), Have: (%s)", w, r)
	}
}

func TestReset(t *testing.T) {
	ph := newph(1, 1, 5, 35)
	for i := 1; i <= 10; i++ {
		ph.inc()
	}
	ph.reset()

	r := ph.p
	w := 0
	if r != w {
		t.Fatalf("Expected: (%d), Have: (%d)", w, r)
	}

	r2 := ph.h
	w2 := 0.0
	if r != w {
		t.Fatalf("Expected: (%f), Have: (%f)", w2, r2)
	}
}
