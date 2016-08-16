package engine

import "strconv"

type Skupiny []Skupina

type Skupina struct {
	Hrx        float64
	HrxDelta   float64
	Xcisla     Xcisla
	R2         float64
	S2         [2]float64
	Sucet      [2]int
	PocetKomb  string
	HHrx       [2]float64
	HHrxDelta  [2]float64
	R1         [2]float64
	S1         [2]float64
	Cislovacky [2]Cislovacky
	Zh         [2]int
}

func (s *Skupina) Record() []string {
	r := append([]string{},
		Ftoa(s.Hrx),
		Ftoa(s.HrxDelta),
		s.Xcisla.String(),
		Ftoa(s.R2),
		Ftoa(s.S2[0]),
		Ftoa(s.S2[1]),
		strconv.Itoa(s.Sucet[0]),
		strconv.Itoa(s.Sucet[1]),
		s.PocetKomb,
		Ftoa(s.HHrx[0]),
		Ftoa(s.HHrx[1]),
		Ftoa(s.HHrxDelta[0]),
		Ftoa(s.HHrxDelta[1]),
		Ftoa(s.R1[0]),
		Ftoa(s.R1[1]),
		Ftoa(s.S1[0]),
		Ftoa(s.S1[1]),
	)
	r = append(r, s.formatCislovackyMinMax()...)
	r = append(r, strconv.Itoa(s.Zh[0]), strconv.Itoa(s.Zh[1]))
	return r

}
func (s *Skupina) formatCislovackyMinMax() []string {
	r := make([]string, 20)
	for i := 0; i < 10; i++ {
		j := i * 2
		r[j] = strconv.Itoa(int(s.Cislovacky[0][i]))
		r[j+1] = strconv.Itoa(int(s.Cislovacky[1][i]))
	}
	return r
}
