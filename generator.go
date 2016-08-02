package engine

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type Generator interface {
	Start()
	Stop()
	Wait()

	// Progress sleduje stav generatora/filtra vracia stav
	// aktualne prehladavanych skupin v intervale 0.5 sekund.
	// Po skonceni generatora hlasi pocet zapisanych riadkov
	// Viacnasobne volanie tejto funkcie sposobi panic !
	Progress() chan string
	Error() error
	RowsWritten() int
}

type generator2 struct {
	archiv  *Archiv
	filters Filters

	lenSkupiny int
	progress   int
	err        error

	startTime string
	subDir    string

	nworkers       int
	workers        []worker
	workerQueue    chan chan work
	collectorQueue chan []string
	workQueue      chan work
	quit           chan struct{}
	done           chan struct{}

	collected chan int

	writer *CsvMaxWriter

	rowsWritten int
}

func newGenerator2(archiv *Archiv, filters Filters) *generator2 {
	nworkers := runtime.NumCPU()

	return &generator2{
		archiv:  archiv,
		filters: filters,

		startTime: time.Now().Format("2006-1-2-15-4-5"),

		nworkers:       nworkers,
		workers:        make([]worker, nworkers),
		workerQueue:    make(chan chan work, nworkers),
		collectorQueue: make(chan []string, nworkers),
		workQueue:      make(chan work, nworkers),
		quit:           make(chan struct{}),
		done:           make(chan struct{}, 1),

		collected: make(chan int, 1),
	}
}

func NewGenerator2(archiv *Archiv, filters Filters) Generator {
	g := newGenerator2(archiv, filters)
	g.subDir = g.startTime + "_Generator"
	g.writer = NewCsvMaxWriter(g.startTime, g.archiv.WorkingDir,
		setSubdir(g.subDir),
		setSuffixFunc(IntSuffix()),
		setHeader(NewV1(g.archiv).Header),
	)

	g.protokol(g.subDir)
	g.startWorkers()
	g.collect()
	// g.start()

	g.produceGenerator()

	return g
}

func NewFilter2(archiv *Archiv, filters Filters) Generator {
	g := newGenerator2(archiv, filters)
	g.subDir = g.startTime + "_Filter"
	g.writer = NewCsvMaxWriter(g.startTime, g.archiv.WorkingDir,
		setSubdir(g.subDir),
		setSuffixFunc(IntSuffix()),
		setHeader(HeaderV2),
	)

	g.protokol(g.subDir)
	g.startWorkers()
	g.collect()
	// g.start()

	g.produceFilter()

	return g
}

func (g *generator2) Start() {
	go func() {
		defer func() {
			close(g.collectorQueue)
		}()
		for {
			select {
			case worker := <-g.workerQueue:
				work, ok := <-g.workQueue
				if ok {
					worker <- work
				} else {
					g.nworkers--
				}
				if g.nworkers == 0 {
					g.Stop()
				}
			case <-g.quit:
				// zastavenie workerov
				for _, worker := range g.workers {
					worker.stop()
				}

				// pockanie kym dokoncia aktualnu pracu
				if g.nworkers != 0 {
					for range g.workerQueue {
						g.nworkers--
						if g.nworkers == 0 {
							break
						}
					}
				}

				// Treba vypraznit workQueue ak generator
				// skoncil pouzivatel
				go func() {
					for range g.workQueue {
					}
				}()
				return
			}
		}
	}()
}

func (g *generator2) Stop() {
	go func() {
		g.quit <- struct{}{}
	}()
}

func (g *generator2) Wait() {
	<-g.done

	// zatvorime done chan, aby sme viacej krat nemohli cakat
	close(g.done)
}

func (g *generator2) Progress() chan string {
	ch := make(chan string)
	go func() {
		// signal ze sme skoncili
		defer close(ch)

		// zatvorime collected chan aby sme nemohli zavolat viackrat Progress
		defer close(g.collected)
		for {
			select {
			case g.rowsWritten = <-g.collected:
				ch <- fmt.Sprintf("Hotovo. Zapisanych %d riadkov", g.rowsWritten)
				return
			case <-time.After(500 * time.Millisecond):
				ch <- fmt.Sprintf("Prehladavam skupinu %d z %d", g.progress, g.lenSkupiny)
			}
		}
	}()
	return ch
}

func (g *generator2) Error() error {
	return g.err
}

func (g *generator2) RowsWritten() int {
	return g.rowsWritten
}

func (g *generator2) protokol(subdir string) {
	os.Mkdir(filepath.Join(g.archiv.WorkingDir, subdir), 0755)

	s := g.filters.String()
	err := ioutil.WriteFile(filepath.Join(g.archiv.WorkingDir, subdir, "protokol.txt"), []byte(s), 0755)
	if err != nil {
		log.Println(err)
	}
}

func (g *generator2) startWorkers() {
	for i := range g.workers {
		g.workers[i] = newWorker(i+1, g.workerQueue, g.filters, g.archiv.n)
		g.workers[i].start()
	}
}

func (g *generator2) produceGenerator() {
	go func() {
		defer close(g.workQueue)

		var skupiny Skupiny
		for _, s := range g.archiv.Skupiny {
			if !g.filters.CheckSkupina(s) {
				continue
			}
			skupiny = append(skupiny, s)
		}

		g.lenSkupiny = len(skupiny)

		vystup := NewV1(g.archiv)

		for i, s := range skupiny {
			g.progress = i + 1
			g.workQueue <- work{
				c:  cisla(g.archiv.Hrx.Cisla, s.Xcisla),
				ch: g.collectorQueue,
				v1: &vystup,
			}
		}
	}()
}

