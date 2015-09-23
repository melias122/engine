package archiv

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/melias122/psl/num"
	"github.com/melias122/psl/rw"
)

func (a *Archiv) statistikaZhoda() error {

	// statistika
	stat := struct {
		celkom map[int]int
		zh     map[int]map[string]int
	}{
		celkom: make(map[int]int),
		zh:     make(map[int]map[string]int),
	}
	for i := 0; i <= a.n; i++ {
		stat.zh[i] = make(map[string]int)
	}
	var zh0 *zhodaRiadok
	for _, r := range a.riadky {
		zh1 := makeZhodaRiadok(r.K, zh0)

		stat.celkom[zh1.zh]++
		pZH := []string{}
		for i, c := range zh1.presun {
			if c > 0 {
				s := strconv.Itoa(c) + "/" + strconv.Itoa(i+1)
				if len(s) > 0 {
					pZH = append(pZH, s)
				}
			}
		}
		stat.zh[zh1.zh][strings.Join(pZH, ", ")]++
		zh0 = zh1
	}
	//

	header := []string{"Zhoda", "Pocetnost teor.", "Teoreticka moznost v %", "Pocetnost", "Realne dosiahnute %"}

	w := rw.NewCsvMaxWriter(a.WorkingDir, "StatistikaZhoda", [][]string{header})
	defer w.Close()

	dbLen := float64(len(a.riadky))
	for i := a.n; i >= 0; i-- {
		var (
			c, m big.Int
			r    big.Rat
		)
		c.Mul(c.Binomial(int64(a.n), int64(i)), m.Binomial(int64(a.m-a.n), int64(a.m+i-(2*a.n))))
		r.SetFrac(&c, m.Binomial(int64(a.m), int64(a.n)))
		f, _ := r.Float64()
		if err := w.Write([]string{
			itoa(i),
			c.String(),
			ftoa(f * 100),
			itoa(stat.celkom[i]),
			ftoa((float64(stat.celkom[i]) / dbLen) * 100),
		}); err != nil {
			return err
		}
	}

	var s [][]string
	for i := 1; i <= a.n; i++ {
		s = append(s,
			[]string{""},
			[]string{fmt.Sprintf("Zhoda %d", i), "Pocetnost", "Realne %"},
			[]string{fmt.Sprintf("Zhoda %d", i), itoa(stat.celkom[i]), ftoa((float64(stat.celkom[i]) / dbLen) * 100)},
		)
		for k, v := range stat.zh[i] {
			s = append(s, []string{
				k,
				itoa(v),
				ftoa((float64(v) / dbLen) * 100),
			})
		}
	}
	for _, r := range s {
		if err := w.Write(r); err != nil {
			return err
		}
	}
	return nil
}

