package hrx

import "github.com/melias122/engine/engine"

type Skupina struct {
	Hrx      float64
	HrxDelta float64
	//Xcisla     Xcisla
	R2         float64
	STL2       [2]float64
	Sucet      [2]int
	PocetKomb  string
	HHrx       [2]float64
	HHrxDelta  [2]float64
	R1         [2]float64
	STL1       [2]float64
	Cislovacky [2]engine.Cislovacka
	Zhoda      [2]int
}
