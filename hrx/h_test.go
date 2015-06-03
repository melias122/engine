package hrx

import "testing"

func TestGenerujPresun(t *testing.T) {
	// sk := []Tab{
	// 	{1, []int{8}},
	// 	{2, []int{6, 13, 16, 18, 31, 35}},
	// 	{4, []int{2, 14, 33, 34}},
	// 	{3, []int{1, 9, 12, 15, 17, 21, 22, 24, 25, 26}},
	// 	{5, []int{5, 10, 19, 27, 30}},
	// 	{6, []int{4, 11, 29}},
	// 	{7, []int{23, 28}},
	// 	{8, []int{3, 20, 32}},
	// 	{10, []int{7}},
	// }
	// sk := []Tab{
	// 	{1, 1},
	// 	{2, 6},
	// 	{4, 4},
	// 	{3, 10},
	// 	{5, 5},
	// 	{6, 3},
	// 	{7, 2},
	// 	{8, 3},
	// 	{10, 1},
	// }
	t.Log(h)
	hrx := h.Get()
	spolu := 0
	ch := GenerujPresun(h.Presun(), 5)
	for p := range ch {
		spolu++
		h.Simul(p)
		if hrx != h.Get() {
			t.Log(h)
			t.Fatalf("Excepted: %f, Have: %f", hrx, h.Get())
		}

	}
	if spolu != 904 {
		t.Fatalf("Excepted: 904, Have: %d", spolu)
	}
}

var (
	//Hrx
	h *H = &H{
		m:   35,
		max: 10,
		sk: map[int]int{
			1:  1,
			2:  6,
			3:  10,
			4:  4,
			5:  5,
			6:  3,
			7:  2,
			8:  3,
			10: 1,
		}}
	//HHrx
	h2 *H = &H{
		m:   35,
		max: 182,
		sk: map[int]int{
			122: 1,
			131: 2,
			132: 2,
			135: 1,
			137: 1,
			138: 3,
			139: 1,
			140: 1,
			141: 1,
			142: 2,
			144: 2,
			147: 1,
			148: 2,
			149: 1,
			150: 1,
			152: 1,
			154: 3,
			155: 1,
			157: 1,
			158: 1,
			159: 1,
			166: 1,
			167: 1,
			169: 1,
			170: 1,
			182: 1,
		}}
)

// func TestHrx(t *testing.T) {

// p := Presun{
// 	Tab{10, 1},
// 	Tab{8, 3},
// 	Tab{7, 1},
// }
// t.Log(h.Simul(p))
// }
