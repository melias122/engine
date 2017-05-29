package hrx

import (
	"math"
	"strconv"
	"testing"

	"github.com/melias122/engine/engine"
)

var (
	Hrx  = NewHrx(30, 90)
	HHrx = NewHHrx(30, 90)
)

func BenchmarkHrxValue(b *testing.B) {
	n, m := 30, 90
	hrx := NewHHrx(n, m)
	for i := 0; i < 1000; i++ {
		hrx.Add((i%m)+1, i%n)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hrx.Value(nil)
	}
	b.ReportAllocs()
}

func BenchmarkValueKombinacia(b *testing.B) {
	n, m := 30, 90
	k := make(engine.Kombinacia, n)
	for i := 1; i <= n; i++ {
		k[i-1] = i
	}
	hrx := NewHHrx(n, m)
	for i := 0; i < 1000; i++ {
		hrx.Add((i%m)+1, i%n)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hrx.Value(k)
	}
	b.ReportAllocs()
}

func BenchmarkPow16(b *testing.B) {
	x := 1.1234567890
	y := 16.0
	for i := 0; i < b.N; i++ {
		math.Pow(x, y)
	}
}

func BenchmarkPow16Multi(b *testing.B) {
	x := 1.1234567890
	// y := 16.0
	for i := 0; i < b.N; i++ {
		x *= x
		x *= x
		x *= x
		x *= x
	}
}

func TestNewHrx(t *testing.T) {
	const n, m = 5, 35
	hrx := NewHrx(n, m)
	if hrx.n != n {
		t.Errorf("Excepted: (5), Got: (%d)", hrx.n)
	}
	if hrx.m != m {
		t.Errorf("Excepted: (5), Got: (%d)", hrx.m)
	}
	if len(hrx.Cisla) != m {
		t.Errorf("Excepted: (%d), Got: (%d)", m, len(hrx.Cisla))
	}
}

func TestHrxValue(t *testing.T) {
	const n, m = 5, 35
	hrx := NewHHrx(n, m)

	if strconv.FormatFloat(hrx.Value(nil), 'f', 2, 64) != "100.00" {
		t.Errorf("Excepted: (%s), Got: (%.2f)", "100.00", hrx.Value(nil))
	}

	tests := []struct {
		k     []int
		value string
	}{
		{[]int{2, 7, 13, 32, 35}, "96.219545819576"},
		{[]int{1, 14, 15, 17, 19}, "91.932271522492"},
		{[]int{4, 9, 10, 25, 27}, "86.944174388998"},
		{[]int{1, 2, 13, 21, 31}, "84.684202839784"},
		{[]int{17, 21, 29, 32, 34}, "82.226978235840"},
	}
	for _, test := range tests {
		for y, x := range test.k {
			hrx.Add(x, y)
		}
		if strconv.FormatFloat(hrx.Value(nil), 'f', 12, 64) != test.value {
			t.Errorf("Excepted: (%s), Got: (%.12f)", test.value, hrx.Value(nil))
		}
	}
}

func TestHrxValueKombinacia(t *testing.T) {
	const n, m = 5, 35
	hrx := NewHHrx(n, m)

	for _, a := range [][]int{
		{2, 7, 13, 32, 35},
		{1, 14, 15, 17, 19},
		{4, 9, 10, 25, 27},
		{1, 2, 13, 21, 31},
	} {
		for y, x := range a {
			hrx.Add(x, y)
		}
	}

	tests := []struct {
		k     engine.Kombinacia
		value string
	}{
		{engine.Kombinacia{2}, "84.709009486369"},
		{engine.Kombinacia{2, 7}, "84.707220404804"},
		{engine.Kombinacia{2, 7, 13}, "84.707220377503"},
		{engine.Kombinacia{2, 7, 13, 32}, "84.705431182569"},
		{engine.Kombinacia{17, 21, 29, 32, 34}, "82.226978235840"},
	}
	for _, test := range tests {
		vk := hrx.Value(test.k)
		if strconv.FormatFloat(vk, 'f', 12, 64) != test.value {
			t.Errorf("Excepted: (%s), Got: (%.12f)", test.value, vk)
		}
	}
}
