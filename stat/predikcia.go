// +build ignore

package stat

import (
	"errors"
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/melias122/engine/csv"

	"github.com/montanaflynn/stats"
)

type predictionData struct {
	last int
	step []float64
}

type prediction struct {
	r1         int
	data       map[string]predictionData
	candidates candidates
	predicted  []string
}

func newStepPrediction(r1 int, initData []string) prediction {
	p := prediction{
		r1:   r1,
		data: make(map[string]predictionData),
	}
	for _, s := range initData {
		p.new(0, s)
	}
	return p
}

func (p *prediction) new(last int, s string) {
	d := p.data[s]
	d.step = append(d.step, 1)
	d.last = last
	p.data[s] = d
}

func (p *prediction) addString(last int, s string) {
	p.new(last, s)
	for k, v := range p.data {
		if k == s {
			continue
		}
		step := v.step
		step[len(step)-1]++
	}
}

func score(r1, predict int) int {
	if predict >= r1 {
		return predict - r1
	}
	return r1 - predict
}

type candidate struct {
	score int
	s     string
}

type candidates []candidate

func (b candidates) Len() int           { return len(b) }
func (b candidates) Less(i, j int) bool { return b[i].score < b[j].score }
func (b candidates) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b candidates) Sort()              { sort.Sort(b) }

func (p *prediction) predict() error {
	if len(p.data) == 0 {
		return errors.New("Z nicoho nedokazem predikovat nticu")
	}
	for k, v := range p.data {
		// udalost nenastala
		if v.last == 0 {
			continue
		}
		mean := int(harmonickyPriemer(v.step))
		last := v.last
		score := score(p.r1, mean+last)
		p.candidates = append(p.candidates, candidate{
			score: score,
			s:     k,
		})
	}
	p.candidates.Sort()

	// chcem najviac 3 najlepsie
	max := p.candidates.Len()
	if max > 3 {
		max = 3
	}

	p.predicted = make([]string, max)
	for i := 0; i < max; i++ {
		p.predicted[i] = p.candidates[i].s
	}
	return nil
}

func predictNtice(n, from int, r []Riadok) []string {
	// nova predikcia ntice 1-DO
	nticeStrings := nticeStr(n)
	r1 := len(r)
	p := newStepPrediction(r1, nticeStrings)
	i := from
	for _, r := range r[from:] {
		p.addString(i+1, r.Ntica.String())
	}
	if err := p.predict(); err != nil {
		log.Println(err)
		return []string{}
	}
	ntice := make([]string, len(p.predicted))
	copy(ntice, p.predicted)
	return ntice
}

func predictCislovacka(n, m, from int, r []Riadok, c Cislovacka) []string {
	cislovackyMax := CislovackyMax(n, m)
	r1 := len(r)

	var init []string
	for i := 0; i <= int(cislovackyMax[c]); i++ {
		init = append(init, itoa(i))
	}
	p := newStepPrediction(r1, init)
	for i, r := range r[from:] {
		p.addString(i+1, itoa(int(r.C[c])))
	}
	if err := p.predict(); err != nil {
		log.Println(err)
		return []string{}
	}
	s := make([]string, len(p.predicted))
	copy(s, p.predicted)
	return s
}

func predictZhoda(n, from int, r []Riadok) []string {
	r1 := len(r)

	var init []string
	for i := 0; i <= n; i++ {
		init = append(init, itoa(i))
	}
	p := newStepPrediction(r1, init)
	for i, r := range r[from:] {
		p.addString(i+1, itoa(r.Zh))
	}
	if err := p.predict(); err != nil {
		log.Println(err)
		return []string{}
	}
	s := make([]string, len(p.predicted))
	copy(s, p.predicted)
	return s
}

func predictCifrovacky(n, m, from int, r []Riadok) [10][]string {
	var cifrovacky [10][]string
	cifrovackyMax := CifrovackyTeorMax(n, m)
	r1 := len(r)

	for i, cmax := range cifrovackyMax {
		var init []string
		for j := 0; j <= int(cmax); j++ {
			init = append(init, itoa(j))
		}
		p := newStepPrediction(r1, init)
		for j, r := range r[from:] {
			c := int(r.Cifrovacky[i])
			p.addString(j+1, itoa(c))
		}
		if err := p.predict(); err != nil {
			log.Println(err)
			return [10][]string{}
		}
		cifrovacky[i] = make([]string, len(p.predicted))
		copy(cifrovacky[i], p.predicted)
	}
	return cifrovacky
}

