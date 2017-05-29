package engine

import (
	"testing"
)

func TestNtica(t *testing.T) {
	tests := []struct {
		t Kombinacia
		w string
	}{
		{Kombinacia{}, ""},
		{Kombinacia{1}, "1"},
		{Kombinacia{1, 3}, "2 0"},
		{Kombinacia{1, 3, 5}, "3 0 0"},
		{Kombinacia{1, 3, 5, 7}, "4 0 0 0"},
		{Kombinacia{1, 3, 5, 7, 9}, "5 0 0 0 0"},

		{Kombinacia{1, 2, 3, 4, 5}, "0 0 0 0 1"},
		{Kombinacia{1, 2, 3, 4, 7}, "1 0 0 1 0"},
		{Kombinacia{1, 2, 3, 5, 6}, "0 1 1 0 0"},
		{Kombinacia{1, 3, 4, 5, 9}, "2 0 1 0 0"},
		{Kombinacia{1, 3, 4, 6, 7}, "1 2 0 0 0"},
		{Kombinacia{1, 3, 5, 7, 8}, "3 1 0 0 0"},
		{Kombinacia{1, 3, 5, 7, 8, 10, 12, 13, 14, 15}, "4 1 0 1 0 0 0 0 0 0"},
	}
	for _, test := range tests {
		tica := NewNtica(test.t)
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
		p := NticaPozicie(test.t)
		if bytesToString(p) != test.w {
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
