package filter

import (
	"testing"

	"github.com/melias122/engine/engine"
)

func TestFilterSmernica(t *testing.T) {
	tests := []struct {
		k engine.Kombinacia
		f Filter
		w bool
	}{
		{engine.Kombinacia{14, 16, 26, 27}, NewFilterSmernica(0.0, 1.0, 5, 35), true},         // 0.499...
		{engine.Kombinacia{14, 16, 26, 27, 30}, NewFilterSmernica(0.0, 1.0, 5, 35), true},     // 0.499...
		{engine.Kombinacia{14, 16, 26, 27, 30}, NewFilterSmernica(0.0, 0.49, 5, 35), true},    // 0.499...
		{engine.Kombinacia{14, 16, 26, 27, 30}, NewFilterSmernica(0.49, 0.5, 5, 35), true},    // 0.499...
		{engine.Kombinacia{14, 16, 26, 27, 30}, NewFilterSmernica(0.499, 0.499, 5, 35), true}, // 0.499...
		{engine.Kombinacia{14, 16, 26, 27, 30}, NewFilterSmernica(0.5, 0.5, 5, 35), true},
		{engine.Kombinacia{14, 16, 26, 27, 30}, NewFilterSmernica(0.500000001, 0.6, 5, 35), false}, // 0.499...
	}
	for _, test := range tests {
		ok := test.f.Check(test.k)
		if ok != test.w {
			t.Errorf("Excepted: (%v), Got: (%v) (test %v)", test.w, ok, test)
		}
	}
}
