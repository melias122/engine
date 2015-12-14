package psl

import (
	"testing"
)

func BenchmarkPlus(b *testing.B) {
	c0 := Cislovacky{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	c1 := Cislovacky{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	for i := 0; i < b.N; i++ {
		c0.Plus(c1)
	}
}

func BenchmarkXP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPr(90)
	}
}

func TestNewCsilovacky(t *testing.T) {
	tests := []struct {
		x int
		w string
	}{
		// P N PR MC VC c19 c0 cC Cc CC
		// {0, "0 0 0 0 0 0 0 0 0 0"},
		{1, "0 1 0 1 0 1 0 0 0 0"},
		{2, "1 0 1 1 0 1 0 0 0 0"},
		{9, "0 1 0 0 1 1 0 0 0 0"},
		{10, "1 0 0 0 1 0 1 0 0 0"},
		{11, "0 1 1 1 0 0 0 0 0 1"},
		{12, "1 0 0 1 0 0 0 1 0 0"},
		{21, "0 1 0 1 0 0 0 0 1 0"},
	}
	for _, test := range tests {
		r := NewCislovacky(test.x)
		rs := r.String()
		if rs != test.w {
			t.Fatalf("Expected: (%s), Have: (%s)", test.w, rs)
		}
	}
}

func TestPlus(t *testing.T) {
	var test struct {
		in      []Cislovacky
		result  Cislovacky
		exepted Cislovacky
	}
	test.in = []Cislovacky{
		Cislovacky{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		Cislovacky{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		Cislovacky{0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	}
	test.exepted = Cislovacky{1, 3, 5, 7, 9, 11, 13, 15, 17, 20}

	for _, i := range test.in {
		test.result.Plus(i)
	}

	if test.result != test.exepted {
		t.Log("got: ", test.result, "expected: ", test.exepted)
		t.Fail()
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

func TestFilterCislovackyRange(t *testing.T) {
	tests := []struct {
		k Kombinacia
		f Filter
		w bool
	}{
		{Kombinacia{1}, NewFilterCislovackyRange(0, 0, P, 1), true},
		{Kombinacia{1}, NewFilterCislovackyRange(0, 1, P, 1), true},
		{Kombinacia{1}, NewFilterCislovackyRange(1, 1, P, 1), false},
		{Kombinacia{1}, NewFilterCislovackyRange(1, 1, P, 3), true},
		{Kombinacia{1, 2}, NewFilterCislovackyRange(1, 1, P, 3), true},
		{Kombinacia{1, 2, 3}, NewFilterCislovackyRange(1, 1, P, 3), true},
		{Kombinacia{1, 2, 4}, NewFilterCislovackyRange(1, 1, P, 3), false},
		{Kombinacia{1, 2, 3, 4, 5}, NewFilterCislovackyRange(0, 1, P, 5), false},
		{Kombinacia{1, 2, 3, 4, 5}, NewFilterCislovackyRange(0, 2, P, 5), true},
		{Kombinacia{1, 2, 3, 4, 5}, NewFilterCislovackyRange(2, 2, P, 5), true},
		{Kombinacia{1, 2, 3, 4, 5}, NewFilterCislovackyRange(2, 3, P, 5), true},
		{Kombinacia{1, 2, 3, 4, 5}, NewFilterCislovackyRange(3, 3, P, 5), false},
	}
	for _, test := range tests {
		ok := test.f.Check(test.k)
		if ok != test.w {
			t.Errorf("Excepted: (%v), Got: (%v)", test.w, ok)
		}
	}
}

func TestFilterCislovackyExact(t *testing.T) {
	tests := []struct {
		k Kombinacia
		// f func() (Filter, error)

		ints []int
		c    Cislovacka
		n    int

		w bool
	}{
		{Kombinacia{2}, []int{0, 2}, P, 4, true},
		{Kombinacia{2, 4}, []int{0, 2}, P, 4, true},
		{Kombinacia{2, 4, 6, 7}, []int{0, 2}, P, 4, false},
		{Kombinacia{2, 4, 6, 7}, []int{1, 3}, P, 4, true},
		{Kombinacia{2, 4, 7, 9}, []int{1, 3}, P, 4, false},

		{Kombinacia{1, 3, 7, 9}, []int{1, 3}, P, 4, false},
	}
	for _, test := range tests {
		f, err := NewFilterCislovackyExact(test.ints, test.c, test.n)
		if err != nil {
			t.Fatal(err)
		}
		ok := f.Check(test.k)
		if ok != test.w {
			t.Error(f)
			t.Error(test)
			t.Errorf("Excepted: (%v), Got: (%v)", test.w, ok)
		}
	}
}
