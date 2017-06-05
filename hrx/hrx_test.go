package hrx

import (
	"math"
	"strconv"
	"testing"

	"github.com/melias122/engine/engine"
)

// BenchmarkHrx-4             	10000000	       130 ns/op	       0 B/op	       0 allocs/op
func BenchmarkHrx(b *testing.B) {
	n, m := 30, 90
	hrx := NewHrx(n, m)
	for i := 0; i < b.N; i++ {
		hrx.X(nil)
	}
	b.ReportAllocs()
}

// BenchmarkHrxKombinacia-4   	10000000	       201 ns/op	       0 B/op	       0 allocs/op
func BenchmarkHrxKombinacia(b *testing.B) {
	n, m := 30, 90
	k := make(engine.Kombinacia, n)
	for i := 1; i <= n; i++ {
		k[i-1] = i
	}
	hrx := NewHrx(n, m)
	for i := 0; i < b.N; i++ {
		hrx.X(k)
	}
	b.ReportAllocs()
}

// BenchmarkMathPow16-4       	20000000	        57.5 ns/op
func BenchmarkMathPow16(b *testing.B) {
	x := 1.1234567890
	y := 16.0
	for i := 0; i < b.N; i++ {
		math.Pow(x, y)
	}
}

// BenchmarkHrxPow16-4        	2000000000	         0.72 ns/op
func BenchmarkHrxPow16(b *testing.B) {
	x := 1.1234567890
	// y := 16.0
	for i := 0; i < b.N; i++ {
		x *= x
		x *= x
		x *= x
		x *= x
	}
}

func float64Equal(f1, f2 float64, prec int) bool {
	v1 := strconv.FormatFloat(f1, 'g', prec, 64)
	v2 := strconv.FormatFloat(f2, 'g', prec, 64)
	return v1 == v2
}

func TestHHrx(t *testing.T) {
	const n, m = 5, 35
	hrx := NewHHrx(n, m)

	if strconv.FormatFloat(hrx.X(nil), 'f', 2, 64) != "100.00" {
		t.Errorf("Excepted: (%s), Got: (%.2f)", "100.00", hrx.X(nil))
	}

	tests := []struct {
		k     engine.Kombinacia
		value string
	}{
		{engine.Kombinacia{2, 7, 13, 32, 35}, "96.219545819576"},
		{engine.Kombinacia{1, 14, 15, 17, 19}, "91.932271522492"},
		{engine.Kombinacia{4, 9, 10, 25, 27}, "86.944174388998"},
		{engine.Kombinacia{1, 2, 13, 21, 31}, "84.684202839784"},
		{engine.Kombinacia{17, 21, 29, 32, 34}, "82.226978235840"},
	}
	for _, test := range tests {
		hrx.Add([]engine.Kombinacia{test.k})
		if strconv.FormatFloat(hrx.X(nil), 'f', 12, 64) != test.value {
			t.Errorf("Excepted: (%s), Got: (%.12f)", test.value, hrx.X(nil))
		}
	}
}

func TestHrx535(t *testing.T) {
	hrx := NewHrx(5, 35)
	for i := range k_hrx535.k {
		hrx.Add(k_hrx535.k[0 : i+1])
		value := hrx.X(nil)
		if !float64Equal(value, k_hrx535.hrx[i], 14) {
			t.Fatalf("expected %v, got %v", k_hrx535.hrx[i], value)
		}
	}
}

func TestHHrx535(t *testing.T) {
	hhrx := NewHHrx(5, 35)
	for _, v := range k_hhrx535 {
		hhrx.Add([]engine.Kombinacia{v.k})
		value := hhrx.X(nil)
		if !float64Equal(value, v.hhrx, 14) {
			t.Fatalf("expected %v, got %v", v.hhrx, value)
		}
	}
}

