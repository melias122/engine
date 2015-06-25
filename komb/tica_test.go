package komb

import "testing"

func TestNtica(t *testing.T) {
	tests := []struct {
		t []byte
		w string
	}{
		{[]byte{1, 3, 5, 7, 9}, "5 0 0 0 0"},
		{[]byte{1, 3, 5, 7, 8}, "3 1 0 0 0"},
		{[]byte{1, 2, 5, 7, 9}, "3 1 0 0 0"},
		{[]byte{1, 3, 4, 7, 9}, "3 1 0 0 0"},
		{[]byte{1, 3, 5, 6, 9}, "3 1 0 0 0"},
		{[]byte{1, 3, 4, 5, 9}, "2 0 1 0 0"},
		{[]byte{1, 2, 3, 5, 7}, "2 0 1 0 0"},
		{[]byte{1, 3, 5, 6, 7}, "2 0 1 0 0"},
		{[]byte{1, 2, 3, 4, 7}, "1 0 0 1 0"},
		{[]byte{1, 3, 4, 5, 6}, "1 0 0 1 0"},
		{[]byte{1, 2, 3, 4, 5}, "0 0 0 0 1"},
		{[]byte{1, 2, 3, 5, 6}, "0 1 1 0 0"},
		{[]byte{1, 2, 4, 5, 6}, "0 1 1 0 0"},
		{[]byte{1, 3, 4, 6, 7}, "1 2 0 0 0"},
		{[]byte{1, 2, 4, 6, 7}, "1 2 0 0 0"},
		{[]byte{1, 2, 4, 5, 7}, "1 2 0 0 0"},
		{[]byte{1, 2, 3, 4}, "0 0 0 1"},
		{[]byte{1, 3, 5, 7}, "4 0 0 0"},
		{[]byte{1, 3, 5}, "3 0 0"},
		{[]byte{1, 3}, "2 0"},
		{[]byte{1}, "1"},
		{[]byte{}, ""},
		{[]byte{1, 14, 15, 17, 19}, "3 1 0 0 0"},
		{[]byte{4, 9, 10, 25, 27}, "3 1 0 0 0"},
		{[]byte{1, 3, 6, 16, 26}, "5 0 0 0 0"},
	}
	for _, test := range tests {
		tica := Ntica(test.t)
		if tica.String() != test.w {
			t.Fatalf("Excepted: (%s), Have: (%s)", test.w, tica)
		}
	}
}

// func TestNticaPush(t *testing.T) {
// 	tests := []struct {
// 		t []byte
// 		w string
// 	}{
// 		{[]byte{1, 3, 5, 7, 9}, "5 0 0 0 0"},
// 		{[]byte{1, 3, 5, 7, 8}, "3 1 0 0 0"},
// 		{[]byte{1, 2, 5, 7, 9}, "3 1 0 0 0"},
// 		{[]byte{1, 3, 4, 7, 9}, "3 1 0 0 0"},
// 		{[]byte{1, 3, 5, 6, 9}, "3 1 0 0 0"},
// 		{[]byte{1, 3, 4, 5, 9}, "2 0 1 0 0"},
// 		{[]byte{1, 2, 3, 5, 7}, "2 0 1 0 0"},
// 		{[]byte{1, 3, 5, 6, 7}, "2 0 1 0 0"},
// 		{[]byte{1, 2, 3, 4, 7}, "1 0 0 1 0"},
// 		{[]byte{1, 3, 4, 5, 6}, "1 0 0 1 0"},
// 		{[]byte{1, 2, 3, 4, 5}, "0 0 0 0 1"},
// 		{[]byte{1, 2, 3, 5, 6}, "0 1 1 0 0"},
// 		{[]byte{1, 2, 4, 5, 6}, "0 1 1 0 0"},
// 		{[]byte{1, 3, 4, 6, 7}, "1 2 0 0 0"},
// 		{[]byte{1, 2, 4, 6, 7}, "1 2 0 0 0"},
// 		{[]byte{1, 2, 4, 5, 7}, "1 2 0 0 0"},
// 		{[]byte{1, 2, 3, 4}, "0 0 0 1"},
// 		{[]byte{1, 3, 5, 7}, "4 0 0 0"},
// 		{[]byte{1, 3, 5}, "3 0 0"},
// 		{[]byte{1, 3}, "2 0"},
// 		{[]byte{1}, "1"},
// 		{[]byte{}, ""},
// 		{[]byte{1, 14, 15, 17, 19}, "3 1 0 0 0"},
// 		{[]byte{4, 9, 10, 25, 27}, "3 1 0 0 0"},
// 		{[]byte{1, 3, 6, 16, 26}, "5 0 0 0 0"},
// 	}
// 	for _, test := range tests {
// 		n := newNtica(len(test.t))
// 		for _, x := range test.t {
// 			n.push(x)
// 		}
// 		if n.t.String() != test.w {
// 			t.Fatalf("Excepted: (%s), Have: (%s)", test.w, n.t.String())
// 		}
// 	}
// }

// func TestNticaPop(t *testing.T) {
// 	tests := []struct {
// 		t []byte
// 		w []string
// 	}{
// 		{[]byte{1, 3, 5}, []string{"3 0 0", "2 0 0", "1 0 0", "0 0 0"}},
// 		{[]byte{1, 2, 3}, []string{"0 0 1", "0 1 0", "1 0 0", "0 0 0"}},
// 		{[]byte{1, 2, 4}, []string{"1 1 0", "0 1 0", "1 0 0", "0 0 0"}},
// 		{[]byte{1, 2, 3, 5, 6}, []string{"1 2 0 0 0", "2 1 0 0 0", "1 1 0 0 0", "0 1 0 0 0", "1 0 0 0 0", "0 0 0 0 0"}},
// 	}
// 	for _, test := range tests {
// 		n := newNtica(len(test.t))
// 		for _, x := range test.t {
// 			n.push(x)
// 		}
// 		t.Log(n)
// 		for i := 0; i < len(test.t); i++ {
// 			if n.t.String() != test.w[i] {
// 				t.Fatalf("Excepted: (%s), Have: (%s)", test.w[i], n.t.String())
// 			}
// 			n.pop()
// 		}
// 	}
// }