func (g *generator2) produceFilter() {
	go func() {
		defer close(g.workQueue)

		var skupiny Skupiny
		for _, s := range g.archiv.Skupiny {
			if !g.filters.CheckSkupina(s) {
				continue
			}
			skupiny = append(skupiny, s)
		}

		g.lenSkupiny = len(skupiny)
		for i, s := range skupiny {
			g.progress = i + 1

			v2 := NewV2(g.archiv, s)

			g.workQueue <- work{
				c:  cisla(g.archiv.Hrx.Cisla, s.Xcisla),
				ch: g.collectorQueue,
				v2: &v2,
			}
		}
	}()
}

func (g *generator2) collect() {
	go func() {
		w := g.writer
		defer func() {
			if err := w.Close(); err != nil {
				log.Println(err)
				g.err = err
			}
			// progresu posleme pocet zapisanych riadkov
			g.collected <- w.TotalRowsWriten()

			// skoncili sme az ked je vsetko zapisane
			g.done <- struct{}{}
		}()

		for r := range g.collectorQueue {
			if err := w.Write(r); err != nil {
				log.Println(err)
				g.err = err
				g.Stop()
			}
		}
	}()
}

type work struct {
	c  []cislo
	ch chan []string

	v1 *V1
	v2 *V2
}

type worker struct {
	id          int
	work        chan work
	workerQueue chan chan work
	quit        chan struct{}

	k       kombinator
	filters Filters
	n       int
}

func newWorker(id int, workerQueue chan chan work, filters Filters, n int) worker {
	return worker{
		id:          id,
		work:        make(chan work),
		workerQueue: workerQueue,
		quit:        make(chan struct{}),

		k:       kombinator{},
		filters: filters,
		n:       n,
	}
}

func (w worker) start() {
	go func() {
		// pridame sa do fronty
		w.workerQueue <- w.work
		for {
			select {
			case work := <-w.work:
				// vykoname pracu na zvlast goroutine
				// aby sme mohli v pripade potreby skoncit
				go func() {
					ch := w.k.run(work.c, w.filters, w.n)
					// vykoname pracu
					if work.v1 != nil {
						for k := range ch {
							work.ch <- work.v1.Riadok(k)
							ch <- nil
						}
					} else if work.v2 != nil {
						for k := range ch {
							work.v2.Add(k)
							ch <- nil
						}
						if !work.v2.Empty() {
							work.ch <- work.v2.Riadok()
						}
					}

					// po skoncime sa pridame do fronty
					w.workerQueue <- w.work
				}()
			case <-w.quit:
				// skoncili sme
				w.k.stop()
				return
			}
		}
	}()
}

func (w worker) stop() {
	go func() {
		w.quit <- struct{}{}
	}()
}

type kombinator struct {
	done bool
}

func (g kombinator) run(cisla []cislo, filters Filters, n int) chan Kombinacia {
	ch := make(chan Kombinacia)
	go func() {
		var (
			indices = make([]int, 1, n)
			k       = make(Kombinacia, 0, n)
		)
		for len(indices) > 0 && !g.stopped() {
			j := len(indices)

			// i je index daneho cisla
			i := indices[j-1]

			// na tomto leveli uz nie su dalsie cisla
			// ideme o level nizsie
			if i == len(cisla) {
				indices = indices[:j-1]
				continue
			}

			// skusime cislo
			cislo := cisla[i]

			// v predchadzajucom kroku sme nasli kombinaciu
			// skusime dalsie cislo na tomto leveli
			if k.Len() == j && cislo.cislo == k[k.Len()-1] {
				k.Pop()
				cislo.increment()
				indices[j-1]++
				continue
			}

			// ak pocet cisiel z danej hrx skupiny
			// je vacsi ako 0, berieme cislo do kombinacie
			// a znizime pocet cisiel v skupine
			if cislo.zeroGroup() {
				indices[j-1]++
				continue
			}

			k.Append(cislo.cislo)
			cislo.decrement()

			// ak kombinacia nevyhovuje filtru
			// skusime dalsie cislo
			if filters != nil && !filters.Check(k) {
				continue
			}

			// cisel v kombinacii este nie je n
			// skusime dalsie cislo
			if k.Len() < n {
				indices = append(indices, i+1)
				continue
			}
			ch <- k
			<-ch
		}
		close(ch)
	}()
	return ch
}

func (k *kombinator) stop() {
	// log.Println("kombinator: stop()")
	k.done = true
}

func (k *kombinator) stopped() bool {
	return k.done
}

func cisla(hrxNums Nums, xcisla Xcisla) []cislo {
	var (
		cisla []cislo
		skMax = make(map[int]*int, xcisla.Len())
	)

	for _, tab := range xcisla {
		pocet := tab.Max
		skMax[tab.Sk] = &pocet
	}

	for _, Num := range hrxNums {
		sk := Num.PocetR()
		if pocet, ok := skMax[sk]; ok {
			cisla = append(cisla, cislo{
				cislo: byte(Num.Cislo()),
				pocet: pocet,
			})
		}
	}

	return cisla
}

type cislo struct {
	cislo byte
	pocet *int
}

func (c *cislo) debug() {
	fmt.Printf("%2d: %2d(%p)\n", c.cislo, *c.pocet, c.pocet)
}

func (c cislo) zeroGroup() bool {
	return *c.pocet == 0
}

func (c *cislo) increment() {
	*c.pocet++
}

func (c *cislo) decrement() {
	*c.pocet--
}
