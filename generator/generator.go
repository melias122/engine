package generator

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"github.com/melias122/psl/archiv"
	"github.com/melias122/psl/filter"
	"github.com/melias122/psl/hrx"
	"github.com/melias122/psl/komb"
	"github.com/melias122/psl/num"
	"github.com/melias122/psl/rw"
)

func protokol(wdir, subdir string, filters filter.Filters) {

	os.Mkdir(filepath.Join(wdir, subdir), 0755)

	s := filters.String()
	err := ioutil.WriteFile(filepath.Join(wdir, subdir, "protokol.txt"), []byte(s), 0755)
	if err != nil {
		log.Println(err)
	}
}

func GenerateKombinacie(n int, a *archiv.Archiv, filters filter.Filters, msg chan string) {
	var (
		wg, wg2        sync.WaitGroup
		chanKombinacie = make(chan komb.Kombinacia, 16)
		chanPresuny    = make(chan hrx.Presun, 2)
		chanRiadok     = make(chan []string, 8)
		vystup         = archiv.NewV1(a)
	)

	fileName := time.Now().Format("2006-1-2-15-4-5")
	w := rw.NewCsvMaxWriter(a.WorkingDir, fileName, [][]string{vystup.Header})
	w.Suffix = rw.IntSuffix()
	w.SubDir = fileName + "_Generator"
	defer w.Close()

	protokol(a.WorkingDir, w.SubDir, filters)

	go func() {
		defer close(chanPresuny)
		for _, s := range a.Skupiny {
			if filters.CheckSkupina(s) {
				chanPresuny <- s.Presun
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
	msg <- fmt.Sprintf("Nájdených %s kombinácií", w.NWrites.String())
}

func GenerateFilter(n int, a *archiv.Archiv, filters filter.Filters, msg chan string) {
	var (
		wg          sync.WaitGroup
		chanRiadok  = make(chan []string, 16)
		chanSkupina = make(chan hrx.Skupina, 2)
	)

	fileName := time.Now().Format("2006-1-2-15-4-5")
	w := rw.NewCsvMaxWriter(a.WorkingDir, fileName, [][]string{archiv.HeaderV2})
	w.Suffix = rw.IntSuffix()
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
					chanKombinacie = make(chan komb.Kombinacia, 16)
					generator      = NewGenerator(n, a.Hrx.Cisla, chanKombinacie, filters)
					vystup         = archiv.NewV2(a, skupina)
				)
				go func(skupina hrx.Skupina) {
					defer close(chanKombinacie)
					generator.Generate(skupina.Presun)
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
	msg <- fmt.Sprintf("Nájdených %s riadkov", w.NWrites.String())
}

type Generator struct {
	n       int
	nums    num.Nums
	ch      chan komb.Kombinacia
	filters filter.Filters
}

func NewGenerator(n int, HrxNums num.Nums, ch chan komb.Kombinacia, f ...filter.Filter) *Generator {
	return &Generator{
		n:       n,
		nums:    HrxNums,
		ch:      ch,
		filters: f,
	}
}

func (g *Generator) Generate(p hrx.Presun) {
	g.generate(newCisla(g.nums, p), make(komb.Kombinacia, 0, g.n))
}

func (g *Generator) generate(cisla cisla, k komb.Kombinacia) {
	for i, c := range cisla {
		if *c.pocet == 0 {
			continue
		}
		k = append(k, c.cislo)
		*c.pocet--
		if g.filters.Check(k) {
			if len(k) == g.n {
				cp := make(komb.Kombinacia, len(k))
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
