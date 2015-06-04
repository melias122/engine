package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"

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

	n, m := 20, 80
	path := fmt.Sprintf("testdata/%d%d.csv", n, m)
	_, err := archiv.Make(path, n, m)
	if err != nil {
		panic(err)
	}
	// if err = a.PocetnostR(); err != nil {
	// 	panic(err)
	// }
	// if err = a.PocetnostS(); err != nil {
	// 	panic(err)
	// }
	// if err = a.HrxHHrx(); err != nil {
	// 	panic(err)
	// }

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
