package psl

import "testing"

// func TestNewArchiv(t *testing.T) {
// 	currentWorkingDir, err := ioutil.TempDir("profile", "")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Log(currentWorkingDir)

// 	removeTmpDir := func() {
// 		if err := os.RemoveAll(currentWorkingDir); err != nil {
// 			t.Fatal(err)
// 		}
// 	}
// 	defer removeTmpDir()

// 	a, err := NewArchiv("profile/535.csv", currentWorkingDir, 5, 35)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	files, err := ioutil.ReadDir(a.WorkingDir)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	for _, file := range files {
// 		t.Log(file.Name())
// 	}
// }

func TestZnamyRiadok745(t *testing.T) {
	n, m := 7, 45
	a, err := NewArchiv("profile/745_r.csv", DiscardCSV, n, m)
	if err != nil {
		t.Fatal(err)
	}
	filters := Filters{
		NewFilterR1(0.000341972264831028, 0.000341972264831028, a.HHrx.Cisla, n),
		NewFilterSTL1(0.0003266465, 0.0003266465, a.HHrx.Cisla, n),
		NewFilterHHrx(0.0989396766512104, 0.0989396766512104, a.HHrx, n),

		NewFilterR2(9.34969738146E-006, 9.349697381461E-006, a.Hrx.Cisla, n),
		NewFilterSTL2(1.58809155486105E-005, 1.58809155486105E-005, a.Hrx.Cisla, n),
		NewFilterHrx(34.8651270233453, 34.8651270233453, a.Hrx, n),

		NewFilterSucet(139, 139, n),

		NewFilterCislovackyRange(n, 3, 3, N),
		NewFilterCislovackyRange(n, 3, 3, Pr),
		NewFilterCislovackyRange(n, 5, 5, Mc),
		NewFilterCislovackyRange(n, 2, 2, C19),
		NewFilterCislovackyRange(n, 0, 0, C0),
		NewFilterCislovackyRange(n, 3, 3, XcC),
		NewFilterCislovackyRange(n, 1, 1, Cc),
		NewFilterCislovackyRange(n, 1, 1, CC),
		NewFilterZhodaRange(0, 0, a.K, n),

		NewFilterSmernica(0.8037878788, 0.8037878788, n, m),
		NewFilterKorelacia(0.488756203, 0.488756203, a.K, n, m),

		// filterCifrovacky{n: n, c: c},

		NewFilterZakazane(Ints{int(a.Uc.Cislo)}, n, m),
		NewFilterZakazaneSTL(MapInts{1: {1}}, n, m),

		NewFilterPovinne(Ints{2, 9, 16, 22, 23, 24, 43}, n, m),
		NewFilterPovinneSTL(MapInts{1: {2}, 2: {9}, 3: {16}, 4: {22}, 5: {23}, 6: {24}, 7: {43}}, n, m),

		NewFilterNtica(n, Tica{4, 0, 1, 0, 0, 0, 0}),
		NewFilterXtica(n, m, Tica{2, 1, 3, 0, 1}),
		// NewFilterXcisla()
	}
	g := NewGenerator2(a, filters)
	g.Start()
	for msg := range g.Progress() {
		t.Log(msg)
	}

	if g.RowsWriten != 1 {
		t.Fatalf("Generator mal najst 1 kombinaciu ale nasiel %d", g.RowsWriten)
	}

	f := NewFilter2(a, filters)
	f.Start()
	for msg := range f.Progress() {
		t.Log(msg)
	}

	if f.RowsWriten != 1 {
		t.Fatalf("Filter mal najst 1 kombinaciu ale nasiel %d", f.RowsWriten)
	}
}
