package archiv

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/melias122/psl/komb"
	"github.com/melias122/psl/rw"
)

type zhodaRiadok struct {
	// k je aktualna v riadku
	k      komb.Kombinacia
	zh     int
	krok   int
	krokZh bool
	presun []int // r-1/r
}

func makeZhodaRiadok(k komb.Kombinacia, z0 *zhodaRiadok) *zhodaRiadok {
	if z0 == nil {
		z0 = &zhodaRiadok{
			k:      komb.Kombinacia{},
			presun: make([]int, len(k)),
			krok:   1,
		}
	}
	z := &zhodaRiadok{
		k:      k,
		zh:     komb.Zhoda(k, z0.k),
		presun: make([]int, len(k)),
	}
	if z0.krokZh {
		z.krok = 1
	} else {
		z.krok = z0.krok
	}
	if z.zh > 0 {
		z.krokZh = true
		for i, c1 := range z.k {
			for j, c0 := range z0.k {
				if c0 == c1 {
					z.presun[i] = j + 1
					break
				}
			}
		}
	} else {
		z.krok++
	}
	return z
}

func zhodaRiadokHeader(n int) []string {
	header := []string{"ZH"}
	for i := 0; i < n; i++ {
		header = append(header, strconv.Itoa(i+1))
	}
	header = append(header, "Pozícia zhoda", "Krok")
	for i := 0; i < n; i++ {
		header = append(header, strconv.Itoa(i+1)+" stl r-1/r")
	}
	return header
}

func (z *zhodaRiadok) Strings() []string {
	var s []string
	s = append(s, strconv.Itoa(z.zh))
	for i, c := range z.presun {
		if c > 0 {
			s = append(s, strconv.Itoa(int(z.k[i])))
		} else {
			s = append(s, "")
		}
	}
	pZH := []string{}
	for i, c := range z.presun {
		if c > 0 {
			pZH = append(pZH, strconv.Itoa(c)+"/"+strconv.Itoa(i+1))
		}
	}
	s = append(s, strings.Join(pZH, ", "))
	if z.zh > 0 {
		s = append(s, strconv.Itoa(z.krok))
	} else {
		s = append(s, "")
	}
	for i, c := range z.presun {
		if c > 0 {
			s = append(s, strconv.Itoa(c)+"/"+strconv.Itoa(i+1))
		} else {
			s = append(s, "")
		}
	}
	return s
}

func (a *Archiv) mapaZhoda() error {
	header := make([]string, len(a.origHeader))
	copy(header, a.origHeader)
	header = append(header, zhodaRiadokHeader(a.n)...)

	w := rw.NewCsvMaxWriter(a.WorkingDir, "MapaZhoda", [][]string{header})
	defer w.Close()

	var zh0 *zhodaRiadok
	var mapaZhoda []*zhodaRiadok
	for _, r := range a.riadky {
		mapaZhoda = append(mapaZhoda, makeZhodaRiadok(r.K, zh0))
		zh0 = mapaZhoda[len(mapaZhoda)-1]
	}
	for i, z := range mapaZhoda {
		var record []string
		record = append(record, a.riadky[i].origStrings...)
		record = append(record, z.Strings()...)
		if err := w.Write(record); err != nil {
			return err
		}
	}
	return nil
}

func nticaPozicie(k komb.Kombinacia) []string {
	var s []string
	for i, ok := range komb.NticaPozicie(k) {
		if ok == 1 {
			s = append(s, itoa(int(k[i])))
		} else {
			s = append(s, "")
		}
	}
	return s
}

func sucetNtic(k komb.Kombinacia, ntica komb.Tica) string {
	if len(k) < 2 || int(ntica[0]) == len(k) {
		return ""
	}
	var (
		s     []string
		sucet int
	)
	for i := range k[:len(k)-1] {
		if k[i]+1 == k[i+1] {
			if sucet == 0 {
				sucet = int(k[i])
			}
			sucet += int(k[i+1])
		} else if sucet != 0 {
			s = append(s, itoa(sucet))
			sucet = 0
		}
	}
	if sucet != 0 {
		s = append(s, itoa(sucet))
	}
	return strings.Join(s, ", ")
}

func sucinNtic(k komb.Kombinacia, ntica komb.Tica) string {
	if len(k) < 2 || int(ntica[0]) == len(k) {
		return ""
	}
	var (
		s     []string
		sucin int
	)
	for i := range k[:len(k)-1] {
		if k[i]+1 == k[i+1] {
			if sucin == 0 {
				sucin = int(i + 1)
			}
			sucin *= int(i + 2)
		} else if sucin != 0 {
			s = append(s, itoa(sucin))
			sucin = 0
		}
	}
	if sucin != 0 {
		s = append(s, itoa(sucin))
	}
	return strings.Join(s, ", ")
}

func krokNtica(ntica komb.Tica, nticeVsetky []string, counter map[string]int) []string {
	var s []string
	for _, nticaString := range nticeVsetky {
		if nticaString == ntica.String() {
			s = append(s, itoa(counter[nticaString]))
			counter[nticaString] = 1
		} else {
			if counter[nticaString] > 0 {
				counter[nticaString]++
			}
			s = append(s, "")
		}
	}
	return s
}