type getFloat64 func(r Riadok) float64

func prepareFloats64Set(riadky []Riadok, getValue getFloat64) []float64 {
	set := make([]float64, len(riadky))
	for i, r := range riadky {
		set[i] = getValue(r)
	}
	return set
}

func predict3Floats64(set []float64) []float64 {
	var s stats.Series
	x := 1.0
	for _, y := range set {
		if y <= 0 {
			y = 1e-10
		}
		s = append(s, stats.Coordinate{X: x, Y: y})
		x++
	}
	lin, _ := stats.LinearRegression(append(s, stats.Coordinate{X: x, Y: 0}))
	exp, _ := stats.ExponentialRegression(s)
	log, _ := stats.LogarithmicRegression(s)

	return []float64{lin[len(lin)-1].Y, exp[len(exp)-1].Y, log[len(log)-1].Y}
}

func (a *Archiv) predikcia() error {

	a.Predict1DO.cislovacky[0] = predictCislovacka(a.n, a.m, 0, a.riadky, P)
	a.PredictODDO.cislovacky[0] = predictCislovacka(a.n, a.m, a.Uc.Riadok, a.riadky, P)

	for _, c := range a.Predict1DO.cislovacky[0] {
		ci, _ := strconv.Atoi(c)
		a.Predict1DO.cislovacky[1] = append(a.Predict1DO.cislovacky[1], itoa(a.n-ci))
	}
	for _, c := range a.PredictODDO.cislovacky[0] {
		ci, _ := strconv.Atoi(c)
		a.PredictODDO.cislovacky[1] = append(a.PredictODDO.cislovacky[1], itoa(a.n-ci))
	}

	a.Predict1DO.cislovacky[2] = predictCislovacka(a.n, a.m, 0, a.riadky, Pr)
	a.PredictODDO.cislovacky[2] = predictCislovacka(a.n, a.m, a.Uc.Riadok, a.riadky, Pr)

	a.Predict1DO.cislovacky[3] = predictCislovacka(a.n, a.m, 0, a.riadky, Mc)
	a.PredictODDO.cislovacky[3] = predictCislovacka(a.n, a.m, a.Uc.Riadok, a.riadky, Mc)

	for _, c := range a.Predict1DO.cislovacky[3] {
		ci, _ := strconv.Atoi(c)
		a.Predict1DO.cislovacky[4] = append(a.Predict1DO.cislovacky[4], itoa(a.n-ci))
	}
	for _, c := range a.PredictODDO.cislovacky[3] {
		ci, _ := strconv.Atoi(c)
		a.PredictODDO.cislovacky[4] = append(a.PredictODDO.cislovacky[4], itoa(a.n-ci))
	}

	a.Predict1DO.cislovacky[5] = predictCislovacka(a.n, a.m, 0, a.riadky, C19)
	a.PredictODDO.cislovacky[5] = predictCislovacka(a.n, a.m, a.Uc.Riadok, a.riadky, C19)

	a.Predict1DO.cislovacky[6] = predictCislovacka(a.n, a.m, 0, a.riadky, C0)
	a.PredictODDO.cislovacky[6] = predictCislovacka(a.n, a.m, a.Uc.Riadok, a.riadky, C0)

	a.Predict1DO.cislovacky[7] = predictCislovacka(a.n, a.m, 0, a.riadky, XcC)
	a.PredictODDO.cislovacky[7] = predictCislovacka(a.n, a.m, a.Uc.Riadok, a.riadky, XcC)

	a.Predict1DO.cislovacky[8] = predictCislovacka(a.n, a.m, 0, a.riadky, Cc)
	a.PredictODDO.cislovacky[8] = predictCislovacka(a.n, a.m, a.Uc.Riadok, a.riadky, Cc)

	a.Predict1DO.cislovacky[9] = predictCislovacka(a.n, a.m, 0, a.riadky, CC)
	a.PredictODDO.cislovacky[9] = predictCislovacka(a.n, a.m, a.Uc.Riadok, a.riadky, CC)

	a.Predict1DO.zhoda = predictZhoda(a.n, 0, a.riadky)
	a.PredictODDO.zhoda = predictZhoda(a.n, a.Uc.Riadok, a.riadky)

	a.Predict1DO.ntice = predictNtice(a.n, 0, a.riadky)
	a.PredictODDO.ntice = predictNtice(a.n, a.Uc.Riadok, a.riadky)

	a.Predict1DO.cifrovacky = predictCifrovacky(a.n, a.m, 0, a.riadky)
	a.PredictODDO.cifrovacky = predictCifrovacky(a.n, a.m, a.Uc.Riadok, a.riadky)

	getVal := func(r Riadok) float64 { return r.Sm }
	a.Predict1DO.sm = predict3Floats64(prepareFloats64Set(a.riadky, getVal))
	a.PredictODDO.sm = predict3Floats64(prepareFloats64Set(a.riadky[a.Uc.Riadok:], getVal))

	getVal = func(r Riadok) float64 { return r.Kk }
	a.Predict1DO.kk = predict3Floats64(prepareFloats64Set(a.riadky, getVal))
	a.PredictODDO.kk = predict3Floats64(prepareFloats64Set(a.riadky[a.Uc.Riadok:], getVal))

	getVal = func(r Riadok) float64 { return r.R1 }
	a.Predict1DO.r1 = predict3Floats64(prepareFloats64Set(a.riadky, getVal))
	a.PredictODDO.r1 = predict3Floats64(prepareFloats64Set(a.riadky[a.Uc.Riadok:], getVal))

	getVal = func(r Riadok) float64 { return r.S1 }
	a.Predict1DO.stl1 = predict3Floats64(prepareFloats64Set(a.riadky, getVal))
	a.PredictODDO.stl1 = predict3Floats64(prepareFloats64Set(a.riadky[a.Uc.Riadok:], getVal))

	getVal = func(r Riadok) float64 { return r.HHrx }
	a.Predict1DO.hhrx = predict3Floats64(prepareFloats64Set(a.riadky, getVal))
	a.PredictODDO.hhrx = predict3Floats64(prepareFloats64Set(a.riadky[a.Uc.Riadok:], getVal))

	getVal = func(r Riadok) float64 { return r.R2 }
	a.Predict1DO.r2 = predict3Floats64(prepareFloats64Set(a.riadky, getVal))
	a.PredictODDO.r2 = predict3Floats64(prepareFloats64Set(a.riadky[a.Uc.Riadok:], getVal))

	getVal = func(r Riadok) float64 { return r.S2 }
	a.Predict1DO.stl2 = predict3Floats64(prepareFloats64Set(a.riadky, getVal))
	a.PredictODDO.stl2 = predict3Floats64(prepareFloats64Set(a.riadky[a.Uc.Riadok:], getVal))

	getVal = func(r Riadok) float64 { return r.Hrx }
	a.Predict1DO.hrx = predict3Floats64(prepareFloats64Set(a.riadky, getVal))
	a.PredictODDO.hrx = predict3Floats64(prepareFloats64Set(a.riadky[a.Uc.Riadok:], getVal))

	getVal = func(r Riadok) float64 { return float64(r.Sucet) }
	a.Predict1DO.sucet = predict3Floats64(prepareFloats64Set(a.riadky, getVal))
	a.PredictODDO.sucet = predict3Floats64(prepareFloats64Set(a.riadky[a.Uc.Riadok:], getVal))

	// fmt.Println(a.Predict1DO.String())

	return nil
}

