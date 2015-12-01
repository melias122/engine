package psl

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestNewArchiv(t *testing.T) {
	currentWorkingDir, err := ioutil.TempDir("profile", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(currentWorkingDir)

	removeTmpDir := func() {
		if err := os.RemoveAll(currentWorkingDir); err != nil {
			t.Fatal(err)
		}
	}
	defer removeTmpDir()

	a, err := NewArchiv("profile/535.csv", currentWorkingDir, 5, 35)
	if err != nil {
		t.Fatal(err)
	}
	files, err := ioutil.ReadDir(a.WorkingDir)
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range files {
		t.Log(file.Name())
	}
}

func TestNewArchivNoOutputs(t *testing.T) {
	a, err := NewArchiv("profile/535.csv", "-", 5, 35)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(a)
}

func TestArchivRPlus1(t *testing.T) {
	// k1 := Kombinacia{2, 9, 16, 22, 23, 24, 43}

	a, err := NewArchiv("profile/745_r.csv", "-", 7, 45)
	if err != nil {
		t.Fatal(err)
	}
	filters := Filters{
		NewFilterR1(0.00034197, 0.00034198, a.HHrx.Cisla, a.n),
		// NewFilterSTL1(n, min, max, cisla, fname),
		// NewFilterHHrx(n, min, max, HHrx),

		NewFilterR2(0.00000934, 0.00000935, a.Hrx.Cisla, a.n),
		// NewFilterSTL1(n, min, max, cisla, fname),
		// NewFilterHHrx(n, min, max, HHrx),

		NewFilterSucet(130, 140, a.n),

		// NewFilterCislovackyExact(a.n, Ints{3}, P),
		// NewFilterCislovackyExact(a.n, Ints{3}, N),
		// NewFilterCislovackyExact(a.n, Ints{3}, Pr),
		// NewFilterCislovackyExact(a.n, Ints{3}, Mc),
		// NewFilterCislovackyExact(a.n, Ints{3}, Vc),
		// NewFilterCislovackyExact(a.n, Ints{3}, C19),
		// NewFilterCislovackyExact(a.n, Ints{3}, C0),
		// NewFilterCislovackyExact(a.n, Ints{3}, Cc),
		// NewFilterCislovackyExact(a.n, Ints{3}, XcC),
		// NewFilterCislovackyExact(a.n, Ints{3}, CC),
		// NewFilterCislovackyRange(a.n, 2, 4, Pr),
		// NewFilterZhodaExact(ints, k, n),

		NewFilterSmernica(a.n, a.m, 0.803, 0.804),
		NewFilterKorelacia(a.n, a.m, 0.488, 0.489, a.K),

		NewFilterNtica(a.n, Tica{4, 0, 1, 0, 0, 0, 0}),
	}
	g := NewGenerator2(a, filters)
	g.Start()
	// for {
	// 	msg, run := g.Progress()
	// 	fmt.Println(msg)
	// 	if !run {
	// 		break
	// 	}
	// }
	g.Wait()
	time.Sleep(500 * time.Millisecond)

	f := NewFilter2(a, filters)
	f.Start()
	f.Wait()
	time.Sleep(500 * time.Millisecond)
	// for {
	// 	msg, run := f.Progress()
	// 	fmt.Println(msg)
	// 	if !run {
	// 		break
	// 	}
	// }
}
