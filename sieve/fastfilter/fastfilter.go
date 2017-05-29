package fastfilter

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/melias122/engine/csv"
	"github.com/melias122/engine/engine"
	"github.com/melias122/engine/filter"
	"github.com/melias122/engine/sieve"
)

type FastFilter struct {
	sieve.Producer

	filters filter.Filters
	skupiny engine.Skupiny

	r *result
}

func New(archiv *engine.Archiv, sk engine.Skupiny, filters filter.Filters) *FastFilter {
	startTime := time.Now().Format("2006-1-2-15-4-5")
	subdir := startTime + "_FastFilter"

	csvw := csv.NewCsvMaxWriter(startTime, archiv.Workdir,
		csv.SetSubdir(subdir),
		csv.SetSuffixFunc(csv.IntSuffix()),
		csv.SetHeader(engine.HeaderHrxHHrx),
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
	resultFilter := newResultFilter(csvw, archiv)
	filters = append(filters, resultFilter)

	return &FastFilter{
		filters: filters,
		skupiny: skupiny,
		r:       resultFilter,
	}
}

func (f *FastFilter) Start(ctx context.Context) <-chan sieve.Task {
	tasks := make(chan sieve.Task)
	go func() {
		defer close(tasks)
		for _, s := range f.skupiny {
			select {
			case <-ctx.Done():
				return
			case tasks <- &task{
				filters: f.filters,
				skupina: s,
			}:
				log.Println("sending task")
			}
		}
	}()
	return tasks
}

func (f *FastFilter) Count() int {
	return len(f.skupiny)
}

func (f *FastFilter) Close() error {
	return f.r.Close()
}

func (f *FastFilter) Found() string {
	return f.r.String()
}
