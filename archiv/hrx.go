package archiv

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/melias122/psl/hrx"
)

func (a *Archiv) HrxHHrx() error {
	header := []string{"p.c.", "HRX pre r+1", "dHRX diferencia s \"r\"", "presun z r do (r+1)cisla", "∑%ROD-DO", "∑%STLOD-DO od do", "∑ kombi od do",
		"Pocet ∑ kombi", "HHRX pre r+1", "dHHRX diferencia s \"r\"", "∑%R1-DO od do", "Teor. max. pocet", "∑%R1-DO", "∑%STL1-DO od do",
	}
	f, err := os.Create(fmt.Sprintf("%d%d/HrxHHrx_%d%d.csv", a.n, a.m, a.n, a.m))
	if err != nil {
		return err
	}
	defer f.Close()
	w := csv.NewWriter(f)
	w.Comma = ';'
	if err = w.Write(header); err != nil {
		return err
	}

	ch := hrx.GenerujPresun(a.Hrx.Presun(), a.n)

	i := 1
	for p := range ch {
		r := make([]string, 0, len(header))
		r = append(r,
			itoa(i),
			"0",
			"0",
			p.String())

		if err := w.Write(r); err != nil {
			return err
		}
		i++
	}

	w.Flush()
	return w.Error()
}