var k_hhrx535 = []struct {
	k    engine.Kombinacia
	hhrx float64
}{
	{engine.Kombinacia{2, 7, 13, 32, 35}, 96.21954581957614},
	{engine.Kombinacia{1, 14, 15, 17, 19}, 91.93227152249185},
	{engine.Kombinacia{4, 9, 10, 25, 27}, 86.94417438899828},
	{engine.Kombinacia{1, 2, 13, 21, 31}, 84.68420283978378},
	{engine.Kombinacia{17, 21, 29, 32, 34}, 82.22697823584026},
	{engine.Kombinacia{14, 16, 23, 30, 34}, 78.0675609713104},
	{engine.Kombinacia{1, 3, 6, 16, 26}, 73.15552713524555},
	{engine.Kombinacia{3, 6, 7, 23, 24}, 71.24954895332954},
	{engine.Kombinacia{11, 16, 23, 30, 31}, 69.18359917532598},
	{engine.Kombinacia{8, 12, 20, 22, 25}, 58.22588452912038},
	{engine.Kombinacia{6, 18, 29, 32, 35}, 54.20407914137429},
	{engine.Kombinacia{4, 15, 17, 20, 23}, 54.59963667061779},
	{engine.Kombinacia{7, 9, 13, 16, 24}, 54.5114576770521},
	{engine.Kombinacia{1, 26, 27, 28, 30}, 49.37657934833153},
	{engine.Kombinacia{8, 10, 23, 33, 34}, 43.04709745060901},
	{engine.Kombinacia{2, 5, 6, 18, 27}, 27.54726745149525},
	{engine.Kombinacia{17, 25, 27, 31, 32}, 27.52795338997487},
	{engine.Kombinacia{24, 27, 29, 31, 33}, 26.49814181994931},
	{engine.Kombinacia{2, 17, 19, 30, 32}, 25.35667664034957},
	{engine.Kombinacia{15, 17, 19, 28, 34}, 28.663291373048384},
	{engine.Kombinacia{5, 14, 23, 27, 29}, 26.859483837626623},
	{engine.Kombinacia{9, 14, 17, 31, 32}, 30.69454314118526},
	{engine.Kombinacia{8, 10, 28, 29, 33}, 30.243820818306865},
	{engine.Kombinacia{7, 13, 20, 29, 33}, 30.117909801024457},
	{engine.Kombinacia{8, 12, 32, 33, 35}, 27.59989615668915},
	{engine.Kombinacia{6, 10, 18, 20, 32}, 30.446920517310343},
	{engine.Kombinacia{7, 11, 12, 26, 35}, 26.523886291639087},
	{engine.Kombinacia{6, 8, 14, 26, 29}, 26.502531195078248},
	{engine.Kombinacia{10, 14, 20, 21, 28}, 26.10867553155068},
	{engine.Kombinacia{1, 4, 6, 25, 32}, 28.113562079606364},
	{engine.Kombinacia{3, 5, 15, 19, 24}, 26.83490932695292},
	{engine.Kombinacia{3, 18, 19, 25, 30}, 26.71882252974132},
	{engine.Kombinacia{5, 8, 25, 31, 32}, 28.479146351600136},
	{engine.Kombinacia{3, 5, 15, 23, 31}, 28.45433149027293},
	{engine.Kombinacia{3, 12, 16, 21, 29}, 28.254930099700417},
	{engine.Kombinacia{1, 2, 3, 29, 33}, 28.245522713229228},
	{engine.Kombinacia{1, 2, 18, 32, 33}, 29.8147856564002},
	{engine.Kombinacia{3, 7, 12, 22, 35}, 22.99619673662761},
	{engine.Kombinacia{12, 16, 18, 29, 32}, 25.09262291664387},
	{engine.Kombinacia{13, 16, 19, 31, 35}, 25.016007811242385},
	{engine.Kombinacia{9, 12, 15, 32, 34}, 26.404426695475674},
	{engine.Kombinacia{10, 12, 20, 25, 35}, 26.371713355347953},
	{engine.Kombinacia{16, 25, 27, 33, 35}, 26.369308352504056},
	{engine.Kombinacia{4, 8, 9, 14, 33}, 25.77712668838702},
	{engine.Kombinacia{7, 16, 27, 28, 29}, 25.67595797275386},
	{engine.Kombinacia{1, 3, 10, 17, 33}, 25.673680948632864},
	{engine.Kombinacia{6, 8, 15, 18, 30}, 25.653714221394917},
	{engine.Kombinacia{9, 19, 21, 24, 32}, 26.952589203640958},
	{engine.Kombinacia{9, 11, 16, 30, 33}, 24.232241481397654},
}

