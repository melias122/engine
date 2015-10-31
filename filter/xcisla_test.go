package filter

import (
	"testing"

	"github.com/melias122/psl/hrx"
)

var xcislaTest = []hrx.Skupina{
	{Presun: hrx.Presun{{1, 5}}},
	{Presun: hrx.Presun{{1, 1}, {2, 1}, {3, 1}, {5, 2}}},
	{Presun: hrx.Presun{{1, 1}, {2, 1}, {3, 1}, {6, 2}}},
	{Presun: hrx.Presun{{1, 1}, {2, 1}, {3, 1}, {7, 2}}},
	{Presun: hrx.Presun{{1, 1}, {2, 1}, {3, 1}, {8, 2}}},
	{Presun: hrx.Presun{{1, 1}, {2, 2}, {3, 1}}},
	{Presun: hrx.Presun{{1, 2}, {2, 2}, {3, 1}}},
	{Presun: hrx.Presun{{1, 2}, {2, 1}, {3, 1}}},
	{Presun: hrx.Presun{{1, 2}, {2, 1}, {3, 1}, {10, 1}}},
}

func TestXcisla(t *testing.T) {
	tabs := hrx.Presun{
		{1, 1},
		{2, 2},
		{2, 1},
		{1, 2},
	}
	filter := Xcisla(tabs)
	for _, s := range xcislaTest {
		if ok := filter.CheckSkupina(s); ok {
			t.Log(s.Presun)
		}
	}
}
