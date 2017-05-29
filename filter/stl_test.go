package filter

import (
	"testing"

	"github.com/melias122/engine/engine"
)

func TestFilterSTL(t *testing.T) {
	cisla := engine.Nums{
		engine.NewNum(1, 5, 35),
		engine.NewNum(2, 5, 35),
		engine.NewNum(3, 5, 35),
	}
	tests := []struct {
		k engine.Kombinacia
		f Filter
		w bool
	}{
		{engine.Kombinacia{1, 2, 3}, NewFilterSTL1(0.0, 1.0, cisla, 3), true},
		{engine.Kombinacia{1, 2}, NewFilterSTL1(0, -0.99, cisla, 3), false},
		{engine.Kombinacia{1, 2}, NewFilterSTL1(0.1, 1, cisla, 3), true},
		{engine.Kombinacia{1, 2, 3}, NewFilterSTL1(0.101, 1, cisla, 3), false},
	}
	for _, test := range tests {
		ok := test.f.Check(test.k)
		if ok != test.w {
			t.Errorf("Excepted: (%v), Got: (%v)", test.w, ok)
		}
	}
}