var k_hrx535 = struct {
	k   []engine.Kombinacia
	hrx []float64
}{
	k: []engine.Kombinacia{
		engine.Kombinacia{2, 7, 13, 32, 35},
		engine.Kombinacia{1, 14, 15, 17, 19},
		engine.Kombinacia{4, 9, 10, 25, 27},
		engine.Kombinacia{1, 2, 13, 21, 31},
		engine.Kombinacia{17, 21, 29, 32, 34},
		engine.Kombinacia{14, 16, 23, 30, 34},
		engine.Kombinacia{1, 3, 6, 16, 26},
		engine.Kombinacia{3, 6, 7, 23, 24},
		engine.Kombinacia{11, 16, 23, 30, 31},
		engine.Kombinacia{8, 12, 20, 22, 25},
		engine.Kombinacia{6, 18, 29, 32, 35},
		engine.Kombinacia{4, 15, 17, 20, 23},
		engine.Kombinacia{7, 9, 13, 16, 24},
		engine.Kombinacia{1, 26, 27, 28, 30},
		engine.Kombinacia{8, 10, 23, 33, 34},
		engine.Kombinacia{2, 5, 6, 18, 27},
		engine.Kombinacia{17, 25, 27, 31, 32},
		engine.Kombinacia{24, 27, 29, 31, 33},
		engine.Kombinacia{2, 17, 19, 30, 32},
		engine.Kombinacia{15, 17, 19, 28, 34},
		engine.Kombinacia{5, 14, 23, 27, 29},
		engine.Kombinacia{9, 14, 17, 31, 32},
		engine.Kombinacia{8, 10, 28, 29, 33},
		engine.Kombinacia{7, 13, 20, 29, 33},
		engine.Kombinacia{8, 12, 32, 33, 35},
		engine.Kombinacia{6, 10, 18, 20, 32},
		engine.Kombinacia{7, 11, 12, 26, 35},
		engine.Kombinacia{6, 8, 14, 26, 29},
		engine.Kombinacia{10, 14, 20, 21, 28},
		engine.Kombinacia{1, 4, 6, 25, 32},
		engine.Kombinacia{3, 5, 15, 19, 24},
		engine.Kombinacia{3, 18, 19, 25, 30},
		engine.Kombinacia{5, 8, 25, 31, 32},
		engine.Kombinacia{3, 5, 15, 23, 31},
		engine.Kombinacia{3, 12, 16, 21, 29},
		engine.Kombinacia{1, 2, 3, 29, 33},
		engine.Kombinacia{1, 2, 18, 32, 33},
		engine.Kombinacia{3, 7, 12, 22, 35},
		engine.Kombinacia{12, 16, 18, 29, 32},
		engine.Kombinacia{13, 16, 19, 31, 35},
		engine.Kombinacia{9, 12, 15, 32, 34},
		engine.Kombinacia{10, 12, 20, 25, 35},
		engine.Kombinacia{16, 25, 27, 33, 35},
		engine.Kombinacia{4, 8, 9, 14, 33},
		engine.Kombinacia{7, 16, 27, 28, 29},
		engine.Kombinacia{1, 3, 10, 17, 33},
		engine.Kombinacia{6, 8, 15, 18, 30},
		engine.Kombinacia{9, 19, 21, 24, 32},
		engine.Kombinacia{9, 11, 16, 30, 33},
	},
	hrx: []float64{
		100,
		100,
		100,
		100,
		100,
		100,
		100,
		100,
		100,
		100,
		100,
		100,
		100,
		100,
		100,
		66.64648353945898,
		65.6681589910004,
		64.32007471644017,
		69.10373995717715,
		67.27162271064574,
		65.90170696584472,
		64.83093510529247,
		63.24888844830738,
		61.762534598563946,
		59.96909384173584,
		59.06255471833764,
		56.83241672305258,
		56.42137450222582,
		59.36562634357559,
		57.139471846705824,
		58.75135107197244,
		56.980020480415696,
		56.57688216670812,
		55.69856545069767,
		53.087249660430444,
		51.95108549834381,
		51.43726712799387,
		64.14849898973152,
		63.286918010644186,
		61.88681261420529,
		62.30507133401868,
		61.8861368233652,
		62.682664545925114,
		61.44235557177621,
		60.12203030690176,
		63.055800722426234,
		61.63952400350403,
		60.0272437696334,
		60.74448222026025,
	},
}
