package psl

import (
	"testing"
)

func TestNtica(t *testing.T) {
	tests := []struct {
		t []byte
		w string
	}{
		{[]byte{}, ""},
		{[]byte{1}, "1"},
		{[]byte{1, 3}, "2 0"},
		{[]byte{1, 3, 5}, "3 0 0"},
		{[]byte{1, 3, 5, 7}, "4 0 0 0"},
		{[]byte{1, 3, 5, 7, 9}, "5 0 0 0 0"},

		{[]byte{1, 2, 3, 4, 5}, "0 0 0 0 1"},
		{[]byte{1, 2, 3, 4, 7}, "1 0 0 1 0"},
		{[]byte{1, 2, 3, 5, 6}, "0 1 1 0 0"},
		{[]byte{1, 3, 4, 5, 9}, "2 0 1 0 0"},
		{[]byte{1, 3, 4, 6, 7}, "1 2 0 0 0"},
		{[]byte{1, 3, 5, 7, 8}, "3 1 0 0 0"},
		{[]byte{1, 3, 5, 7, 8, 10, 12, 13, 14, 15}, "4 1 0 1 0 0 0 0 0 0"},
	}
	for _, test := range tests {
		tica := Ntica(test.t)
		if tica.String() != test.w {
			t.Fatalf("Excepted: (%s), Have: (%s)", test.w, tica)
		}
	}
}

func TestNticaPozicie(t *testing.T) {
	tests := []struct {
		t Kombinacia
		w string
	}{
		{Kombinacia{}, ""},
		{Kombinacia{1}, "0"},
		{Kombinacia{1, 3}, "0 0"},
		{Kombinacia{1, 3, 5}, "0 0 0"},
		{Kombinacia{1, 3, 5, 7}, "0 0 0 0"},
		{Kombinacia{1, 3, 5, 7, 9}, "0 0 0 0 0"},

		{Kombinacia{1, 2, 3, 4, 5}, "1 1 1 1 1"},
		{Kombinacia{1, 2, 3, 4, 7}, "1 1 1 1 0"},
		{Kombinacia{1, 2, 3, 5, 6}, "1 1 1 1 1"},
		{Kombinacia{1, 3, 4, 5, 9}, "0 1 1 1 0"},
		{Kombinacia{1, 3, 4, 6, 7}, "0 1 1 1 1"},
		{Kombinacia{1, 3, 5, 7, 8}, "0 0 0 1 1"},
		{Kombinacia{1, 2, 4, 6, 7, 8, 10, 12, 13, 14}, "1 1 0 1 1 1 0 1 1 1"},
	}

	for _, test := range tests {
		p := Kombinacia(NticaPozicie(test.t))
		if p.String() != test.w {
			t.Fatalf("Excepted: (%s), Have: (%s) Kombinacia: (%s)", test.w, p, test.t)
		}
	}
}

func TestNticaSucet(t *testing.T) {
	tests := []struct {
		k     Kombinacia
		sucet nticaSS
	}{
		{Kombinacia{1, 3, 5}, nticaSS{}},
		{Kombinacia{1, 2, 3}, nticaSS{6}},
		{Kombinacia{1, 2, 5, 7, 8}, nticaSS{3, 15}},
	}
	for _, test := range tests {
		sucet := NticaSucet(test.k)
		if sucet.String() != test.sucet.String() {
			t.Errorf("Expected: %s, Got: %s", test.sucet.String(), sucet.String())
		}
	}
}

func TestNticaSucin(t *testing.T) {
	tests := []struct {
		k     Kombinacia
		sucin nticaSS
	}{
		{Kombinacia{1, 3, 5}, nticaSS{}},
		{Kombinacia{1, 2, 3}, nticaSS{6}},
		{Kombinacia{1, 2, 5, 7, 8}, nticaSS{2, 20}},
	}
	for _, test := range tests {
		sucin := NticaSucin(test.k)
		if sucin.String() != test.sucin.String() {
			t.Errorf("Expected: %s, Got: %s", test.sucin.String(), sucin.String())
		}
	}
}

func TestFilterNtica(t *testing.T) {
	tests := []struct {
		k Kombinacia
		f Filter
		w bool
	}{
		{Kombinacia{1, 2, 3, 4, 5}, NewNtica(5, Tica{0, 0, 0, 0, 1}), true},
		{Kombinacia{1, 2, 3, 4, 6}, NewNtica(5, Tica{1, 0, 0, 1, 0}), true},
		{Kombinacia{1, 2, 3, 5, 6}, NewNtica(5, Tica{0, 1, 1, 0, 0}), true},
		{Kombinacia{1, 2, 3, 5, 7}, NewNtica(5, Tica{2, 0, 1, 0, 0}), true},
		{Kombinacia{1, 2, 4, 5, 7}, NewNtica(5, Tica{1, 2, 0, 0, 0}), true},
		{Kombinacia{1, 2, 4, 6, 8}, NewNtica(5, Tica{3, 1, 0, 0, 0}), true},
		{Kombinacia{1, 3, 5, 7, 9}, NewNtica(5, Tica{5, 0, 0, 0, 0}), true},

		{Kombinacia{1, 2, 3, 4, 5}, NewNtica(5, Tica{5, 0, 0, 0, 0}), false},
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
		k Kombinacia
	}{
		{NewStlNtica(4, Tica{4, 0, 0, 0}, []byte{0, 0, 0, 0}), Kombinacia{1, 3, 5, 7}},
		{NewStlNtica(4, Tica{2, 1, 0, 0}, []byte{1, 1, 0, 0}), Kombinacia{1, 2, 5, 7}},
		{NewStlNtica(4, Tica{2, 1, 0, 0}, []byte{0, 1, 1, 0}), Kombinacia{1, 3, 4, 7}},
		{NewStlNtica(4, Tica{2, 1, 0, 0}, []byte{0, 0, 1, 1}), Kombinacia{1, 3, 6, 7}},
		{NewStlNtica(4, Tica{1, 0, 1, 0}, []byte{1, 1, 1, 0}), Kombinacia{1, 2, 3, 7}},
		{NewStlNtica(4, Tica{1, 0, 1, 0}, []byte{0, 1, 1, 1}), Kombinacia{1, 3, 4, 5}},
		{NewStlNtica(4, Tica{0, 0, 0, 1}, []byte{1, 1, 1, 1}), Kombinacia{2, 3, 4, 5}},
		{
			Filter: NewStlNtica(5, Tica{0, 0, 0, 0, 1}, []byte{1, 1, 1, 1, 1}),
			k:      Kombinacia{1, 2, 3, 4, 5},
		},
		{
			Filter: NewStlNtica(6, Tica{4, 1, 0, 0, 0, 0}, []byte{0, 0, 0, 0, 1, 1}),
			k:      Kombinacia{1, 3, 5, 7, 9, 10},
		},
	}
	for _, test := range tests {
		ok := test.Check(test.k)
		if !ok {
			t.Errorf("Excepted: (%v), Got: (%v)", true, ok)
		}
	}
}
