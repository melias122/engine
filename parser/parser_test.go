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

func TestParseDigit(t *testing.T) {
	tests := []struct {
		s string
		d int
		e bool
	}{
		{"", 0, true},
		{"foo", 0, true},
		{"0", 0, false},
		// {"-1", -1, false},
		{"1", 1, false},
		// {"2e2", 2e2, false},
	}
	for _, test := range tests {
		r := strings.NewReader(test.s)
		p := NewParser(r)
		d, err := p.ParseInt()
		if err != nil && !test.e {
			t.Log("Should not err: ", test)
		}
		if d != test.d {
			t.Fatal(err)
		}
	}
}

func TestParseColon(t *testing.T) {
	tests := []struct {
		s string
		d []int
	}{
		{"3-2", nil},
		{"1-3", []int{1, 2, 3}},
		{"  1  ,  2  ,4    ", []int{1, 2, 4}},
		{"1,2,4-5", []int{1, 2, 4, 5}},

		{"", nil},
		{"1,2,4-5,", nil},
		{"1,2, 4 -5, foo    ", nil},
	}
	for _, test := range tests {
		r := strings.NewReader(test.s)
		p := NewParser(r)
		d, err := p.ParseInts()
		if !reflect.DeepEqual(test.d, d) {
			t.Fatal(err, test, d)
		}
	}
}

func TestParseSemiColon(t *testing.T) {
	tests := []struct {
		s string
		r MapInts
	}{
		{"1:1", MapInts{1: Ints{1}}},
		{"1:1,2-4;2:3;5:8", MapInts{1: Ints{1, 2, 3, 4}, 2: Ints{3}, 5: Ints{8}}},

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
		p := NewParser(r)
		res, err := p.ParseMapInts()
		if !reflect.DeepEqual(test.r, res) {
			t.Fatal(err, test, res)
		}
	}
}
