package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"github.com/melias122/psl/archiv"
	"github.com/melias122/psl/filter"
	"github.com/melias122/psl/generator"
	"github.com/melias122/psl/komb"
)

var a *archiv.Archiv

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

	n, m := 5, 35
	path := fmt.Sprintf("testdata/535.csv")

	workingDir, err := os.Getwd()
	a, err = archiv.Make(path, workingDir, n, m)
	if err != nil {
		panic(err)
	}

	filters := filter.Filters{
		filter.NewSucet(n, 100, 150),
		filter.NewStlNtica(n, komb.Tica{2, 0, 1, 0, 0}, []byte{0, 0, 1, 1, 1}),
		filter.NewXtica(n, m, komb.Tica{0, 0, 2, 3}),
		// filter.NewZhoda(n, 1, 1, a.K),
		// filter.NewZakazane(m, []byte{a.Uc.Cislo}),
		// filter.NewHrx(n, 36.23, 36.23, a.Hrx, "HRX"),
		// filter.NewR(n, 3.464E-014, 3.464E-014, a.HHrx.Cisla, "ƩR 1-DO"),
		// filter.NewCislovacky(n, 7, 9, num.IsN, "N"),
		// filter.NewCislovacky(n, 3, 5, num.IsPr, "Pr"),
	}
	msg := make(chan string)
	go func() {
		generator.GenerateKombinacie(n, a, filters, msg)
	}()
	fmt.Println(<-msg)
}

// func ForwardLinearPrediction(coefs, x []float64) {
// 	// GET SIZE FROM INPUT VECTORS
// 	N := len(x) - 1
// 	m := len(coefs)
//
// 	// INITIALIZE R WITH AUTOCORRELATION COEFFICIENTS
// 	R := make([]float64, m+1)
// 	for i := 0; i <= m; i++ {
// 		for j := 0; j <= N-i; j++ {
// 			R[i] += x[j] * x[j+i]
// 		}
// 	}
//
// 	// INITIALIZE Ak
// 	Ak := make([]float64, m+1)
// 	Ak[0] = 1.0
//
// 	// INITIALIZE Ek
// 	Ek := R[0]
//
// 	// LEVINSON-DURBIN RECURSION
// 	for i := 0; i < m; i++ {
// 		// COMPUTE LAMBDA
// 		lambda := 0.0
// 		for j := 0; j <= i; j++ {
// 			lambda -= Ak[j] * R[i+1-j] //7
// 		}
// 		lambda /= Ek
//
// 		// UPDATE Ak
// 		for k := 0; k <= (i+1)/2; k++ {
// 			temp := Ak[i+1-k] + lambda*Ak[k]
// 			Ak[k] = Ak[k] + lambda*Ak[i+1-k]
// 			Ak[i+1-k] = temp
// 		}
//
// 		// UPDATE Ek
// 		Ek *= 1.0 - lambda*lambda
// 	}
//
// 	// TODO: assisgn...
// 	// ASSIGN COEFFICIENTS
// 	// coeffs.assign( ++Ak.begin(), Ak.end() );
// }
