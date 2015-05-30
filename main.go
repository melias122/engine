package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"

	"github.com/melias122/psl/archiv"
)

func main() {

	cpuprofile := flag.String("cpuprofile", "", "Write cpu profile to file")

	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	n, m := 5, 35
	f, err := os.Open(fmt.Sprintf("testdata/%d%d.csv", n, m))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	a := archiv.New(n, m)
	if err = a.Parse(f); err != nil {
		panic(err)
	}
	// if err = a.PocetnostR(); err != nil {
	// 	panic(err)
	// }
	// if err = a.PocetnostS(); err != nil {
	// 	panic(err)
	// }

	// test()
}

type x struct {
	sk int
	c  []int
}

var (
	sk = []x{
		{1, []int{26, 52, 61, 68, 69}},
		{2, []int{7, 18, 21, 23, 27, 38, 46, 48, 55, 60, 63, 65, 70, 76}},
		{3, []int{2, 15, 20, 25, 29, 32, 34, 41, 42, 43, 49, 50, 57, 59, 77, 79}},
		{4, []int{4, 5, 6, 10, 11, 24, 30, 31, 33, 36, 40, 44, 47, 56, 62, 64, 66, 72, 75, 78, 80}},
		{5, []int{17, 22, 45, 58, 67, 71, 74}},
		{6, []int{3, 8, 19, 39, 51, 54}},
		{7, []int{9, 13, 14, 16, 28, 53, 73}},
		{8, []int{1, 35}},
		{9, []int{37}},
		{10, []int{12}},
	}
)

type w struct {
	s, w int
}

var blb = [][]w{}
var y = make([]w, 0, 30)
var spolu int

func ss(sk []x, n int) {
	if len(sk) == 0 {
		return
	}
	var xx int
	if n > len(sk[0].c) {
		xx = len(sk[0].c)
	} else {
		xx = n
	}
	for i := xx; i > 0; i-- {
		y = append(y, w{sk[0].sk, i})
		if n-i > 0 {
			ss(sk[1:], n-i)
		} else {
			// fmt.Println(y)
			cp := make([]w, len(y))
			copy(cp, y)
			blb = append(blb, cp)
			spolu++
		}
		y = y[:len(y)-1]
	}
	ss(sk[1:], n)
}

func test() {
	ss(sk, 30)
	fmt.Println(spolu)
	fmt.Println(len(blb))
	time.Sleep(10 * time.Second)
}

func perm(n int, emit func([]int)) {
	var rc func([]int, int, int)
	rc = func(a []int, i, sum int) {
		a[i] = sum
		emit(a[:i+1])
		for a[i]--; a[i] > 0; a[i]-- {
			rc(a, i+1, sum-a[i])
		}
	}
	rc(make([]int, n), 0, n)
}

func comb(n, m int, emit func([]int)) {
	s := make([]int, n)
	last := n - 1

	var rc func(int, int)

	rc = func(i, next int) {
		for j := next; j < m; j++ {
			s[i] = j
			if i == last {
				emit(s)
			} else {
				rc(i+1, j+1)
			}
		}
	}
	rc(0, 0)
}

func GenNtice(n int) [][]byte {

	var ntice [][]byte

	nt := make([]byte, n)
	nt_end := make([]byte, n)

	nt[0] = byte(n)
	nt_end[n-1] = byte(1)

	idx := 0
	for !bytes.Equal(nt, nt_end) {

		sum := 0
		for i, e := range nt {
			sum += int(e) * (i + 1)
		}

		if sum == n {

			s := make([]byte, len(nt))
			copy(s, nt)
			ntice = append(ntice, s)

			nt[idx]--
			idx++
		} else if sum < n {
			nt[idx]++
		} else {
			nt[idx]--
			idx++
		}

		if idx == len(nt) {
			idx--
			for nt[idx] == 0 {
				idx--
			}
			nt[idx]--
			idx++
		}
	}

	return ntice
}

func ForwardLinearPrediction(coefs, x []float64) {
	// GET SIZE FROM INPUT VECTORS
	N := len(x) - 1
	m := len(coefs)

	// INITIALIZE R WITH AUTOCORRELATION COEFFICIENTS
	R := make([]float64, m+1)
	for i := 0; i <= m; i++ {
		for j := 0; j <= N-i; j++ {
			R[i] += x[j] * x[j+i]
		}
	}

	// INITIALIZE Ak
	Ak := make([]float64, m+1)
	Ak[0] = 1.0

	// INITIALIZE Ek
	Ek := R[0]

	// LEVINSON-DURBIN RECURSION
	for i := 0; i < m; i++ {
		// COMPUTE LAMBDA
		lambda := 0.0
		for j := 0; j <= i; j++ {
			lambda -= Ak[j] * R[i+1-j] //7
		}
		lambda /= Ek

		// UPDATE Ak
		for k := 0; k <= (i+1)/2; k++ {
			temp := Ak[i+1-k] + lambda*Ak[k]
			Ak[k] = Ak[k] + lambda*Ak[i+1-k]
			Ak[i+1-k] = temp
		}

		// UPDATE Ek
		Ek *= 1.0 - lambda*lambda
	}

	// TODO: assisgn...
	// ASSIGN COEFFICIENTS
	// coeffs.assign( ++Ak.begin(), Ak.end() );
}
