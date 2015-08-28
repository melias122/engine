package archiv

// func ntice(n int, emit func([]int)) {
// 	if n < 2 {
// 		panic("ntice: n < 2")
// 	}
// 	tica := make([]int, n)
// 	end := make([]int, n)
// 	tica[0] = n
// 	end[n-1] = 1
//
// 	i := 0
// 	for {
// 		sum := 0
// 		for j, e := range tica {
// 			sum += int(e) * (j + 1)
// 		}
// 		if sum == n {
// 			emit(tica)
// 			tica[i]--
// 			i++
// 		} else if sum < n {
// 			tica[i]++
// 		} else {
// 			tica[i]--
// 			i++
// 		}
// 		if i == n {
// 			i--
// 			for tica[i] == 0 {
// 				i--
// 			}
// 			tica[i]--
// 			i++
// 		}
//
// 		equal := true
// 		for p := 0; p < n; p++ {
// 			if tica[p] != end[p] {
// 				equal = false
// 			}
// 		}
// 		if equal {
// 			break
// 		}
// 	}
// 	emit(tica)
// }

// func nticeStr(n int) []string {
// 	var (
// 		tice []string
// 		buf  bytes.Buffer
// 	)
// 	ntice(n, func(tica []int) {
// 		buf.Reset()
// 		for i, t := range tica {
// 			if i > 0 {
// 				buf.WriteString(" ")
// 			}
// 			buf.WriteString(itoa(t))
// 		}
// 		tice = append(tice, buf.String())
// 	})
// 	return tice
// }

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

// void mapaZhoda(CSV csv, uint n){
// func MapaZhoda(csv [][]string, n int) error {

// // CSV mapaZhoda;
// var (
// 	mapaZhoda                   [][]string
// 	zhoda, l_len                int
// 	krok                        int = -1
// 	out_stringl, header, zh_anl []string
// 	prd, akt                    []int
// 	zh_pos                      map[int]int
// )

// // dlzka riadku
// if csv.size() > 0 {
// 	l_len = csv[1].length()
// } else {
// 	l_len = csv.front().length()
// }

// // header
// header = append(header, "Zhoda")
// for i := 1; i <= n; i++ {
// 	header = append(header, itoa(i))
// }
// header = append(header, "Pozicia ZH", "Krok")
// for i := 1; i <= n; i++ {
// 	header = append(header, fmt.Sprint("%d Stl r-1/r", i))
// }

// // TODO: header
// // out_stringl << csv.front() << header;
// // mapaZhoda.append(out_stringl);
// // header end

// //header pop
// csv.pop_front()

// for i := 0; i <= n; i++ {
// 	zhoda_pocet.insert(i, 0)
// }

// // foreach (const QStringList &akt_r, csv){
// for _, akt_r := range csv {

// 	if !akt.isEmpty() {
// 		akt.clear()
// 	}

// 	for i := 3; i < n+3; i++ {
// 		akt.push_back(akt_r[i].toInt())
// 	}

// 	zhoda = 0
// 	zh_anl.clear()
// 	zh_pos.clear()

// 	for i := 0; i < 2*n+3; i++ {
// 		zh_anl = append(zh_anl, "")
// 	}

// 	for i := 0; i < akt.size() && !prd.empty(); i++ {
// 		for j := 0; j < prd.size(); j++ {
// 			if akt[i] == prd[j] {
// 				zhoda++
// 				zh_pos.insert(i, j)
// 			}
// 		}
// 	}

// 	// Zhoda
// 	zh_anl[0] = itoa(zhoda)

// 	// Krok
// 	if zhoda > 0 {
// 		if krok == -1 {
// 			zh_anl[n+2] = itoa(0)
// 		} else {
// 			zh_anl[n+2] = itoa(krok)
// 		}
// 		krok = 1
// 	} else if krok >= 0 {
// 		krok++
// 	}

// 	var pos_zh string

// 	// TODO: sort
// 	// auto zhPosKeys = zh_pos.keys();
// 	// if(!zhPosKeys.isEmpty())
// 	// qSort(zhPosKeys);

// 	for _, i := range zhPosKeys {
// 		// 1-n
// 		mapaZhoda.back()[l_len+1+zh_pos[i]] = itoa(prd[zh_pos[i]])
// 		zh_anl[i+1] = itoa(akt[i])

// 		// Pozicia zhoda
// 		pos_zh = itoa(zh_pos[i] + 1).append("|").append(itoa(i + 1))
// 		if zh_anl[n+1].size() > 0 {
// 			zh_anl[n+1].append(", ").append(pos_zh)
// 		} else {
// 			zh_anl[n+1].append(pos_zh)
// 		}

// 		zh_anl[n+3+i] = pos_zh
// 	}

// 	zhs := zh_anl[n+1]

// 	var zhoda_kolko int
// 	if !zhs.isEmpty() {
// 		zhoda_kolko = zhs.split(",").size()
// 	} else {
// 		zhoda_kolko = 0
// 	}
// 	zhoda_pocet[zhoda_kolko]++

// 	if zhoda_kolko > 0 { // len zhoda 1,2,3,..
// 		if !zhoda_typ_pocet[zhoda_kolko].contains(zhs) {
// 			zhoda_typ_pocet[zhoda_kolko].insert(zhs, 1)
// 		} else {
// 			zhoda_typ_pocet[zhoda_kolko][zhs]++
// 		}
// 	}

// 	out_stringl.clear()
// 	out_stringl << akt_r << zh_anl

// 	mapaZhoda.push_back(QStringList(out_stringl))

// 	prd = akt
// }

// exportCsv(mapaZhoda, pwd()+"MapaZhoda_"+suborName()+".csv")

// 	return nil
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
