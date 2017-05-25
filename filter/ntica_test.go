package filter

import (
	"testing"

	"github.com/melias122/engine"
)

func TestFilterNtica(t *testing.T) {
	tests := []struct {
		k engine.Kombinacia
		f Filter
		w bool
	}{
		{engine.Kombinacia{1, 2, 3, 4, 5}, NewFilterNtica(5, engine.Tica{0, 0, 0, 0, 1}), true},
		{engine.Kombinacia{1, 2, 3, 4, 6}, NewFilterNtica(5, engine.Tica{1, 0, 0, 1, 0}), true},
		{engine.Kombinacia{1, 2, 3, 5, 6}, NewFilterNtica(5, engine.Tica{0, 1, 1, 0, 0}), true},
		{engine.Kombinacia{1, 2, 3, 5, 7}, NewFilterNtica(5, engine.Tica{2, 0, 1, 0, 0}), true},
		{engine.Kombinacia{1, 2, 4, 5, 7}, NewFilterNtica(5, engine.Tica{1, 2, 0, 0, 0}), true},
		{engine.Kombinacia{1, 2, 4, 6, 8}, NewFilterNtica(5, engine.Tica{3, 1, 0, 0, 0}), true},
		{engine.Kombinacia{1, 3, 5, 7, 9}, NewFilterNtica(5, engine.Tica{5, 0, 0, 0, 0}), true},

		{engine.Kombinacia{1, 2, 3, 4, 5}, NewFilterNtica(5, engine.Tica{5, 0, 0, 0, 0}), false},
	}
	for _, test := range tests {
		ok := test.f.Check(test.k)
		if ok != test.w {
			t.Errorf("Excepted: (%v), Got: (%v)", test.w, ok)
		}
	}
}

func TestStlNtica(t *testing.T) {
	tests := []struct {
		Filter
		k engine.Kombinacia
	}{
		{NewFilterSTLNtica(4, engine.Tica{4, 0, 0, 0}, []byte{0, 0, 0, 0}), engine.Kombinacia{1, 3, 5, 7}},
		{NewFilterSTLNtica(4, engine.Tica{2, 1, 0, 0}, []byte{1, 1, 0, 0}), engine.Kombinacia{1, 2, 5, 7}},
		{NewFilterSTLNtica(4, engine.Tica{2, 1, 0, 0}, []byte{0, 1, 1, 0}), engine.Kombinacia{1, 3, 4, 7}},
		{NewFilterSTLNtica(4, engine.Tica{2, 1, 0, 0}, []byte{0, 0, 1, 1}), engine.Kombinacia{1, 3, 6, 7}},
		{NewFilterSTLNtica(4, engine.Tica{1, 0, 1, 0}, []byte{1, 1, 1, 0}), engine.Kombinacia{1, 2, 3, 7}},
		{NewFilterSTLNtica(4, engine.Tica{1, 0, 1, 0}, []byte{0, 1, 1, 1}), engine.Kombinacia{1, 3, 4, 5}},
		{NewFilterSTLNtica(4, engine.Tica{0, 0, 0, 1}, []byte{1, 1, 1, 1}), engine.Kombinacia{2, 3, 4, 5}},
		{
			Filter: NewFilterSTLNtica(5, engine.Tica{0, 0, 0, 0, 1}, []byte{1, 1, 1, 1, 1}),
			k:      engine.Kombinacia{1, 2, 3, 4, 5},
		},
		{
			Filter: NewFilterSTLNtica(6, engine.Tica{4, 1, 0, 0, 0, 0}, []byte{0, 0, 0, 0, 1, 1}),
			k:      engine.Kombinacia{1, 3, 5, 7, 9, 10},
		},
	}
	for _, test := range tests {
		ok := test.Check(test.k)
		if !ok {
			t.Errorf("Excepted: (%v), Got: (%v)", true, ok)
		}
	}
}
