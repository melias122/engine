package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"github.com/melias122/psl/archiv"
	"github.com/melias122/psl/hrx"
)

func main() {

	var (
		cpuprofile = flag.String("cpuprofile", "", "Write cpu profile to file")
	)

	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	// n, m := 5, 35
	n, m := 20, 80
	path := fmt.Sprintf("testdata/%d%d.csv", n, m)
	a, err := archiv.Make(path, n, m)
	if err != nil {
		panic(err)
	}

	htab := hrx.NewHrxTab(a.Hrx, a.HHrx, n, m)
	htab.Make()
}

// func max101perioda(m int) {
// 	var (
// 		max101perioda, pocetPerioda int
// 		pocetRiedka                 int
// 		nKALKULUJ                   int
// 		a101perioda                 []int
// 		aRiedka                     [][]int
// 	)
// 	for kk := 1; kk <= pocetRiedka+nKALKULUJ; kk++ {
// 		for i := 1; i <= m; i++ {
// 			a101perioda[i] = 0
// 		}
// 		var (
// 			kk_od                    = kk
// 			boolSTOP, bool101perioda bool
// 		)
// 		for !boolSTOP {
// 			bool101perioda = true
// 			for i := 1; i <= m; i++ {
// 				a101perioda[i] += aRiedka[kk_od][i]
// 				if a101perioda[i] < 1 {
// 					bool101perioda = false
// 				}
// 			}
// 			if bool101perioda {
// 				boolSTOP = true
// 			}

// 			kk_od--
// 			if kk_od == 0 {
// 				boolSTOP = true
// 			}
// 		}

// 		if bool101perioda || (kk_od == 0) {
// 			if bool101perioda {
// 				pocetPerioda++
// 			}
// 			kk_od++
// 			for i := 1; i <= m; i++ {
// 				if a101perioda[i] > max101perioda {
// 					max101perioda = a101perioda[i]
// 				}
// 			}
// 		}
// 	}
// 	max101perioda++
// 	fmt.Println(max101perioda)
// }

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
