package fastfilter

import (
	"context"
	"testing"

	"github.com/melias122/engine/engine"
	. "github.com/melias122/engine/filter"
	"github.com/melias122/engine/sieve"
)

func Test745(t *testing.T) {
	n, m := 7, 45
	a, err := engine.NewArchiv("../../testdata/745.csv", "../../testdata", n, m)
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

		NewFilterCislovackyRange(3, 3, engine.N, n),
		NewFilterCislovackyRange(3, 3, engine.Pr, n),
		NewFilterCislovackyRange(5, 5, engine.Mc, n),
		NewFilterCislovackyRange(2, 2, engine.C19, n),
		NewFilterCislovackyRange(0, 0, engine.C0, n),
		NewFilterCislovackyRange(3, 3, engine.XcC, n),
		NewFilterCislovackyRange(1, 1, engine.Cc, n),
		NewFilterCislovackyRange(1, 1, engine.CC, n),
		NewFilterZhodaRange(0, 0, a.K, n),

		NewFilterSmernica(0.8037878788, 0.8037878788, n, m),
		NewFilterKorelacia(0.488756203, 0.488756203, a.K, n, m),

		NewFilterZakazane(Ints{int(a.Uc.Cislo)}, n, m),
		NewFilterZakazaneSTL(MapInts{1: {1}}, n, m),

		NewFilterPovinne(Ints{2, 9, 16, 22, 23, 24, 43}, n, m),
		NewFilterPovinneSTL(MapInts{1: {2}, 2: {9}, 3: {16}, 4: {22}, 5: {23}, 6: {24}, 7: {43}}, n, m),

		NewFilterNtica(n, engine.Tica{4, 0, 1, 0, 0, 0, 0}),
		NewFilterXtica(n, m, engine.Tica{2, 1, 3, 0, 1}),
	}

	ff := New(a, filters)
	s, err := sieve.New(ff, nil)
	if err != nil {
		t.Fatal(err)
	}

	s.Start(context.Background())
	<-s.Done()
	if s.Error() != nil {
		t.Log(s.Error())
	}

	if ff.Found() != "1" {
		t.Fail()
	}
}
