package fastfilter

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"gitlab.com/melias122/engine"
	"gitlab.com/melias122/engine/filter"
	"gitlab.com/melias122/engine/sieve"
)

type FastFilter struct {
	sieve.Producer

	filters filter.Filters
	skupiny engine.Skupiny

	r *result
}

func New(archiv *engine.Archiv, filters filter.Filters) *FastFilter {
	startTime := time.Now().Format("2006-1-2-15-4-5")
	subdir := startTime + "_FastFilter"

	csvw := engine.NewCsvMaxWriter(startTime, archiv.WorkingDir,
		engine.SetSubdir(subdir),
		engine.SetSuffixFunc(engine.IntSuffix()),
		engine.SetHeader(engine.HeaderHrxHHrx),
	)

	os.Mkdir(filepath.Join(archiv.WorkingDir, subdir), 0755)
	filename := filepath.Join(archiv.WorkingDir, subdir, "protokol.txt")
	ioutil.WriteFile(filename, []byte(filters.String()), 0755)

	// fast filter
	skupiny := make(engine.Skupiny, 0, len(archiv.Skupiny))
	for _, s := range archiv.Skupiny {
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
