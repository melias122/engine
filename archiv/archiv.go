package archiv

import (
	"github.com/melias122/engine/engine"
)

type Riadok struct {
	Kombinacia engine.Kombinacia `csv:"K"`
	Sucet      int               `csv:"ƩK"`
	SucetΔ     int               `csv:"ΔƩK"`
	engine.Cislovacka
	Zhoda      int
	Smernica   float64 `csv:"Sm"`
	SmernicaΔ  float64 `csv:"ΔSm"`
	Korelacia  float64 `csv:"Kk"`
	KorelaciaΔ float64 `csv:"ΔKk"`
	Ntica      engine.Tica
	Xtica      engine.Tica
	R1         float64 `csv:"ƩR 1-DO"`
	R1Δ        float64 `csv:"ΔƩR 1-DO"`
	STL1       float64 `csv:"ƩSTL 1-DO"`
	STL1Δ      float64 `csv:"ΔƩSTL 1-DO"`
	HHrx       float64
	HHrxΔ      float64 `csv:"ΔHHrx"`
	R2         float64 `csv:"ƩR OD-DO"`
	R2Δ        float64 `csv:"ΔƩR OD-DO"`
	STL2       float64 `csv:"ΔƩSTL OD-DO"`
	STL2Δ      float64 `csv:"ΔƩSTL OD-DO"`
	Hrx        float64
	HrxΔ       float64 `csv:"ΔHrx"`
	engine.Cifrovacka
}

type Hrx interface {
	engine.Hrx
	Add([]engine.Kombinacia)
}

type Archiv struct {
	n    int
	m    int
	hrx  Hrx
	hhrx Hrx

	Riadok []Riadok
}

func NewArchiv(hhrx Hrx, hrx Hrx, n, m int) *Archiv {
	return &Archiv{
		n:    n,
		m:    m,
		hhrx: hhrx,
		hrx:  hrx,
	}
}

func (a *Archiv) Process(all []engine.Kombinacia) {
	var r0 Riadok
	for i, k := range all {
		a.hhrx.Add(all[i : i+1])
		a.hrx.Add(all[:i+1])
		r1 := Riadok{
			Kombinacia: k,
			Sucet:      k.Sucet(),
			Cislovacka: engine.NewCislovacka(k),
			Zhoda:      engine.Zhoda(r0.Kombinacia, k),
			Smernica:   engine.Smernica(k, a.n, a.m),
			Korelacia:  engine.Korelacia(r0.Kombinacia, k, a.n, a.m),
			Ntica:      engine.NewNtica(k),
			Xtica:      engine.NewXtica(k, a.m),
			R1:         a.hhrx.R(k),
			STL1:       a.hhrx.STL(k),
			HHrx:       a.hhrx.X(nil),
			R2:         a.hrx.R(k),
			STL2:       a.hrx.STL(k),
			Hrx:        a.hrx.X(nil),
			Cifrovacka: engine.NewCifrovacka(k),
		}

		r1.SucetΔ = r1.Sucet - r0.Sucet
		r1.SmernicaΔ = r1.Smernica - r0.Smernica
		r1.KorelaciaΔ = r1.Korelacia - r0.Korelacia
		r1.R1Δ = r1.R1 - r0.R1
		r1.STL1Δ = r1.STL1 - r0.STL1
		r1.HHrxΔ = r1.HHrx - r0.HHrx
		r1.R2Δ = r1.R2 - r0.R2
		r1.STL2Δ = r1.STL2 - r0.STL2
		r1.HrxΔ = r1.Hrx - r0.Hrx

		a.Riadok = append(a.Riadok, r1)
		r0 = r1
	}
}
