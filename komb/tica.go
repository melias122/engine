package komb

import (
	"bytes"
	"strconv"
)

type Tica []byte

func (t Tica) String() string {
	var buf bytes.Buffer
	for i, el := range t {
		if i > 0 {
			buf.WriteString(" ")
		}
		buf.WriteString(strconv.Itoa(int(el)))
	}
	return buf.String()
}

type Ntica struct {
	t Tica
	n []int
	p int
}

func newNtica(n int) Ntica {
	return Ntica{
		t: make(Tica, n),
		n: make([]int, 0, n),
	}
}

func (n *Ntica) push(x int) {
	n.n = append(n.n, x)
	len := len(n.n)
	if len > 1 {
		if x-n.n[len-2] == 1 {
			n.t[n.p]--
			n.p++
		} else {
			n.p = 0
		}
	}
	n.t[n.p]++
}

func (n *Ntica) pop() {
	len := len(n.n)
	if len > 1 {
		x := n.n[len-1]
		n.n = n.n[:len-1]
		if x-n.n[len-2] == 1 {
			// V pripade ze posledna cislica je volna, tak n.p == 0
			// To znamena ze potrebujeme zistit predchadzajucu nticu
			if n.p == 0 {
				for i := 1; i < len && n.n[len-i]-n.n[len-i-1] == 1; i++ {
					n.p++
				}
			}
			n.t[n.p]--
			n.p--
			n.t[n.p]++
		} else {
			n.t[0]--
		}

	}
}
