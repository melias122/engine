package fastfilter

import (
	"fmt"
	"sync"

	"github.com/melias122/engine"
)

type syncWriter struct {
	sync.Mutex
	w *engine.CsvMaxWriter
}

func newSyncWriter(w *engine.CsvMaxWriter) *syncWriter {
	return &syncWriter{
		w: w,
	}
}

func (s *syncWriter) Write(record []string) (err error) {
	s.Lock()
	defer s.Unlock()
	return s.w.Write(record)
}

func (s *syncWriter) Close() error {
	s.Lock()
	defer s.Unlock()
	return s.w.Close()
}

type result struct {
	w   *syncWriter
	err error
}

func newResultFilter(w *engine.CsvMaxWriter, a *engine.Archiv) *result {
	return &result{
		w: newSyncWriter(w),
	}
}

func (f *result) Check(engine.Kombinacia) bool {
	return true
}

func (f *result) CheckSkupina(s engine.Skupina) bool {
	if err := f.w.Write(s.Record()); err != nil {
		if f.err == nil {
			f.err = err
		}
		return false
	}
	return true
}

func (f *result) String() string {
	return fmt.Sprint(f.w.w.TotalRowsWriten())
}

func (r *result) Close() error {
	return r.w.Close()
}
