package engine

import (
	"testing"
)

func TestNewCislovacka(t *testing.T) {
	tests := []struct {
		x int
		c Cislovacka
	}{
		// P N PR MC VC c19 c0 cC Cc CC
		// {0, "0 0 0 0 0 0 0 0 0 0"},
		{1, Cislovacka{0, 1, 0, 1, 0, 1, 0, 0, 0, 0}},
		{2, Cislovacka{1, 0, 1, 1, 0, 1, 0, 0, 0, 0}},
		{9, Cislovacka{0, 1, 0, 0, 1, 1, 0, 0, 0, 0}},
		{10, Cislovacka{1, 0, 0, 0, 1, 0, 1, 0, 0, 0}},
		{11, Cislovacka{0, 1, 1, 1, 0, 0, 0, 0, 0, 1}},
		{12, Cislovacka{1, 0, 0, 1, 0, 0, 0, 1, 0, 0}},
		{21, Cislovacka{0, 1, 0, 1, 0, 0, 0, 0, 1, 0}},
	}
	for _, test := range tests {
		c := NewCislovacka(Kombinacia{test.x})
		if c != test.c {
			t.Fatalf("expected %v, got: %v", test.c, c)
		}
	}
}

func TestIsP(t *testing.T) {
	test := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res := []bool{false, true, false, true, false, true, false, true, false, true}

	for i, e := range test {
		if IsP(e) != res[i] {
			t.Fail()
		}
	}
}

func TestIsN(t *testing.T) {
	test := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res := []bool{true, false, true, false, true, false, true, false, true, false}

	for i, e := range test {
		if IsN(e) != res[i] {
			t.Fail()
		}
	}
}

func TestIsPr(t *testing.T) {
	test := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res := []bool{false, true, true, false, true, false, true, false, false, false}

	for i, e := range test {
		if IsPr(e) != res[i] {
			t.Fail()
		}
	}
}

func TestIsMc(t *testing.T) {
	test := []int{1, 2, 3, 10, 15, 23, 29, 31, 32, 90}
	res := []bool{true, true, true, false, true, true, false, true, true, false}

	for i, e := range test {
		if IsMc(e) != res[i] {
			t.Log(e, " nie je Mc")
			t.Fail()
		}
	}
}

func TestIsVc(t *testing.T) {
	test := []int{2, 4, 5, 6, 7, 10, 14, 16, 29, 31, 32, 55, 56, 11, 16, 23, 30, 31}
	res := []bool{false, false, false, true, true, true, false, true, true, false, false, false, true, false, true, false, true, false}

	for i, e := range test {
		if IsVc(e) != res[i] {
			t.Log(e, " nie je Vc")
			t.Fail()
		}
	}
}

func TestIsC19(t *testing.T) {
	test := []int{1, 2, 3, 9, 10, 55}
	res := []bool{true, true, true, true, false, false}

	for i, e := range test {
		if IsC19(e) != res[i] {
			t.Fail()
		}
	}
}

func TestIsC0(t *testing.T) {
	test := []int{1, 9, 10, 11, 90}
	res := []bool{false, false, true, false, true}

	for i, e := range test {
		if IsC0(e) != res[i] {
			t.Fail()
		}
	}
}

func TestIscC(t *testing.T) {
	test := []int{1, 9, 10, 11, 12, 13, 14, 20, 21, 22, 89, 90}
	res := []bool{false, false, false, false, true, true, true, false, false, false, true, false}

	for i, e := range test {
		if IscC(e) != res[i] {
			t.Fail()
		}
	}
}

func TestIsCc(t *testing.T) {
	test := []int{1, 2, 9, 10, 11, 19, 20, 21, 22, 41, 42, 43}
	res := []bool{false, false, false, false, false, false, false, true, false, true, true, true}

	for i, e := range test {
		if IsCc(e) != res[i] {
			t.Log(e, " nie je Cc")
			t.Fail()
		}
	}
}

func TestIsCC(t *testing.T) {
	test := []int{1, 9, 10, 11, 12, 21, 22, 23}
	res := []bool{false, false, false, true, false, false, true, false}

	for i, e := range test {
		if IsCC(e) != res[i] {
			t.Fail()
		}
	}
}
