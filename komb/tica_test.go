package komb

import "testing"

func TestNticaPush(t *testing.T) {
	tests := []struct {
		t []int
		w string
	}{
		{[]int{1, 3, 5, 7, 9}, "5 0 0 0 0"},
		{[]int{1, 3, 5, 7, 8}, "3 1 0 0 0"},
		{[]int{1, 2, 5, 7, 9}, "3 1 0 0 0"},
		{[]int{1, 3, 4, 7, 9}, "3 1 0 0 0"},
		{[]int{1, 3, 5, 6, 9}, "3 1 0 0 0"},
		{[]int{1, 3, 4, 5, 9}, "2 0 1 0 0"},
		{[]int{1, 2, 3, 5, 7}, "2 0 1 0 0"},
		{[]int{1, 3, 5, 6, 7}, "2 0 1 0 0"},
		{[]int{1, 2, 3, 4, 7}, "1 0 0 1 0"},
		{[]int{1, 3, 4, 5, 6}, "1 0 0 1 0"},
		{[]int{1, 2, 3, 4, 5}, "0 0 0 0 1"},
		{[]int{1, 2, 3, 5, 6}, "0 1 1 0 0"},
		{[]int{1, 2, 4, 5, 6}, "0 1 1 0 0"},
		{[]int{1, 3, 4, 6, 7}, "1 2 0 0 0"},
		{[]int{1, 2, 4, 6, 7}, "1 2 0 0 0"},
		{[]int{1, 2, 4, 5, 7}, "1 2 0 0 0"},
		{[]int{1, 2, 3, 4}, "0 0 0 1"},
		{[]int{1, 3, 5, 7}, "4 0 0 0"},
		{[]int{1, 3, 5}, "3 0 0"},
		{[]int{1, 3}, "2 0"},
		{[]int{1}, "1"},
		{[]int{}, ""},
		{[]int{1, 14, 15, 17, 19}, "3 1 0 0 0"},
		{[]int{4, 9, 10, 25, 27}, "3 1 0 0 0"},
		{[]int{1, 3, 6, 16, 26}, "5 0 0 0 0"},
	}
	for _, test := range tests {
		n := Ntica{n: []int{}, Tica: make(Tica, len(test.t))}
		for _, x := range test.t {
			n.push(x)
		}
		if n.String() != test.w {
			t.Fatalf("test: %s; Excepted: (%s), Have: (%s)", test, test.w, n.String())
		}
	}
}

// func TestNticaPop(t *testing.T) {

// }
