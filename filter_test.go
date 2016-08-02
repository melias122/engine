package engine

import (
	"strconv"
	"testing"
)

// func BenchmarkFilter(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Filters30.Check(K30)
// 	}
// }

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
		{f: 1.58809155486105E-005},
		{f: 9.34969738146142E-006},
		{f: 9.34969738146142E-027},
	}
	for _, test := range tests {
		LSS := nextLSS(test.f)
		GRT := nextGRT(test.f)
		if !(LSS < test.f) {
			t.Fail()
			t.Errorf("%f expected to be smaller than %f", LSS, test.f)
		}
		if !(GRT > test.f) {
			t.Fail()
			t.Errorf("%f expected to be greater than %f", GRT, test.f)
		}
		if !(LSS < test.f && GRT > test.f) {
			t.Fail()
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
