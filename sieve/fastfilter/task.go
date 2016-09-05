package fastfilter

import (
	"gitlab.com/melias122/engine"
	"gitlab.com/melias122/engine/filter"
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
