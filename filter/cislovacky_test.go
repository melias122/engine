package filter

import (
	"testing"

	"github.com/melias122/psl/komb"
	"github.com/melias122/psl/num"
)

func TestRangeCislovacky(t *testing.T) {
	tests := []struct {
		k komb.Kombinacia
		f Filter
		w bool
	}{
		{komb.Kombinacia{1}, CislovackyRange(1, 0, 0, num.P), true},
		{komb.Kombinacia{1}, CislovackyRange(1, 0, 1, num.P), true},
		{komb.Kombinacia{1}, CislovackyRange(1, 1, 1, num.P), false},
		{komb.Kombinacia{1}, CislovackyRange(3, 1, 1, num.P), true},
		{komb.Kombinacia{1, 2}, CislovackyRange(3, 1, 1, num.P), true},
		{komb.Kombinacia{1, 2, 3}, CislovackyRange(3, 1, 1, num.P), true},
		{komb.Kombinacia{1, 2, 4}, CislovackyRange(3, 1, 1, num.P), false},
		{komb.Kombinacia{1, 2, 3, 4, 5}, CislovackyRange(5, 0, 1, num.P), false},
		{komb.Kombinacia{1, 2, 3, 4, 5}, CislovackyRange(5, 0, 2, num.P), true},
		{komb.Kombinacia{1, 2, 3, 4, 5}, CislovackyRange(5, 2, 2, num.P), true},
		{komb.Kombinacia{1, 2, 3, 4, 5}, CislovackyRange(5, 2, 3, num.P), true},
		{komb.Kombinacia{1, 2, 3, 4, 5}, CislovackyRange(5, 3, 3, num.P), false},
	}
	for _, test := range tests {
		ok := test.f.Check(test.k)
		if ok != test.w {
			t.Errorf("Excepted: (%v), Got: (%v)", test.w, ok)
		}
	}
}

func TestExactCislovacky(t *testing.T) {
	tests := []struct {
		k komb.Kombinacia
		f Filter
		w bool
	}{
		{komb.Kombinacia{2}, CislovackyExact(4, []int{0, 2}, num.P), true},
		{komb.Kombinacia{2, 4}, CislovackyExact(4, []int{0, 2}, num.P), true},
		{komb.Kombinacia{2, 4, 6, 7}, CislovackyExact(4, []int{0, 2}, num.P), false},
		{komb.Kombinacia{2, 4, 6, 7}, CislovackyExact(4, []int{1, 3}, num.P), true},
		{komb.Kombinacia{2, 4, 7, 9}, CislovackyExact(4, []int{1, 3}, num.P), false},

		{komb.Kombinacia{1, 3, 7, 9}, CislovackyExact(4, []int{1, 3}, num.P), false},
	}
	for _, test := range tests {
		ok := test.f.Check(test.k)
		if ok != test.w {
			t.Errorf("Excepted: (%v), Got: (%v)", test.w, ok)
		}
	}
}