func (a *Archiv) statistikaNtice() error {
	stat := struct {
		teorMax map[string]*big.Int
		celkom  map[string]int
		sucin   map[string]map[string]int
	}{
		teorMax: make(map[string]*big.Int),
		celkom:  make(map[string]int),
		sucin:   make(map[string]map[string]int),
	}
	var (
		nticeVsetky = nticeStr(a.n)
		counter     = make(map[string]int)
	)
	for _, ntica := range nticeNtice(a.n) {
		var (
			k        = a.m - a.n + 1
			pocetMax = big.NewInt(1)
			b        big.Int
		)
		for _, n := range ntica {
			if n == 0 {
				continue
			}
			pocetMax.Mul(pocetMax, b.Binomial(int64(k), int64(k-int(n))))
			k -= int(n)
		}
		stat.teorMax[ntica.String()] = pocetMax
	}
	for _, tica := range nticeVsetky {
		counter[tica] = 0
	}
	for _, r := range a.riadky {
		ntica := r.Ntica.String()
		stat.celkom[ntica]++
		sucin := sucinNtic(r.K, r.Ntica)
		if _, ok := stat.sucin[ntica]; !ok {
			stat.sucin[ntica] = make(map[string]int)
		}
		stat.sucin[ntica][sucin]++
	}

	var (
		dbLen = float64(len(a.riadky))
		s     [][]string
	)
	for _, ntica := range nticeVsetky {
		var r big.Rat
		r.SetFrac(stat.teorMax[ntica], big.NewInt(0).Binomial(int64(a.m), int64(a.n)))
		teorPercento, _ := r.Float64()
		s = append(s, []string{
			ntica,
			stat.teorMax[ntica].String(),                    // teor max pocet
			ftoa(teorPercento * 100),                        // teor percento
			itoa(stat.celkom[ntica]),                        // skutocny pocet za DB
			ftoa(float64(stat.celkom[ntica]) / dbLen * 100), // skutocne percento za DB
		})
	}
	s = append(s,
		[]string{""},
		[]string{
			"N-tica", "Sucin pozicie a stlpca", "Pocet vyskytov", "%",
		})
	for _, ntica := range nticeVsetky {
		s = append(s, []string{
			ntica,
			"vsetky",
			itoa(stat.celkom[ntica]),
			ftoa(float64(stat.celkom[ntica]) / dbLen * 100),
		},
		)
		for k, v := range stat.sucin[ntica] {
			s = append(s, []string{
				ntica,
				k,
				itoa(v),
				ftoa(float64(v) / dbLen * 100),
			})
		}
		s = append(s, []string{""})
	}
	header := []string{
		"N-tica", "Pocetnost teor.", "Teoreticka moznost v %",
		"Realne dosiahnuta pocetnost", "Realne dosiahnute %",
	}
	w := rw.NewCsvMaxWriter(a.WorkingDir, "StatistikaNtice", [][]string{header})
	defer w.Close()

	for _, r := range s {
		if err := w.Write(r); err != nil {
			return err
		}
	}
	return nil
}

func (a *Archiv) statistikaCislovacky() error {

	f := func(r Riadok) []byte {
		b := r.C[:]
		b = append(b, byte(r.Zh))
		return b
	}

	header := []string{
		"", "", "", "", "", "", "", "",
		"P", "N", "PR", "Mc", "Vc", "c1-c9", "C0", "cC", "Cc", "CC", "ZH",
	}
	w := rw.NewCsvMaxWriter(a.WorkingDir, "StatistikaCislovacky", [][]string{header})
	defer w.Close()

	tmax := func() []byte {
		var c num.C
		for i := 1; i <= a.m; i++ {
			c.Plus(num.NewC(i))
		}
		sc := c[:]
		sc = append(sc, byte(a.n-1))
		for i, n := range sc {
			if int(n) > a.n {
				sc[i] = byte(a.n)
			}
		}
		return sc
	}
	stat1do := makeStatCifrovacky(a.n, a.m, a.riadky, f, tmax)
	stat1doStrings := statCifrovackyStrings(a.n, a.m, len(a.origHeader), stat1do, tmax)
	for _, s := range stat1doStrings {
		if err := w.Write(s); err != nil {
			return err
		}
	}

	// statistika od-do
	w.Write([]string{})
	w.Write([]string{})
	w.Write([]string{"R OD-DO"})
	if a.Uc.Cislo != 0 && a.Uc.Riadok > 1 {
		statoddo := makeStatCifrovacky(a.n, a.m, a.riadky[a.Uc.Riadok-1:], f, tmax)
		statoddoStrings := statCifrovackyStrings(a.n, a.m, len(a.origHeader), statoddo, tmax)
		for _, s := range statoddoStrings {
			if err := w.Write(s); err != nil {
				return err
			}
		}
	} else {
		w.Write([]string{"Nenastala udalost 101..."})
	}

	return nil
}

// func statArchiv(n, m int, kombinacie [][]byte) {
// csv := [][]string{}
//     QStringList line;
//     QVector<Cislovacky> _cislovacky;