func (a *Archiv) mapaNtice() error {
	var (
		s           [][]string
		nticeVsetky = nticeStr(a.n)
		counter     = make(map[string]int)
	)
	for _, tica := range nticeVsetky {
		counter[tica] = 0
	}
	for _, r := range a.riadky {
		var riadok []string
		riadok = append(riadok, r.Ntica.String())
		riadok = append(riadok, nticaPozicie(r.K)...)
		riadok = append(riadok, sucetNtic(r.K, r.Ntica))
		riadok = append(riadok, sucinNtic(r.K, r.Ntica))
		riadok = append(riadok, krokNtica(r.Ntica, nticeVsetky, counter)...)
		s = append(s, riadok)
	}

	header := make([]string, len(a.origHeader))
	copy(header, a.origHeader)
	header = append(header, "N-tica")
	for i := 1; i <= a.n; i++ {
		header = append(header, itoa(i))
	}
	header = append(header, "Sucet N-tic", "Sucin pozicie a stlpca")
	header = append(header, nticeVsetky...)

	w := rw.NewCsvMaxWriter(a.WorkingDir, "MapaNtice", [][]string{header})
	defer w.Close()

	for i, r := range a.riadky {
		var riadok []string
		riadok = append(riadok, r.origStrings...)
		riadok = append(riadok, s[i]...)
		if err := w.Write(riadok); err != nil {
			return err
		}
	}
	return nil
}

func ntice(n int, emit func([]int)) {
	if n < 2 {
		panic("ntice: n < 2")
	}
	tica := make([]int, n)
	end := make([]int, n)
	tica[0] = n
	end[n-1] = 1

	i := 0
	for {
		sum := 0
		for j, e := range tica {
			sum += int(e) * (j + 1)
		}
		if sum == n {
			emit(tica)
			tica[i]--
			i++
		} else if sum < n {
			tica[i]++
		} else {
			tica[i]--
			i++
		}
		if i == n {
			i--
			for tica[i] == 0 {
				i--
			}
			tica[i]--
			i++
		}

		equal := true
		for p := 0; p < n; p++ {
			if tica[p] != end[p] {
				equal = false
			}
		}
		if equal {
			break
		}
	}
	emit(tica)
}

func nticeNtice(n int) []komb.Tica {
	var nticeVsetky []komb.Tica
	ntice(n, func(tica []int) {
		n := make(komb.Tica, len(tica))
		for i := range tica {
			n[i] = byte(tica[i])
		}
		nticeVsetky = append(nticeVsetky, n)
	})
	return nticeVsetky
}

func nticeStr(n int) []string {
	var (
		tice []string
		buf  bytes.Buffer
	)
	ntice(n, func(tica []int) {
		buf.Reset()
		for i, t := range tica {
			if i > 0 {
				buf.WriteString(" ")
			}
			buf.WriteString(itoa(t))
		}
		tice = append(tice, buf.String())
	})
	return tice
}

// type Mapa []struct {
// k     []int
// ntica komb.Tica
// xtica komb.Tica
// }

// func (a *Archiv) MapaNtice(csv [][]string, n int) error {
// func (a *Archiv) MapaNtice() error {

// var (
// mapaNtice   [][]string
// out_stringl []string
// akt                 []int
// csv [][]string
// )

// ntice
// ntice_pos := nticeStr(a.n)
// ntice_krok := make(map[string]int, len(ntice_pos))
// ntice_pocet := make(map[string]int, len(ntice_pos))
// for _, tica := range ntice_pos {
// ntice_pocet[tica] = 0
// ntice_krok[tica] = -1
// }

// header
// mHeader := []string{"N-tica"}
// for i := 1; i <= a.n; i++ {
// mHeader = append(mHeader, itoa(i))
// }
// mHeader = append(mHeader, "Sucet N-tic", "Sucin pozicie a stlpca")
// for _, tica := range ntice_pos {
// mHeader = append(mHeader, fmt.Sprintf("Krok %s", tica))
// }

// out_stringl = append(out_stringl, csv[0]...)
// out_stringl = append(out_stringl, mHeader...)
// mapaNtice = append(mapaNtice, out_stringl)

// analyza
// pos := 0
// for _, akt_r := range csv[1:] {
// for _, akt_r := range a.riadky {
// pos++
// akt := make([]int, n)
// for i := 3; i < n+3; i++ {
// 	c, err := strconv.Atoi(akt_r[i])
// 	if err != nil {
// 		return err
// 	}
// 	akt[i-3] = c
// }
// akt := akt_r.K

// ntica_anl := make([]string, len(ntice_pos)+a.n+3)
// ntica_anl[0] = kombToString(akt)

