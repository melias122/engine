package psl

import "bytes"

func nticaPozicie(k Kombinacia) []string {
	var s []string
	for i, ok := range NticaPozicie(k) {
		if ok == 1 {
			s = append(s, itoa(int(k[i])))
		} else {
			s = append(s, "")
		}
	}
	return s
}

func krokNtica(ntica Tica, nticeVsetky []string, counter map[string]int) []string {
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
		riadok = append(riadok, NticaSucet(r.K).String())
		riadok = append(riadok, NticaSucin(r.K).String())
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

	w := NewCsvMaxWriter("MapaNtice", a.WorkingDir, setHeader(header))
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

func nticeNtice(n int) []Tica {
	var nticeVsetky []Tica
	ntice(n, func(tica []int) {
		n := make(Tica, len(tica))
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
