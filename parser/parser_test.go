package parser

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
		d, err := p.ParseInts(test.zhoda)
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
		res, err := p.ParseMapInts(nil)
		if !reflect.DeepEqual(test.r, res) {
			t.Fatal(err, test, res)
		}
	}
}