// if ntica_anl[0] != ntice_pos[0] { // ak neni "n 0 0 0 ...
// 	start, end := -1, -1
// 	for i := 0; i < len(akt)-1; i++ {
// 		for j := i; j < len(akt)-1; j++ {
// 			if (akt[j+1] - akt[j]) == 1 {
// 				if start == -1 {
// 					start = j
// 				}
// 				end = j + 1
// 				i = j
// 			} else {
// 				break
// 			}
// 		}
// 		if start != end {
// 			sum := 0
// 			sum_stl := 1
// 			for j := start; j <= end; j++ {
// 				sum += int(akt[j])
// 				sum_stl *= (j + 1)
// 				// cislo v stl
// 				ntica_anl[j+1] = itoa(int(akt[j]))
// 			}
//
// 			// sucet ntic
// 			if ntica_anl[a.n+1] == "" {
// 				ntica_anl[a.n+1] = itoa(sum)
// 			} else {
// 				s := ntica_anl[a.n+1]
// 				ntica_anl[a.n+1] = fmt.Sprint("%s, %d", s, sum)
// 			}
//
// 			// sucin stlpcov
// 			if ntica_anl[a.n+2] == "" {
// 				ntica_anl[a.n+2] = itoa(sum_stl)
// 			} else {
// 				s := ntica_anl[a.n+2]
// 				ntica_anl[a.n+2] = fmt.Sprint("%s, %d", s, sum_stl)
// 			}
//
// 			start = -1
// 			end = -1
// 		}
// 	}
// }

// ntice stat
// tica := ntica_anl[0]
// sucin_stl := ntica_anl[a.n+2]
// ntice_pocet[tica]++
// if sucin_stl != "" {
// 	if ntice_typ_pocet[tica].value(sucin_stl) > 0 {
// 		p := ntice_typ_pocet[tica].value(sucin_stl)
// 		ntice_typ_pocet[tica].insert(sucin_stl, p+1)
// 	} else {
// 		ntice_typ_pocet[ntica_anl[0]].insert(ntica_anl[a.n+2], 1)
// 	}
// }
//

// for strp := range ntice_krok {
// 	if ntice_krok[strp] != -1 {
// 		ntice_krok[strp]++
// 	}
// }

// if ntice_krok[ntica_anl[0]] == -1 {
// 	ntica_anl[a.n+3+ntice_pos[ntica_anl[0]]] = itoa(0)
// } else {
// 	ntica_anl[a.n+3+ntice_pos[ntica_anl[0]]] = itoa(ntice_krok[ntica_anl[0]])
// }
// ntice_krok[ntica_anl[0]] = 0

// out_stringl.clear()
// out_stringl << akt_r << ntica_anl

// mapaNtice = append(mapaNtice, akt_r, ntica_anl)
// mapaNtice = append(mapaNtice, ntica_anl)
// }
// for _, r := range mapaNtice {
// fmt.Println(r)
// }
// return nil
// }

// func MapaXtice(path string, n int) error {
// 	// f.Seek(0, 0)
// 	file, err := os.Open(path)
// 	if err != nil {
// 		return err
// 	}
// 	r := csv.NewReader(file)
// 	r.Comma = ';'
// 	records, err := r.ReadAll()
// 	if err != nil {
// 		return err
// 	}
//
// 	xlsxFile := xlsx.NewFile()
// 	sheet := xlsxFile.AddSheet("Mapa Xtice")
//
// 	for _, record := range records {
// 		row := sheet.AddRow()
// 		for range record {
// 			cell := row.AddCell()
// 			cell.Value = "Value"
// 			// cell.SetString(c)
// 		}
// 		kombinacia, _ := parse(record, n)
// 		for _, cislo := range kombinacia {
// 			cell := row.AddCell()
// 			cell.SetInt(int(cislo))

// TODO: farby
// style := xlsx.NewStyle()
// switch (int(cislo) - 1) / 10 {
// case 0: // Red
// 	style.Fill = *xlsx.NewFill("solid", "FF0123456X", "FF0123456X")
// case 1: // Green
// 	style.Fill = *xlsx.NewFill("solid", "FF0123456X", "FF0123456X")
// case 2: // Blue
// 	style.Fill = *xlsx.NewFill("solid", "FF0123456X", "FF0123456X")
// case 3: // Yellow
// 	style.Fill = *xlsx.NewFill("solid", "FF0123456X", "FF0123456X")
// case 4: // Cyan
// 	style.Fill = *xlsx.NewFill("solid", "FF0123456X", "FF0123456X")
// case 5: // Dark Red
// 	style.Fill = *xlsx.NewFill("solid", "FF0123456X", "FF0123456X")
// case 6: // Dark Green
// 	style.Fill = *xlsx.NewFill("solid", "FF0123456X", "FF0123456X")
// case 7: // Dark Blue
// 	style.Fill = *xlsx.NewFill("solid", "FF0123456X", "FF0123456X")
// case 8: // Dark Yellow
// 	style.Fill = *xlsx.NewFill("solid", "FF0123456X", "FF0123456X")
// default:
// }
// cell.SetStyle(style)
// }
// }
// return xlsxFile.Save("MapaXtice.xlsx")
// }
