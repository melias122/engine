package hrx

import "testing"

func TestContains(t *testing.T) {
	tests := []struct {
		presun Presun
		test   map[Tab]bool
	}{
		{presun: Presun{}, test: map[Tab]bool{Tab{}: false, Tab{1, 1}: false}},
		{presun: Presun{{0, 0}}, test: map[Tab]bool{Tab{0, 0}: true, Tab{1, 1}: false}},
		{
			presun: Presun{
				{1, 1}, {2, 2}, {4, 0},
			},
			test: map[Tab]bool{
				Tab{0, 0}: false, Tab{1, 1}: true, Tab{2, 2}: true, Tab{3, 3}: false,
				Tab{3, 0}: false, Tab{3, 1}: false, Tab{4, 0}: true, Tab{4, 1}: false,
				Tab{5, 1}: false,
			},
		},
	}
	for _, test := range tests {
		for tab, exp := range test.test {
			if ok := test.presun.Contains(tab); ok != exp {
				t.Fatalf("Expected: (%v:%v), got: (%v:%v)", tab, exp, test.presun, ok)
			}
		}
	}
}
