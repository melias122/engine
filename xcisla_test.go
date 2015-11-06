package psl

import (
	"testing"
)

var xcislaTest = []Skupina{
	{Xcisla: Xcisla{{1, 5}}},
	{Xcisla: Xcisla{{1, 1}, {2, 1}, {3, 1}, {5, 2}}},
	{Xcisla: Xcisla{{1, 1}, {2, 1}, {3, 1}, {6, 2}}},
	{Xcisla: Xcisla{{1, 1}, {2, 1}, {3, 1}, {7, 2}}},
	{Xcisla: Xcisla{{1, 1}, {2, 1}, {3, 1}, {8, 2}}},
	{Xcisla: Xcisla{{1, 1}, {2, 2}, {3, 1}}},
	{Xcisla: Xcisla{{1, 2}, {2, 2}, {3, 1}}},
	{Xcisla: Xcisla{{1, 2}, {2, 1}, {3, 1}}},
	{Xcisla: Xcisla{{1, 2}, {2, 1}, {3, 1}, {10, 1}}},
}

func TestFilterXcisla(t *testing.T) {
	tabs := Xcisla{
		{1, 1},
		{2, 2},
		{2, 1},
		{1, 2},
	}
	filter := NewFilterXcisla(tabs)
	for _, s := range xcislaTest {
		if ok := filter.CheckSkupina(s); ok {
			t.Log(s.Xcisla)
		}
	}
}
