package komb

import "testing"

func TestXtica(t *testing.T) {
	tests := []struct {
		m int
		t Kombinacia
		w string
	}{
		{9, Kombinacia{1, 2, 3, 4, 5}, "5"},
		{10, Kombinacia{1, 2, 3, 4, 5}, "5"},
		{11, Kombinacia{1, 2, 3, 4, 5}, "5 0"},
		{11, Kombinacia{1, 10, 11, 12, 13}, "2 3"},
		{90, Kombinacia{1, 10, 11, 12, 13}, "2 3 0 0 0 0 0 0 0"},
		{90, Kombinacia{1, 10, 11, 12, 90}, "2 2 0 0 0 0 0 0 1"},
		{90, Kombinacia{10, 20, 30, 40, 50, 60, 70, 80, 90}, "1 1 1 1 1 1 1 1 1"},
		{90, Kombinacia{9, 19, 29, 39, 49, 59, 69, 79, 89}, "1 1 1 1 1 1 1 1 1"},
		{90, Kombinacia{11, 21, 31, 41, 51, 61, 71, 81}, "0 1 1 1 1 1 1 1 1"},
	}
	for _, test := range tests {
		tica := Xtica(test.m, test.t)
		if tica.String() != test.w {
			t.Fatalf("Excepted: (%s), Have: (%s)", test.w, tica)
		}
	}
}

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
	}
	for _, test := range tests {
		tica := Ntica(test.t)
		if tica.String() != test.w {
			t.Fatalf("Excepted: (%s), Have: (%s)", test.w, tica)
		}
	}
}