//     Kombinacia pred{};
//     for(auto &akt : kombinacie){
//         _cislovacky.append(cislovacky(akt, pred));
//         pred = akt;
//     }

//     //    Počet možných
//     line << "" << "Pocet moznych";
//     QVector<uint> pm(11,0);
//     for(uint cislo{1}; cislo <= m; ++cislo){
//         if(P(cislo)) ++pm[0];
//         if(N(cislo)) ++pm[1];
//         if(PR(cislo)) ++pm[2];
//         if(Mc(cislo)) ++pm[3];
//         if(Vc(cislo)) ++pm[4];
//         if(C19(cislo)) ++pm[5];
//         if(C0(cislo)) ++pm[6];
//         if(cC(cislo)) ++pm[7];
//         if(Cc(cislo)) ++pm[8];
//         if(CC(cislo)) ++pm[9];
//     }
//     pm[10] = n - 1;

//     for(auto el : pm){
//         line << uintToQString(el);
//     }
//     csv.append(line);

//     //    teoret možné
//     line.clear();
//     line << "" << "Teor mozne";
//     for(auto el : pm){
//         QString tm("0-");
//         if(el >= n)
//             tm += uintToQString(n);
//         else
//             tm += uintToQString(el);
//         line << tm;
//     }
//     csv.append(line);

//     QVector<QVector<double>> ap(11,{}), hp(11,{}), vp(11,{});

//     for(unsigned i{0}; i <= n; ++i){
//         csv.push_back({});
//         line.clear();

//         QVector<int> pu(11,0), krok(11,0), vyskyt(11,0), sums(11,0);
//         QVector<QVector<int>> krok_pocet(11,{});             // krok, pocet pre N,P,PR...

//         int riadok = 1;
//         for (auto &qv : _cislovacky) {
//             for(int poz{0}; poz < 11; ++poz){
//                 if(qv[poz] == i){
//                     ++pu[poz];                          // pocet udalost i
//                     vyskyt[poz] = riadok;
//                     if(krok[poz] > 0){
//                         krok_pocet[poz].push_back(krok[poz]);
//                         sums[poz] += krok[poz];
//                     }
//                     krok[poz] = 0;
//                 }
//                 else{
//                     krok[poz]++;
//                 }
//             }
//             ++riadok;
//         }

//         line << "" << "Pocet udalost " + uintToQString(i) ;
//         for(auto c : pu){
//             line << uintToQString(c);
//         }
//         csv.append(line);

//         //    Súčet diferencií krok udalosť i
//         line.clear();
//         line << "" << "Sucet diferencii krok udalost " + uintToQString(i);
//         for(int j{0}; j< 11; ++j){
//             line << uintToQString(sums[j]);
//         }
//         csv.append(line);

//         //    Krok aritmetický priemer krok udalosť  i
//         line.clear();
//         line << "" << "Krok aritmeticky priemer udalost " + uintToQString(i);
//         QVector<double> aritmeticky_priemer(11, 0.f);
//         for(int j{0}; j< 11; ++j){
//             if(krok_pocet[j].size() > 0)
//                 aritmeticky_priemer[j] = static_cast<double>(sums[j])/krok_pocet[j].size();
//             line << uintToQString(round(aritmeticky_priemer[j]));
//         }

//         csv.append(line);

//         //    Krok harmonický priemer krok udalosť i
//         line.clear();
//         line << "" << "Krok harmonicky priemer udalost " + uintToQString(i);
//         QVector<double> harmonicky_priemer(11, 0.f);
//         for(int j{0}; j< 11; ++j){
//             double hp = 0.f;

//             for(auto &c : krok_pocet[j]){
//                 hp += 1.0f/c;
//             }
//             if(hp > 0)
//                 harmonicky_priemer[j] = krok_pocet[j].size()/hp;
//             line << uintToQString(round(harmonicky_priemer[j]));
//         }

//         csv.append(line);

