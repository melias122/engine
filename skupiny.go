package psl

type Skupiny []Skupina

type Skupina struct {
	Hrx    float64
	HHrx   [2]float64
	R1     [2]float64
	R2     float64
	S1     [2]float64
	S2     [2]float64
	Sucet  [2]uint16
	Xcisla Xcisla
}
