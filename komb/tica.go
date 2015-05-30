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
	Tica
	n []int
	t int
}

func (n *Ntica) push(x int) {
	n.n = append(n.n, x)
	len := len(n.n)
	if len > 1 {
		t := false
		if x-n.n[len-2] == 1 {
			t = true
		}
		if t {
			n.Tica[n.t]--
			n.t++
		} else {
			n.t = 0
		}
		n.Tica[n.t]++
	} else {
		n.Tica[n.t]++
	}
}

func (n *Ntica) pop() {
}
