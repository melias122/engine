package hrx

import "testing"

func TestHrx(t *testing.T) {
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
	sk := []Tab{
		{1, 1},
		{2, 6},
		{4, 4},
		{3, 10},
		{5, 5},
		{6, 3},
		{7, 2},
		{8, 3},
		{10, 1},
	}
	spolu := 0
	ch := GenerujPresun(sk, 5)
	for p := range ch {
		spolu++
		// fmt.Println(p)
	}
	if spolu != 904 {
		t.Fatalf("Excepted: 904, Have: %d", spolu)
	}
}