//         //    Krok vážený priemer krok udalosť i
//         QSet<int> set;
//         line.clear();
//         line << "" << "Krok vazeny priemer udalost " + uintToQString(i);
//         QVector<double> vazeny_priemer(11, 0.f);
//         for(int j{0}; j< 11; ++j){
//             double sum = 0.f;

//             for(auto el : krok_pocet[j])
//                 set.insert(el);
//             for(auto el : set)
//                 sum += el;

//             if(sum > 0)
//                 vazeny_priemer[j] = sums[j]/sum;
//             line << uintToQString(round(vazeny_priemer[j]));
//         }

//         csv.append(line);

//         //    Riadok posedný výskyt udalosť i
//         line.clear();
//         line << "" << "Riadok posledny vyskyt udalost " + uintToQString(i);
//         for(int j{0}; j< 11; ++j){
//             line << uintToQString(vyskyt[j]);
//         }
//         csv.append(line);

//         //    Riadokposedný výskyt i + Krok aritmetický priemer
//         line.clear();
//         line << "" << "Krok aritmeticky priemer + riadok posledny vyskyt " + uintToQString(i);
//         for(int j{0}; j< 11; ++j){
//             line << uintToQString(round(vyskyt[j] + aritmeticky_priemer[j]));
//             ap[j].push_back(round(vyskyt[j] + aritmeticky_priemer[j]));
//         }

//         csv.append(line);

//         //    Riadok posedný výskyt i + Krok harmonický priemer
//         line.clear();
//         line << "" << "Krok harmonicky priemer + riadok posledny vyskyt " + uintToQString(i);
//         for(int j{0}; j< 11; ++j){
//             line << uintToQString(round(vyskyt[j] + harmonicky_priemer[j]));
//             hp[j].push_back(round(vyskyt[j] + harmonicky_priemer[j]));
//         }

//         csv.append(line);

//         //    Riadok posedný výskyt i + Krok vážený priemer
//         line.clear();
//         line << "" << "Krok vazeny priemer + riadok posledny vyskyt " + uintToQString(i);
//         for(int j{0}; j< 11; ++j){
//             line << uintToQString(round(vyskyt[j] + vazeny_priemer[j]));
//             vp[j].push_back(vyskyt[j] + vazeny_priemer[j]);
//         }

//         csv.append(line);
//     }

//     auto najblizsia = [](QVector<double> &vec, unsigned dlzka){
//         double naj = dlzka + 1;
//         int i = 0, ret = -1;

//         for(auto c : vec){
//             auto val = abs(dlzka - c);
//             if(val < naj){
//                 naj = val;
//                 ret = i;
//             }
//             ++i;
//         }
//         return ret;
//     };

// //    csv.append({});
//     csv.push_back({});
//     line.clear();
//     line << "" << "Riadok posedny vyskyt + Krok aritmeticky priemer";
//     for(int i{0}; i < 11; ++i){
//         line << QString::number(najblizsia(ap[i], _cislovacky.size()));
//     }
//     csv.append(line);

//     line.clear();
//     line << "" << "Riadok posedny vyskyt + Krok harmonicky priemer";
//     for(int i{0}; i < 11; ++i){
//         line << QString::number(najblizsia(hp[i], _cislovacky.size()));
//     }
//     csv.append(line);

//     line.clear();
//     line << "" << "Riadok posedny vyskyt + Krok vazeny priemer";
//     for(int i{0}; i < 11; ++i){
//         line << QString::number(najblizsia(vp[i], _cislovacky.size()));
//     }
//     csv.append(line);

//     line.clear();
//     line << "" << "Riadok posedný výskyt - Kriterium v  r+1 (riadku)";
//     //    for(int i{0}; i < 11; ++i){
//     //        auto naj = najblizsia(najblizsia(hp[i], _cislovacky.size()));
//     //        if()
//     //    }
//     csv.append(line);

//     exportCsv(csv, pwd() + "StatistikaCislovacky_" + suborName() + ".csv");
// }
