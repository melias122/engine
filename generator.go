package psl

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func protokol(wdir, subdir string, filters Filters) {

	os.Mkdir(filepath.Join(wdir, subdir), 0755)

	s := filters.String()
	err := ioutil.WriteFile(filepath.Join(wdir, subdir, "protokol.txt"), []byte(s), 0755)
	if err != nil {
		log.Println(err)
	}
}

func GenerateKombinacie(n int, a *Archiv, filters Filters, msg chan string) {
	var (
		wg, wg2        sync.WaitGroup
		chanKombinacie = make(chan Kombinacia, 16)
		chanPresuny    = make(chan Xcisla, 2)
		chanRiadok     = make(chan []string, 8)
		vystup         = NewV1(a)
	)

	fileName := time.Now().Format("2006-1-2-15-4-5")
	w := NewCsvMaxWriter(a.WorkingDir, fileName, [][]string{vystup.Header})
	w.Suffix = IntSuffix()
	w.SubDir = fileName + "_Generator"
	defer w.Close()

	protokol(a.WorkingDir, w.SubDir, filters)

	go func() {
		defer close(chanPresuny)
		for _, s := range a.Skupiny {
			if filters.CheckSkupina(s) {
				chanPresuny <- s.Xcisla
			}
		}
	}()

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var (
				generator = NewGenerator(n, a.Hrx.Cisla, chanKombinacie, filters)
			)
			for p := range chanPresuny {
				generator.Generate(p)
			}
		}()
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			for k := range chanKombinacie {
				chanRiadok <- vystup.Riadok(k)
			}
		}()
	}
	go func() {
		defer close(chanKombinacie)
		wg.Wait()
	}()
	go func() {
		defer close(chanRiadok)
		wg2.Wait()
	}()
	for r := range chanRiadok {
		if err := w.Write(r); err != nil {
			log.Println(err)
		}
	}
	msg <- fmt.Sprintf("Nájdených %s kombinácií", strconv.FormatUint(w.NWrites, 10))
}

func GenerateFilter(n int, a *Archiv, filters Filters, msg chan string) {
	var (
		wg          sync.WaitGroup
		chanRiadok  = make(chan []string, 16)
		chanSkupina = make(chan Skupina, 2)
	)

	fileName := time.Now().Format("2006-1-2-15-4-5")
	w := NewCsvMaxWriter(a.WorkingDir, fileName, [][]string{HeaderV2})
	w.Suffix = IntSuffix()
	w.SubDir = fileName + "_Filter"
	defer w.Close()

	protokol(a.WorkingDir, w.SubDir, filters)

	go func() {
		defer close(chanSkupina)
		for _, s := range a.Skupiny {
			if filters.CheckSkupina(s) {
				chanSkupina <- s
			}
		}
	}()

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for skupina := range chanSkupina {
				var (
					chanKombinacie = make(chan Kombinacia, 16)
					generator      = NewGenerator(n, a.Hrx.Cisla, chanKombinacie, filters)
					vystup         = NewV2(a, skupina)
				)
				go func(skupina Skupina) {
					defer close(chanKombinacie)
					generator.Generate(skupina.Xcisla)
				}(skupina)
				for k := range chanKombinacie {
					vystup.Add(k)
				}
				if !vystup.Empty() {
					chanRiadok <- vystup.Riadok()
				}
			}
		}()
	}
	go func() {
		defer close(chanRiadok)
		wg.Wait()
	}()
	for r := range chanRiadok {
		if err := w.Write(r); err != nil {
			log.Println(err)
		}
	}
	msg <- fmt.Sprintf("Nájdených %s riadkov", strconv.FormatUint(w.NWrites, 10))
}

type Generator struct {
	n       int
	nums    Nums
	ch      chan Kombinacia
	filters Filters
}

func NewGenerator(n int, HrxNums Nums, ch chan Kombinacia, f ...Filter) *Generator {
	return &Generator{
		n:       n,
		nums:    HrxNums,
		ch:      ch,
		filters: f,
	}
}

func (g *Generator) Generate(p Xcisla) {
	g.generate(newCisla(g.nums, p), make(Kombinacia, 0, g.n))
}

func (g *Generator) generate(cisla cisla, k Kombinacia) {
	for i, c := range cisla {
		if *c.pocet == 0 {
			continue
		}
		k = append(k, c.cislo)
		*c.pocet--
		if g.filters.Check(k) {
			if len(k) == g.n {
				cp := make(Kombinacia, len(k))
				copy(cp, k)
				g.ch <- cp
			} else {
				g.generate(cisla[i+1:], k)
			}
		}
		k = k[:len(k)-1]
		*c.pocet++
	}
}
