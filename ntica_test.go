package psl

import (
	"testing"

	// "github.com/melias122/psl/komb"
)

func TestNtica(t *testing.T) {
	tests := []struct {
		k komb.Kombinacia
		f Filter
		w bool
	}{
		{komb.Kombinacia{1, 2, 3, 4, 5}, NewNtica(5, komb.Tica{0, 0, 0, 0, 1}), true},
		{komb.Kombinacia{1, 2, 3, 4, 6}, NewNtica(5, komb.Tica{1, 0, 0, 1, 0}), true},
		{komb.Kombinacia{1, 2, 3, 5, 6}, NewNtica(5, komb.Tica{0, 1, 1, 0, 0}), true},
		{komb.Kombinacia{1, 2, 3, 5, 7}, NewNtica(5, komb.Tica{2, 0, 1, 0, 0}), true},
		{komb.Kombinacia{1, 2, 4, 5, 7}, NewNtica(5, komb.Tica{1, 2, 0, 0, 0}), true},
		{komb.Kombinacia{1, 2, 4, 6, 8}, NewNtica(5, komb.Tica{3, 1, 0, 0, 0}), true},
		{komb.Kombinacia{1, 3, 5, 7, 9}, NewNtica(5, komb.Tica{5, 0, 0, 0, 0}), true},

		{komb.Kombinacia{1, 2, 3, 4, 5}, NewNtica(5, komb.Tica{5, 0, 0, 0, 0}), false},
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
		k komb.Kombinacia
	}{
		{NewStlNtica(4, komb.Tica{4, 0, 0, 0}, []byte{0, 0, 0, 0}), komb.Kombinacia{1, 3, 5, 7}},
		{NewStlNtica(4, komb.Tica{2, 1, 0, 0}, []byte{1, 1, 0, 0}), komb.Kombinacia{1, 2, 5, 7}},
		{NewStlNtica(4, komb.Tica{2, 1, 0, 0}, []byte{0, 1, 1, 0}), komb.Kombinacia{1, 3, 4, 7}},
		{NewStlNtica(4, komb.Tica{2, 1, 0, 0}, []byte{0, 0, 1, 1}), komb.Kombinacia{1, 3, 6, 7}},
		{NewStlNtica(4, komb.Tica{1, 0, 1, 0}, []byte{1, 1, 1, 0}), komb.Kombinacia{1, 2, 3, 7}},
		{NewStlNtica(4, komb.Tica{1, 0, 1, 0}, []byte{0, 1, 1, 1}), komb.Kombinacia{1, 3, 4, 5}},
		{NewStlNtica(4, komb.Tica{0, 0, 0, 1}, []byte{1, 1, 1, 1}), komb.Kombinacia{2, 3, 4, 5}},
		{
			Filter: NewStlNtica(5, komb.Tica{0, 0, 0, 0, 1}, []byte{1, 1, 1, 1, 1}),
			k:      komb.Kombinacia{1, 2, 3, 4, 5},
		},
		{
			Filter: NewStlNtica(6, komb.Tica{4, 1, 0, 0, 0, 0}, []byte{0, 0, 0, 0, 1, 1}),
			k:      komb.Kombinacia{1, 3, 5, 7, 9, 10},
		},
	}
	for _, test := range tests {
		ok := test.Check(test.k)
		if !ok {
			t.Errorf("Excepted: (%v), Got: (%v)", true, ok)
		}
	}
}