func normLoop(sk Skupiny, values []float64, f func(s Skupina) []float64) {
	for i, x := range values {
		// pravdepobna hodnota
		p := [2]float64{x, math.MaxFloat64}
		for _, s := range sk {
			for _, y := range f(s) {
				dt := x - y
				if dt < 0 {
					dt = -dt
				}
				if dt < p[1] {
					p[0] = y
					p[1] = dt
				}
			}
		}
		values[i] = p[0]
	}
}

func normalizePrediction(p *Prediction, s Skupiny) {

	normLoop(s, p.hrx, func(s Skupina) []float64 { return []float64{s.Hrx} })
	normLoop(s, p.hhrx, func(s Skupina) []float64 { return s.HHrx[:] })
	normLoop(s, p.r1, func(s Skupina) []float64 { return s.R1[:] })
	normLoop(s, p.r2, func(s Skupina) []float64 { return []float64{s.R2} })
	normLoop(s, p.stl1, func(s Skupina) []float64 { return s.S1[:] })
	normLoop(s, p.stl2, func(s Skupina) []float64 { return s.S2[:] })
}

func savePredictions(workingDir string, p1DO, pODDO Prediction) error {
	w := csv.NewCsvMaxWriter("Predikcia", workingDir)
	defer w.Close()

	s1 := p1DO.Record("Predikcia 1-DO")
	for _, rec := range s1 {
		if err := w.Write(rec); err != nil {
			return err
		}
	}
	w.Write([]string{})
	w.Write([]string{})

	s2 := pODDO.Record("Predikcia OD-DO")
	for _, rec := range s2 {
		if err := w.Write(rec); err != nil {
			return err
		}
	}
	return nil
}

