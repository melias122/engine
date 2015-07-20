package komb

import "testing"

func TestString(t *testing.T) {
	k := Kombinacia{1, 2, 3, 4, 5}
	if k.String() != "1 2 3 4 5" {
		t.Errorf("Excepted: (1 2 3 4 5), Got: (%s)", k.String())
	}
}
