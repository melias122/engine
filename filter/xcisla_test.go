package filter

import (
	"testing"

	"gitlab.com/melias122/engine"
)

var xcislaTest = []engine.Skupina{
	{Xcisla: engine.Xcisla{{1, 5}}},
	{Xcisla: engine.Xcisla{{1, 1}, {2, 1}, {3, 1}, {5, 2}}},
	{Xcisla: engine.Xcisla{{1, 1}, {2, 1}, {3, 1}, {6, 2}}},
	{Xcisla: engine.Xcisla{{1, 1}, {2, 1}, {3, 1}, {7, 2}}},
	{Xcisla: engine.Xcisla{{1, 1}, {2, 1}, {3, 1}, {8, 2}}},
	{Xcisla: engine.Xcisla{{1, 1}, {2, 2}, {3, 1}}},
	{Xcisla: engine.Xcisla{{1, 2}, {2, 2}, {3, 1}}},
	{Xcisla: engine.Xcisla{{1, 2}, {2, 1}, {3, 1}}},
	{Xcisla: engine.Xcisla{{1, 2}, {2, 1}, {3, 1}, {10, 1}}},
}

func TestFilterXcisla(t *testing.T) {
	tabs := engine.Xcisla{
		{1, 1},
		{2, 2},
		{2, 1},
		{1, 2},
	}
	filter := NewFilterXcisla(tabs)
	for _, s := range xcislaTest {
		if ok := filter.CheckSkupina(s); !ok {
			t.Fatal(s.Xcisla)
		}
	}
}