type Prediction struct {
	cislovacky     [10][]string
	zhoda          []string
	sm, kk         []float64
	ntice          []string
	r1, stl1, hhrx []float64
	r2, stl2, hrx  []float64
	sucet          []float64
	cifrovacky     [10][]string
}

func (p *Prediction) Record(name string) [][]string {
	var s [][]string
	s = append(s,
		[]string{name},
		[]string{""},
		[]string{"r+1 na zaklade kroku", "Vitaz 1", "Vitaz 2", "Vitaz 3"},
	)
	// cislovacky
	for i := 0; i < 10; i++ {
		var c []string
		c = append(c, fmt.Sprint(Cislovacka(i)))
		for _, val := range p.cislovacky[i] {
			c = append(c, val)
		}
		s = append(s, c)
	}

	// zhoda
	zh := []string{"Zh"}
	for _, val := range p.zhoda {
		zh = append(zh, val)
	}
	s = append(s, zh)

	// cifrovacky
	for i := 0; i < 10; i++ {
		var c []string
		c = append(c, fmt.Sprint("Cifra ", (i+1)%10))
		for _, val := range p.cifrovacky[i] {
			c = append(c, val)
		}
		s = append(s, c)
	}
	// ntice
	ntice := []string{"Ntice"}
	for _, val := range p.ntice {
		ntice = append(ntice, val)
	}
	s = append(s, ntice)
	s = append(s, []string{})
	s = append(s, []string{"r+1 predikcia", "Linearna", "Exponencialna", "Logaritmicka"})

	//r1
	r1 := []string{"R 1-DO"}
	for _, val := range p.r1 {
		r1 = append(r1, ftoa(val))
	}
	s = append(s, r1)

	//s1
	s1 := []string{"STL 1-DO"}
	for _, val := range p.stl1 {
		s1 = append(s1, ftoa(val))
	}
	s = append(s, s1)

	//HHrx
	hhrx := []string{"HHrx"}
	for _, val := range p.hhrx {
		hhrx = append(hhrx, ftoa(val))
	}
	s = append(s, hhrx)

	//r2
	r2 := []string{"R OD-DO"}
	for _, val := range p.r2 {
		r2 = append(r2, ftoa(val))
	}
	s = append(s, r2)

	//s2
	s2 := []string{"STL OD-DO"}
	for _, val := range p.stl2 {
		s2 = append(s2, ftoa(val))
	}
	s = append(s, s2)

	//Hrx
	hrx := []string{"Hrx"}
	for _, val := range p.hrx {
		hrx = append(hrx, ftoa(val))
	}
	s = append(s, hrx)

	//Sucet
	sucet := []string{"Sucet"}
	for _, val := range p.sucet {
		sucet = append(sucet, itoa(int(val)))
	}
	s = append(s, sucet)

	// sm
	sm := []string{"Smernica"}
	for _, val := range p.sm {
		sm = append(sm, ftoa(val))
	}
	s = append(s, sm)

	// kk
	kk := []string{"Korelacia"}
	for _, val := range p.kk {
		kk = append(kk, ftoa(val))
	}
	s = append(s, kk)

	return s
}

func (p *Prediction) String() string {
	var s []string
	for i := 0; i < 10; i++ {
		s = append(s, fmt.Sprint(Cislovacka(i), p.cislovacky[i]))
	}
	s = append(s, fmt.Sprint("Zh: ", p.zhoda))
	s = append(s, fmt.Sprint("Sm: ", p.sm))
	s = append(s, fmt.Sprint("Kk: ", p.kk))
	s = append(s, fmt.Sprint("Ntice: ", p.ntice))
	s = append(s, fmt.Sprint("R1: ", p.r1))
	s = append(s, fmt.Sprint("S1: ", p.stl1))
	s = append(s, fmt.Sprint("HHrx: ", p.hhrx))
	s = append(s, fmt.Sprint("R2: ", p.r2))
	s = append(s, fmt.Sprint("S2: ", p.stl2))
	s = append(s, fmt.Sprint("Hrx: ", p.hrx))
	s = append(s, fmt.Sprint("Sucet: ", p.sucet))
	for i := 0; i < 10; i++ {
		s = append(s, fmt.Sprint("Cifra ", (i+1)%10, p.cifrovacky[i]))
	}
	return strings.Join(s, "\n")
}
