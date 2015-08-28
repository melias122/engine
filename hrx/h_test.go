package hrx

import (
	"math"
	"strconv"
	"testing"
)

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
	hrx := New(n, m)
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

// func TestAdd(t *testing.T) {
// 	const n, m = 2, 4
// 	hrx := New(n, m)
//
// 	tests := []struct {
// 		x, y int
// 		w    []int8
// 	}{
// 		{1, 0, []int8{3, 1, 0, 0}},
// 		{1, 0, []int8{3, 0, 1, 0}},
// 		{2, 0, []int8{2, 1, 1, 0}},
// 		{3, 0, []int8{1, 2, 1, 0}},
// 		{4, 0, []int8{0, 3, 1, 0}},
// 		{4, 0, []int8{0, 2, 2, 0}},
// 		{1, 0, []int8{0, 2, 1, 1}},
// 	}
// 	for _, test := range tests {
// 		hrx.Add(test.x, test.y)
// 		if !reflect.DeepEqual(hrx.skupiny, test.w) {
// 			t.Errorf("Excepted: (%v), Got: (%v)", test.w, hrx.skupiny)
// 		}
// 	}
// }

func TestIs101(t *testing.T) {
	const n, m = 2, 4
	hrx := New(n, m)

	tests := []struct {
		x, y  int
		is101 bool
	}{
		{1, 0, false},
		{2, 0, false},
		{3, 0, false},
		{4, 0, true},
		{1, 0, true},
	}
	for _, test := range tests {
		hrx.Add(test.x, test.y)
		if hrx.Is101() != test.is101 {
			t.Errorf("Excepted: (%v), Got: (%v)", test.is101, hrx.Is101())
		}
	}
}

func TestValue(t *testing.T) {
	const n, m = 5, 35
	hrx := New(n, m)

	if strconv.FormatFloat(hrx.Value(), 'f', 2, 64) != "100.00" {
		t.Errorf("Excepted: (%s), Got: (%.2f)", "100.00", hrx.Value())
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
		if strconv.FormatFloat(hrx.Value(), 'f', 12, 64) != test.value {
			t.Errorf("Excepted: (%s), Got: (%.12f)", test.value, hrx.Value())
		}
	}
}

func TestValueKombinacia(t *testing.T) {
	const n, m = 5, 35
	hrx := New(n, m)

	for _, a := range [][]int{
		[]int{2, 7, 13, 32, 35},
		[]int{1, 14, 15, 17, 19},
		[]int{4, 9, 10, 25, 27},
		[]int{1, 2, 13, 21, 31},
	} {
		for y, x := range a {
			hrx.Add(x, y)
		}
	}

	tests := []struct {
		k     []byte
		value string
	}{
		{[]byte{2}, "84.709009486369"},
		{[]byte{2, 7}, "84.707220404804"},
		{[]byte{2, 7, 13}, "84.707220377503"},
		{[]byte{2, 7, 13, 32}, "84.705431182569"},
		{[]byte{17, 21, 29, 32, 34}, "82.226978235840"},
	}
	for _, test := range tests {
		vk := hrx.ValueKombinacia(test.k)
		if strconv.FormatFloat(vk, 'f', 12, 64) != test.value {
			t.Errorf("Excepted: (%s), Got: (%.12f)", test.value, vk)
		}
	}
}

func TestHrx(t *testing.T) {
	k_hrx := [][]int{
		[]int{1, 3, 5},
		[]int{2, 4, 6},
	}
	// k_hhrx := [][]int{
	// 	[]int{1, 5, 6},
	// 	[]int{1, 3, 5},
	// 	[]int{2, 4, 6},
	// }
	hrx := New(3, 6)
	for _, i := range k_hrx {
		for y, x := range i {
			hrx.Add(x, y)
		}
	}
	// t.Log(hrx)
}
