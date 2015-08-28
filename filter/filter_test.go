package filter

import (
	"strconv"
	"testing"

	"github.com/melias122/psl/komb"
)

func TestParseNtica(t *testing.T) {
	tests := []struct {
		s string
		w komb.Tica
		n int
	}{
		{s: "", w: komb.Tica{}, n: 5},
		{s: "             \t\n   \t\t\n", w: komb.Tica{}, n: 5},
		{s: "5", w: komb.Tica{5, 0, 0, 0, 0}, n: 5},
		{s: "5 0 0 0 0", w: komb.Tica{5, 0, 0, 0, 0}, n: 5},
		{s: "5 0 0 0 0 0", w: komb.Tica{}, n: 5},
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
		w    komb.Tica
		e    bool
	}{
		{n: 5, m: 35, s: "", w: komb.Tica{}, e: true},
		{n: 5, m: 35, s: "    ", w: komb.Tica{}, e: true},
		{n: 5, m: 35, s: "    \t\t\t\t\t\t \n\n  \t      ", w: komb.Tica{}, e: true},
		{n: 5, m: 35, s: "1 2", e: true}, // 1+2 != 5
		{n: 5, m: 35, s: "1 2 0 0 1", e: true},
		{n: 5, m: 35, s: "1 2 2 2", e: true},
		{n: 5, m: 35, s: "5,0,0", e: true},
		{n: 5, m: 35, s: "5;", e: true},

		{n: 5, m: 35, s: "5", w: komb.Tica{5, 0, 0, 0}},
		{n: 5, m: 35, s: "5 ", w: komb.Tica{5, 0, 0, 0}},
		{n: 5, m: 35, s: "3 2", w: komb.Tica{3, 2, 0, 0}},
		{n: 5, m: 35, s: "1 2 0 2", w: komb.Tica{1, 2, 0, 2}},
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

func TestGRTLSS(t *testing.T) {
	tests := []struct {
		f float64
	}{
		{f: 10},
		{f: 0},
		{f: 0.1},
		{f: 0.9},
		{f: 0.11},
		{f: 0.19},
		{f: 0.0156977747110574},
		{f: 32.3354},
	}
	for _, test := range tests {
		LSS := nextLSS(test.f)
		GRT := nextGRT(test.f)
		if !(LSS < test.f) {
			t.Errorf("%f expected to be smaller than %f", LSS, test.f)
		}
		if !(GRT > test.f) {
			t.Errorf("%f expected to be greater than %f", GRT, test.f)
		}
	}

	f := 0.0156977747110574
	for _, fi := range []float64{
		0.0156977747110574,
		0.015697774711057,
		0.01569777471105,
		0.0156977747110,
		0.015697774711,
		0.01569777471,
		0.0156977747,
		0.015697774,
		0.01569777,
		0.0156977,
		0.015697,
		0.01569,
		0.0156,
		0.015,
		0.01,
		0.02,
		0.016,
		0.0157,
		0.015698,
		0.0156978,
		0.015697775,
	} {
		LSS := nextLSS(fi)
		GRT := nextGRT(fi)
		if !(LSS < f) || !(GRT > f) {
			t.Errorf("%f, %f, %f", LSS, GRT, f)
		}
	}
}

func TestDt(t *testing.T) {
	tests := []struct {
		f    float64
		w    string
		prec int
	}{
		{f: 32.123, w: "0.001", prec: 3},
		{f: 0.1, w: "0.1", prec: 1},
		{f: 99, w: "1", prec: -1},
		{f: 1, w: "1", prec: -1},
		{f: 0, w: "1", prec: -1},
		{f: 1.4e-5, w: "0.000001", prec: 6},
	}
	for _, test := range tests {
		s := strconv.FormatFloat(dt(test.f), 'f', test.prec, 64)
		if s != test.w {
			t.Errorf("Expected: %s, Got: %s (%v)", test.w, s, test)
		}
	}
}

// func TestParseBytes(t *testing.T) {
// 	tests := []struct {
// 		s string
// 		w string
// 	}{
// 		{s: "", w: ""},
// 		{s: "       \t \t \t   \n", w: ""},
// 	}
// }
