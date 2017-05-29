package fastfilter

import (
	"github.com/melias122/engine/engine"
	"github.com/melias122/engine/filter"
)

type task struct {
	filters filter.Filters
	skupina engine.Skupina
}

func (t *task) Run() error {
	t.filters.CheckSkupina(t.skupina)
	return nil
}

func (t *task) Cancel() {}
