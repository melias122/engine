package psl

import (
	"reflect"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	// tests := []struct {
	// }{}
}

func TestParseInt(t *testing.T) {
	tests := []struct {
		s string
		d int
	}{
		{"", 0},
		{"foo", 0},
		{"0", 0},
		// {"-1", -1},
		{"1", 1},
		// {"2e2", 2e2},
	}
	for _, test := range tests {
		r := strings.NewReader(test.s)
		p := NewParser(r, 5, 35)
		d, err := p.ParseInt()
		if d != test.d {
			t.Fatal(err)
		}
	}
}

func TestParseInts(t *testing.T) {
	tests := []struct {
		s     string
		d     []int
		zhoda []byte
	}{
		// {"3-2", nil, nil},
		{"1-3", []int{1, 2, 3}, nil},
		{"  1  ,  2  ,4    ", []int{1, 2, 4}, nil},
		{"1,2,4-5", []int{1, 2, 4, 5}, nil},

		{"", nil, nil},
		{"1,2,4-5,", nil, nil},
		{"1,2, 4 -5, foo    ", nil, nil},

		{"P", nil, nil},
		{"N", nil, nil},
		{"Pr", nil, nil},
		{"Mc", nil, nil},
		{"Vc", nil, nil},
		{"C19", nil, nil},
		{"C0", nil, nil},
		{"cC", nil, nil},
		{"Cc", nil, nil},
		{"CC", nil, nil},
		{s: "p,zh,CC", d: nil, zhoda: []byte{1, 2, 3, 4, 5, 6}},
	}
	for _, test := range tests {
		r := strings.NewReader(test.s)
		p := NewParser(r, 5, 35)
		d, err := p.ParseInts()
		p.Zhoda = test.zhoda
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(test.d, d) {
			t.Fatal(err, test, d)
		}
	}
}

func TestParseMapInts(t *testing.T) {
	tests := []struct {
		s string
		r MapInts
	}{
		{"1:1", MapInts{1: Ints{1}}},
		{"1:1,2-4;2:3;5:8", MapInts{1: Ints{1, 2, 3, 4}, 2: Ints{3}, 5: Ints{8}}},
		{s: "1:P;2:N;3:Pr", r: MapInts{1: cislovacky[P], 2: cislovacky[N], 3: cislovacky[Pr]}},

		{"", nil},
		{"1-3", nil},
		{"foo", nil},
		{"1:1;", nil},
		{"1:1;2", nil},
		{"1,1;", nil},
		{";1:1", nil},
	}
	for _, test := range tests {
		r := strings.NewReader(test.s)
		p := NewParser(r, 5, 99)
		res, err := p.ParseMapInts()
		if !reflect.DeepEqual(test.r, res) {
			t.Fatal(err, test, res)
		}
	}
}

func TestParseNtica(t *testing.T) {
	tests := []struct {
		s string
		w Tica
		n int
	}{
		{s: "", w: Tica{}, n: 5},
		{s: "             \t\n   \t\t\n", w: Tica{}, n: 5},
		{s: "5", w: Tica{5, 0, 0, 0, 0}, n: 5},
		{s: "5 0 0 0 0", w: Tica{5, 0, 0, 0, 0}, n: 5},
		{s: "5 0 0 0 0 0", w: Tica{}, n: 5},
	}
	for _, test := range tests {
		n, e := ParseNtica(test.n, test.s)
		if e != nil {
			if n.String() != test.w.String() {
				t.Errorf("Expected: %s, Got: %s", test.w, n)
			}
		}
	}
}

func TestParseXtica(t *testing.T) {
	tests := []struct {
		s    string
		n, m int
		w    Tica
		e    bool
	}{
		{n: 5, m: 35, s: "", w: Tica{}, e: true},
		{n: 5, m: 35, s: "    ", w: Tica{}, e: true},
		{n: 5, m: 35, s: "    \t\t\t\t\t\t \n\n  \t      ", w: Tica{}, e: true},
		{n: 5, m: 35, s: "1 2", e: true}, // 1+2 != 5
		{n: 5, m: 35, s: "1 2 0 0 1", e: true},
		{n: 5, m: 35, s: "1 2 2 2", e: true},
		{n: 5, m: 35, s: "5,0,0", e: true},
		{n: 5, m: 35, s: "5;", e: true},

		{n: 5, m: 35, s: "5", w: Tica{5, 0, 0, 0}},
		{n: 5, m: 35, s: "5 ", w: Tica{5, 0, 0, 0}},
		{n: 5, m: 35, s: "3 2", w: Tica{3, 2, 0, 0}},
		{n: 5, m: 35, s: "1 2 0 2", w: Tica{1, 2, 0, 2}},
	}
	for _, test := range tests {
		x, e := ParseXtica(test.n, test.m, test.s)
		if e != nil {
			if x.String() != test.w.String() {
				t.Errorf("Expected: %s, Got: %s", test.w, x)
			}
		} else {
			if test.e {
				t.Errorf("Expected: error (%s)", test.s)
			}
		}
	}
}
