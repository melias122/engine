package generator

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/melias122/engine/archiv"
	"github.com/melias122/engine/csv"
	"github.com/melias122/engine/engine"
	"github.com/melias122/engine/filter"
	"github.com/melias122/engine/sieve"
)

type Generator struct {
	archiv  *archiv.Archiv
	filters filter.Filters
	skupiny engine.Skupiny

	r *result
	n int
}

func New(archiv *archiv.Archiv, sk engine.Skupiny, filters filter.Filters, n, m int) *Generator {
	startTime := time.Now().Format("2006-1-2-15-4-5")
	subdir := startTime + "_Generator"

	csvw := csv.NewCsvMaxWriter(startTime, archiv.Workdir,
		csv.SetSubdir(subdir),
		csv.SetSuffixFunc(csv.IntSuffix()),
		csv.SetHeader(newResultFilter(nil, archiv, n, m).header),
	)

	os.Mkdir(filepath.Join(archiv.Workdir, subdir), 0755)
	filename := filepath.Join(archiv.Workdir, subdir, "protokol.txt")
	ioutil.WriteFile(filename, []byte(filters.String()), 0755)

	// fast filter
	var skupiny engine.Skupiny
	for _, s := range sk {
		if !filters.CheckSkupina(s) {
			continue
		}
		skupiny = append(skupiny, s)
	}

	resultFilter := newResultFilter(csvw, archiv, n, m)
	filters = append(filters, resultFilter)

	return &Generator{
		archiv:  archiv,
		filters: filters,
		skupiny: skupiny,
		r:       resultFilter,
		n:       n,
	}
}

func (g *Generator) Start(ctx context.Context) <-chan sieve.Task {
	tasks := make(chan sieve.Task)
	go func() {
		defer close(tasks)
		for _, s := range g.skupiny {
			select {
			case <-ctx.Done():
				return
			case tasks <- &task{
				n:       g.n,
				hrxNums: g.archiv.Hrx.Cisla,
				xcisla:  s.Xcisla,
				filters: g.filters,
			}:
			}
		}
	}()
	return tasks
}

func (g *Generator) Count() int {
	return len(g.skupiny)
}

func (g *Generator) Close() error {
	return g.r.Close()
}

func (g *Generator) Found() string {
	return g.r.String()
}
