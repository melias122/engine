package psl

import "fmt"

func maker(xcisla Xcisla, n int) {
	var (
		sum     int
		x       Xcisla
		indices = make([]int, 1, xcisla.Len())
	)
	for len(indices) > 0 {
		j := len(indices)

		// i je index daneho xcisla
		i := indices[j-1]

		// na tomto leveli uz nie su dalsie xcisla
		// ideme o level nizsie
		if i == xcisla.Len() {
			indices = indices[:j-1]
			continue
		}

		// skusime xcislo
		t := xcisla[i]

		if x.Len() > 0 {
			last := &x[x.Len()-1]
			if last.Sk == t.Sk {
				last.Max--
				sum--
				if last.Max == 0 {
					x = x[:x.Len()-1]
					indices[j-1]++
					continue
				} else {
					indices = append(indices, i+1)
					continue
				}
			}
		}

		t.Max = min(t.Max, n-sum)
		sum += t.Max
		x = append(x, t)

		// cisel v kombinacii este nie je n
		// skusime dalsie cislo
		if sum < n {
			indices = append(indices, i+1)
			continue
		}
		fmt.Println(x)
	}
}

// func TestItt(t *testing.T) {
// 	x := Xcisla{{1, 1}, {2, 1}, {4, 2}, {5, 1}, {6, 9}, {7, 6}, {8, 8}, {9, 4}, {10, 6}, {11, 3}, {12, 2}, {14, 2}}
// 	maker(x, 7)
// }
