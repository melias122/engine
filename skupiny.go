package engine

type Skupiny []Skupina

type Skupina struct {
	R1   [2]float64
	S1   [2]float64
	HHrx [2]float64

	R2  float64
	S2  [2]float64
	Hrx float64

	Sucet [2]uint16

	Cislovacky [2]Cislovacky
	Zh         [2]byte

	// Ntica Ntica
	// Xtica

	Cifrovacky Cifrovacky

	Xcisla Xcisla

	// STLNtica
}
