package filter

import (
	"testing"

	"github.com/melias122/psl/komb"
)

func TestZakazane(t *testing.T) {
	tests := []struct {
		k komb.Kombinacia
		f Filter
		w bool
	}{
		{komb.Kombinacia{1}, NewZakazane([]byte{2, 3}), true},
		{komb.Kombinacia{1, 2}, NewZakazane([]byte{2, 3}), false},
		{komb.Kombinacia{1, 3}, NewZakazane([]byte{2, 3}), false},
		{komb.Kombinacia{1, 4}, NewZakazane([]byte{2, 3}), true},
	}
	for _, test := range tests {
		ok := test.f.Check(test.k)
		if ok != test.w {
			t.Errorf("Excepted: (%v), Got: (%v)", test.w, ok)
		}
	}
}

func TestZakazaneStl(t *testing.T) {
	tests := []struct {
		k komb.Kombinacia
		f Filter
		w bool
	}{
		{komb.Kombinacia{1}, NewZakazaneStl(1, [][]byte{{1}}), false},
		{komb.Kombinacia{1}, NewZakazaneStl(1, [][]byte{{2, 3}}), true},
		{komb.Kombinacia{2, 3, 5}, NewZakazaneStl(3, [][]byte{{1}, {2}, {3, 5}}), false},
		{komb.Kombinacia{2, 3, 6}, NewZakazaneStl(3, [][]byte{{1}, {2}, {3, 5}}), true},
		{komb.Kombinacia{2, 3, 6}, NewZakazaneStl(3, [][]byte{{1}, {2}, {}}), true},

		{komb.Kombinacia{1}, NewZakazaneStl(1, [][]byte{{1}}), false},
		{komb.Kombinacia{2}, NewZakazaneStl(1, [][]byte{{1}}), true},
		{komb.Kombinacia{2, 3}, NewZakazaneStl(2, [][]byte{{1}, {3}}), false},
		{komb.Kombinacia{2, 4}, NewZakazaneStl(2, [][]byte{{1}, {3}}), true},
		{komb.Kombinacia{2, 4, 5}, NewZakazaneStl(3, [][]byte{{1}, {3}, {5}}), false},
		{komb.Kombinacia{2, 4, 6}, NewZakazaneStl(3, [][]byte{{1}, {3}, {5}}), true},
	}
	for i, test := range tests {
		ok := test.f.Check(test.k)
		if ok != test.w {
			t.Errorf("Excepted: (%v), Got: (%v), test: %d", test.w, ok, i+1)
		}
	}
}
