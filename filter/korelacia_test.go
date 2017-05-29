package filter

import (
	"testing"

	"github.com/melias122/engine/engine"
)

func TestFilterKorelacia(t *testing.T) {
	tests := []struct {
		k engine.Kombinacia
		f Filter
		w bool
	}{
		{engine.Kombinacia{1, 14, 15, 17, 19}, NewFilterKorelacia(0.0, 1.0, engine.Kombinacia{2, 7, 13, 32, 35}, 5, 35), true},   // "0.34137300"
		{engine.Kombinacia{1, 14, 15, 17, 19}, NewFilterKorelacia(0.0, 0.34, engine.Kombinacia{2, 7, 13, 32, 35}, 5, 35), true},  // "0.34137300"
		{engine.Kombinacia{1, 14, 15, 17, 19}, NewFilterKorelacia(0.41, 1.0, engine.Kombinacia{2, 7, 13, 32, 35}, 5, 35), false}, // "0.34137300"
		{engine.Kombinacia{1, 2, 3, 4, 5}, NewFilterKorelacia(1.0, 1.0, engine.Kombinacia{1, 2, 3, 4, 5}, 5, 35), true},
		{engine.Kombinacia{1, 2, 3, 4}, NewFilterKorelacia(1.0, 1.0, engine.Kombinacia{1, 2, 3, 4}, 5, 35), true},
	}
	for _, test := range tests {
		ok := test.f.Check(test.k)
		if ok != test.w {
			t.Errorf("Excepted: (%v), Got: (%v) %v", test.w, ok, test)
		}
	}
}
