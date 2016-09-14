package engine

var archivRiadokHeader = []string{
	"Poradové číslo",
	"Kombinacie",
	"P", "N", "Sled PN",
	"Pr", "Sled PNPr",
	"Mc", "Vc", "Sled McVc",
	"C19", "C0", "cC", "Cc", "CC", "Sled prirodzené kritéria",
	"ZH", "SPZH",
	"Sm", "ΔSm",
	"Kk", "ΔKk",
	"Ntica", "Ntica súčet", "Ntica súčin pozície a stĺpca",
	"Xtica",
	"ƩR 1-DO", "ΔƩR 1-DO", "ƩR 1-DO \"r+1\"",
	"ƩSTL1-DO", "ΔƩSTL 1-DO", "ƩSTL 1-DO \"r+1\"",
	"Δ(ƩR 1-DO - ƩSTL 1-DO)",
	"HHRX", "ΔHHRX",
	"ƩR OD-DO", "ΔƩR OD-DO",
	"ƩSTL OD-DO", "ΔƩSTL OD-DO",
	"Δ(ƩR OD-DO - ƩSTL OD-DO)",
	"HRX", "ΔHRX",
	"ƩKombinacie", "ΔƩKombinacie",
	"UC číslo", "UC riadok",
	"Cifra 1", "Cifra 2", "Cifra 3", "Cifra 4", "Cifra 5",
	"Cifra 6", "Cifra 7", "Cifra 8", "Cifra 9", "Cifra 0",
}

type Uc struct {
	Cislo  int
	Riadok int
}

type Riadok struct {
	n, m        int
	origStrings []string

	Pc             int
	K              Kombinacia
	C              Cislovacky
	Zh             int
	ZhPresun       string
	Sm, DtSm       float64
	Kk, DtKk       float64
	Ntica          Tica
	Xtica          Tica
	R1, DtR1       float64
	Rp1            float64
	S1, DtS1       float64
	Sp1            float64
	R1mS1          float64
	HHrx, DtHHrx   float64
	R2, DtR2       float64
	S2, DtS2       float64
	R2mS2          float64
	Hrx, DtHrx     float64
	Sucet, DtSucet int
	Uc
	Cifrovacky Cifrovacky
}

func (r *Riadok) Add(k Kombinacia, n1, n2 Nums, hrx, hhrx float64) {

	Sm := Smernica(k, r.n, r.m)
	R1, S1 := k.SucetRS(n1)
	R2, S2 := k.SucetRS(n2)

	if r.K != nil {
		r.Zh = Zhoda(r.K, k)
		r.ZhPresun = NewZhodaPresun(r.K, k).String()
		Kk := Korelacia(r.K, k, r.n, r.m)
		r.DtKk = Kk - r.Kk
		r.Kk = Kk

		// Dt
		r.DtR1 = R1 - r.R1
		r.DtS1 = S1 - r.S1
		r.DtR2 = R2 - r.R2
		r.DtS2 = S2 - r.S2
		r.DtHrx = hrx - r.Hrx
		r.DtHHrx = hhrx - r.HHrx
		r.DtSm = Sm - r.Sm
		r.DtSucet = k.Sucet() - r.Sucet
	}
	r.Sm = Sm
	r.R1 = R1
	// r.Rp1 = n1.RNext()
	r.S1 = S1
	// r.Sp1 = sumSnext(k, n1)
	r.Rp1, r.Sp1 = k.SucetRSNext(n1)
	r.R1mS1 = R1 - S1
	r.R2 = R2
	r.S2 = S2
	r.R2mS2 = R2 - S2
	r.Hrx = hrx
	r.HHrx = hhrx
	r.Sucet = k.Sucet()

	r.K = k
	r.C = k.Cislovacky()
	r.Ntica = Ntica(k)
	r.Xtica = Xtica(r.m, k)

	r.Cifrovacky = k.Cifrovacky()
}

func (r Riadok) record() []string {
	rec := make([]string, 0, len(archivRiadokHeader))
	rec = append(rec, itoa(int(r.Pc)), r.K.String())
	cislovacky := r.C.Strings()
	rec = append(rec, cislovacky[0:2]...)
	rec = append(rec, r.K.SledPN())
	rec = append(rec, cislovacky[2])
	rec = append(rec, r.K.SledPNPr())
	rec = append(rec, cislovacky[3:5]...)
	rec = append(rec, r.K.SledMcVc())
	rec = append(rec, cislovacky[5:]...)
	rec = append(rec, r.K.SledPrirodzene())
	rec = append(rec,
		itoa(r.Zh),
		r.ZhPresun,
		ftoa(r.Sm),
		ftoa(r.DtSm),
		ftoa(r.Kk),
		ftoa(r.DtKk),
		r.Ntica.String(),
		NticaSucet(r.K).String(),
		NticaSucin(r.K).String(),
		r.Xtica.String(),
		ftoa(r.R1),
		ftoa(r.DtR1),
		ftoa(r.Rp1),
		ftoa(r.S1),
		ftoa(r.DtS1),
		ftoa(r.Sp1),
		ftoa(r.R1mS1),
		ftoa(r.HHrx),
		ftoa(r.DtHHrx),
		ftoa(r.R2),
		ftoa(r.DtR2),
		ftoa(r.S2),
		ftoa(r.DtS2),
		ftoa(r.R2mS2),
		ftoa(r.Hrx),
		ftoa(r.DtHrx),
		itoa(r.Sucet),
		itoa(r.DtSucet),
		itoa(int(r.Cislo)),
		itoa(r.Riadok),
	)
	rec = append(rec, r.Cifrovacky.Strings()...)
	return rec
}
