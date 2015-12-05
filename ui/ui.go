package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/melias122/psl"
)

const (
	R1    = "ƩR 1-DO"
	R2    = "ƩR OD-DO"
	STL1  = "ƩSTL 1-DO"
	STL2  = "ƩSTL OD-DO"
	HRX   = "HRX"
	HHRX  = "HHRX"
	Sucet = "ƩKombinacie"

	Delta1 = "Δ(ƩR 1-DO - ƩSTL OD-DO)"
	Delta2 = "Δ(ƩR OD-DO - ƩSTL OD-DO)"

	P     = "P"
	N     = "N"
	Pr    = "Pr"
	Mc    = "Mc"
	Vc    = "Vc"
	C19   = "C19"
	C0    = "C0"
	cC    = "cC"
	Cc    = "Cc"
	CC    = "CC"
	Zhoda = "Zhoda"

	Kk = "Kk"
	Sm = "Sm"

	Ntica      = "Ntica"
	Xtica      = "Xtica"
	Xcisla     = "Xcisla"
	Cifrovacky = "Cifrovacky"
	STLNtica   = "STL Ntica"
)

//UI
var (
	mainWindow *walk.MainWindow
	nacitajPB  *walk.PushButton
	generujPB  *walk.PushButton
	filtrujPB  *walk.PushButton
	stopPB     *walk.PushButton
	archivrPB  *walk.PushButton
	mNE, nNE   *walk.NumberEdit
	riadokLE   *walk.LineEdit
	ucL        *walk.Label
	infoL      *walk.Label

	stlNtica   *StlNtica
	cifrovacky *CifrovackyPanel

	//Vars
	Archiv *psl.Archiv
	// Generator  psl.Generator = nil
	stop       = make(chan struct{})
	workingDir string

	lines = make(map[string]Line)
)

func UpperFilters() Widget {
	var widgets []Widget

	r1 := NewUiLine(R1, 3)
	r1.filter = func() (psl.Filter, error) {
		min, max, err := r1.MinMax()
		if err != nil {
			return nil, err
		}
		return psl.NewFilterR1(min, max, Archiv.HHrx.Cisla, n()), nil
	}

	r2 := NewUiLine(R2, 3)
	r2.filter = func() (psl.Filter, error) {
		min, max, err := r2.MinMax()
		if err != nil {
			return nil, err
		}
		return psl.NewFilterR2(min, max, Archiv.Hrx.Cisla, n()), nil
	}

	s1 := NewUiLine(STL1, 3)
	s1.filter = func() (psl.Filter, error) {
		min, max, err := s1.MinMax()
		if err != nil {
			return nil, err
		}
		return psl.NewFilterSTL1(min, max, Archiv.HHrx.Cisla, n()), nil
	}

	s2 := NewUiLine(STL2, 3)
	s2.filter = func() (psl.Filter, error) {
		min, max, err := s2.MinMax()
		if err != nil {
			return nil, err
		}
		return psl.NewFilterSTL2(min, max, Archiv.Hrx.Cisla, n()), nil
	}

	hhrx := NewUiLine(HHRX, 3)
	hhrx.filter = func() (psl.Filter, error) {
		min, max, err := hhrx.MinMax()
		if err != nil {
			return nil, err
		}
		return psl.NewFilterHHrx(min, max, Archiv.HHrx, n()), nil
	}

	hrx := NewUiLine(HRX, 3)
	hrx.filter = func() (psl.Filter, error) {
		min, max, err := hrx.MinMax()
		if err != nil {
			return nil, err
		}
		return psl.NewFilterHrx(min, max, Archiv.Hrx, n()), nil
	}

	sucet := NewUiLine(Sucet, 3)
	sucet.filter = func() (psl.Filter, error) {
		min, max, err := sucet.MinMax()
		if err != nil {
			return nil, err
		}
		return psl.NewFilterSucet(int(min), int(max), n()), nil
	}

	delta1 := NewUiLine(Delta1, 0)
	delta1.filter = func() (psl.Filter, error) {
		if delta1.rbDelta1.Checked() {
			return psl.NewFilterR1MinusSTL1(psl.POSSITIVE, Archiv.HHrx.Cisla, n()), nil
		} else if delta1.rbDelta2.Checked() {
			return psl.NewFilterR1MinusSTL1(psl.NEGATIVE, Archiv.HHrx.Cisla, n()), nil
		}
		return nil, errors.New("chyba")
	}

	delta2 := NewUiLine(Delta2, 0)
	delta2.filter = func() (psl.Filter, error) {
		if delta2.rbDelta1.Checked() {
			return psl.NewFilterR2MinusSTL2(psl.POSSITIVE, Archiv.Hrx.Cisla, n()), nil
		} else if delta2.rbDelta2.Checked() {
			return psl.NewFilterR2MinusSTL2(psl.NEGATIVE, Archiv.Hrx.Cisla, n()), nil
		}
		return nil, errors.New("chyba")
	}

	widgets = append(widgets, UiLineToWidgetDelta(delta1))
	widgets = append(widgets, UiLineToWidgetDelta(delta2))
	lines[delta1.name] = delta1
	lines[delta2.name] = delta2

	for _, line := range []*UiLine{
		r1,
		r2,
		s1,
		s2,
		hhrx,
		hrx,
		// delta1,
		// delta2,
		sucet,
	} {
		lines[line.name] = line
		widgets = append(widgets, UiLineToWidget(line, 60))
	}

	return Composite{
		Layout: VBox{
			MarginsZero: true,
		},
		Children: []Widget{
			Composite{
				Layout: Grid{
					MarginsZero: true,
					Columns:     2,
				},
				Children: widgets[:len(widgets)-1],
			},
			Composite{
				Layout:   HBox{MarginsZero: true},
				Children: widgets[len(widgets)-1:],
			},
		},
	}
}

func MiddleFilters() Widget {
	var widgets []Widget
	p := NewUiLine(P, 3)
	p.filter = func() (psl.Filter, error) {
		if p.exactMode.Checked() {
			s := p.lines[1].Text()
			return psl.NewFilterCislovackyExactFromString(s, psl.P, n(), m())
		} else {
			min, max, err := p.MinMax()
			if err != nil {
				return nil, err
			}
			return psl.NewFilterCislovackyRange(n(), int(min), int(max), psl.P), nil
		}
	}

	mc := NewUiLine(Mc, 3)
	mc.filter = func() (psl.Filter, error) {
		if mc.exactMode.Checked() {
			s := mc.lines[1].Text()
			return psl.NewFilterCislovackyExactFromString(s, psl.Mc, n(), m())
		} else {
			min, max, err := mc.MinMax()
			if err != nil {
				return nil, err
			}
			return psl.NewFilterCislovackyRange(n(), int(min), int(max), psl.Mc), nil
		}
	}

	c0 := NewUiLine(C0, 3)
	c0.filter = func() (psl.Filter, error) {
		if c0.exactMode.Checked() {
			s := c0.lines[1].Text()
			return psl.NewFilterCislovackyExactFromString(s, psl.C0, n(), m())
		} else {
			min, max, err := c0.MinMax()
			if err != nil {
				return nil, err
			}
			return psl.NewFilterCislovackyRange(n(), int(min), int(max), psl.C0), nil
		}
	}

	fCC := NewUiLine(CC, 3)
	fCC.filter = func() (psl.Filter, error) {
		if fCC.exactMode.Checked() {
			s := fCC.lines[1].Text()
			return psl.NewFilterCislovackyExactFromString(s, psl.CC, n(), m())
		} else {
			min, max, err := fCC.MinMax()
			if err != nil {
				return nil, err
			}
			return psl.NewFilterCislovackyRange(n(), int(min), int(max), psl.CC), nil
		}
	}

	nui := NewUiLine(N, 3)
	nui.filter = func() (psl.Filter, error) {
		if nui.exactMode.Checked() {
			s := nui.lines[1].Text()
			return psl.NewFilterCislovackyExactFromString(s, psl.N, n(), m())
		} else {
			min, max, err := nui.MinMax()
			if err != nil {
				return nil, err
			}
			return psl.NewFilterCislovackyRange(n(), int(min), int(max), psl.N), nil
		}
	}

	vc := NewUiLine(Vc, 3)
	vc.filter = func() (psl.Filter, error) {
		if vc.exactMode.Checked() {
			s := vc.lines[1].Text()
			return psl.NewFilterCislovackyExactFromString(s, psl.Vc, n(), m())
		} else {
			min, max, err := vc.MinMax()
			if err != nil {
				return nil, err
			}
			return psl.NewFilterCislovackyRange(n(), int(min), int(max), psl.Vc), nil
		}
	}

	fcC := NewUiLine(cC, 3)
	fcC.filter = func() (psl.Filter, error) {
		if fcC.exactMode.Checked() {
			s := fcC.lines[1].Text()
			return psl.NewFilterCislovackyExactFromString(s, psl.XcC, n(), m())
		} else {
			min, max, err := fcC.MinMax()
			if err != nil {
				return nil, err
			}
			return psl.NewFilterCislovackyRange(n(), int(min), int(max), psl.XcC), nil
		}
	}

	zhoda := NewUiLine(Zhoda, 3)
	zhoda.filter = func() (psl.Filter, error) {
		if zhoda.exactMode.Checked() {
			s := zhoda.lines[1].Text()
			return psl.NewFilterZhodaExactFromString(s, Archiv.K, n(), m())
		} else {
			min, max, err := zhoda.MinMax()
			if err != nil {
				return nil, err
			}
			return psl.NewFilterZhodaRange(int(min), int(max), Archiv.K, n()), nil
		}
	}

	pr := NewUiLine(Pr, 3)
	pr.filter = func() (psl.Filter, error) {
		if pr.exactMode.Checked() {
			s := pr.lines[1].Text()
			return psl.NewFilterCislovackyExactFromString(s, psl.Pr, n(), m())
		} else {
			min, max, err := pr.MinMax()
			if err != nil {
				return nil, err
			}
			return psl.NewFilterCislovackyRange(n(), int(min), int(max), psl.Pr), nil
		}
	}

	c19 := NewUiLine(C19, 3)
	c19.filter = func() (psl.Filter, error) {
		if c19.exactMode.Checked() {
			s := c19.lines[1].Text()
			return psl.NewFilterCislovackyExactFromString(s, psl.C19, n(), m())
		} else {
			min, max, err := c19.MinMax()
			if err != nil {
				return nil, err
			}
			return psl.NewFilterCislovackyRange(n(), int(min), int(max), psl.C19), nil
		}
	}

	fCc := NewUiLine(Cc, 3)
	fCc.filter = func() (psl.Filter, error) {
		if fCc.exactMode.Checked() {
			s := fCc.lines[1].Text()
			return psl.NewFilterCislovackyExactFromString(s, psl.Cc, n(), m())
		} else {
			min, max, err := fCc.MinMax()
			if err != nil {
				return nil, err
			}
			return psl.NewFilterCislovackyRange(n(), int(min), int(max), psl.Cc), nil
		}
	}

	kk := NewUiLine(Kk, 3)
	kk.filter = func() (psl.Filter, error) {
		min, max, err := kk.MinMax()
		if err != nil {
			return nil, err
		}
		return psl.NewFilterKorelacia(min, max, Archiv.K, n(), m()), nil
	}

	sm := NewUiLine(Sm, 3)
	sm.filter = func() (psl.Filter, error) {
		min, max, err := sm.MinMax()
		if err != nil {
			return nil, err
		}
		return psl.NewFilterSmernica(min, max, n(), m()), nil
	}

	pw, nw := UiLineToWidgetPair(p, nui, 35)
	lines[p.name] = p
	lines[nui.name] = nui

	mcw, vcw := UiLineToWidgetPair(mc, vc, 35)
	lines[mc.name] = mc
	lines[vc.name] = vc

	widgets = append(widgets, pw, mcw)
	for _, line := range []*UiLine{
		c0,
		fCC,
	} {
		widgets = append(widgets, UiLineToWidgetWithExact(line, 35))
		lines[line.name] = line
	}

	widgets = append(widgets, nw, vcw)
	for _, line := range []*UiLine{
		// nui,
		// vc,
		fcC,
		zhoda,

		pr,
		c19,
		fCc,
	} {
		widgets = append(widgets, UiLineToWidgetWithExact(line, 35))
		lines[line.name] = line
	}

	for _, line := range []*UiLine{
		kk,
		sm,
	} {
		widgets = append(widgets, UiLineToWidget(line, 35))
		lines[line.name] = line
	}

	return Composite{
		Layout: VBox{
			MarginsZero: true,
		},
		Children: []Widget{
			Composite{
				Layout: Grid{
					MarginsZero: true,
					Columns:     4,
				},
				Children: widgets[:len(widgets)-2],
			},
			Composite{
				Layout: HBox{
					MarginsZero: true,
				},
				Children: widgets[len(widgets)-2:],
			},
		},
	}
}

func DownFilters() Widget {
	var widgets []Widget
	povinne := NewUiLine("Povinne", 1)
	povinne.filter = func() (psl.Filter, error) {
		return psl.NewFilterPovinneFromString(povinne.lines[0].Text(), Archiv.K, n(), m())
	}

	povinneSTL := NewUiLine("Povinne STL", 1)
	povinneSTL.filter = func() (psl.Filter, error) {
		return psl.NewFilterPovinneSTLFromString(povinneSTL.lines[0].Text(), Archiv.K, n(), m())
	}

	zakazane := NewUiLine("Zakazane", 1)
	zakazane.filter = func() (psl.Filter, error) {
		return psl.NewFilterZakazaneFromString(zakazane.lines[0].Text(), Archiv.K, n(), m())
	}

	zakazaneSTL := NewUiLine("Zakazane STL", 1)
	zakazaneSTL.filter = func() (psl.Filter, error) {
		return psl.NewFilterZakazaneSTLFromString(zakazaneSTL.lines[0].Text(), Archiv.K, n(), m())
	}

	ntica := NewUiLine(Ntica, 1)
	ntica.filter = func() (psl.Filter, error) {
		tica, err := psl.ParseNtica(n(), ntica.lines[0].Text())
		if err != nil {
			return nil, err
		}
		return psl.NewFilterNtica(n(), tica), nil
	}

	xtica := NewUiLine(Xtica, 1)
	xtica.filter = func() (psl.Filter, error) {
		tica, err := psl.ParseXtica(n(), m(), xtica.lines[0].Text())
		if err != nil {
			return nil, err
		}
		return psl.NewFilterXtica(n(), m(), tica), nil
	}

	xcisla := NewUiLine(Xcisla, 1)
	xcisla.filter = func() (psl.Filter, error) {
		return psl.NewFilterXcislaFromString(xcisla.lines[0].Text(), n(), m())
	}

	for _, line := range []*UiLine{
		povinne,
		povinneSTL,
		zakazane,
		zakazaneSTL,
		ntica,
		xtica,
		xcisla,
	} {
		widgets = append(widgets, UiLineToWidget2(line))
		lines[line.name] = line
	}

	cifrovacky := new(CifrovackyPanel)
	cifrovacky.name = Cifrovacky
	cifrovacky.filter = func() (psl.Filter, error) {
		c, err := psl.ParseCifrovacky(cifrovacky.String(), n(), m())
		if err != nil {
			return nil, err
		}
		return psl.NewFilterCifrovacky(c, n(), m())
	}
	cifrovackyWidget := []Widget{
		Label{
			Text:    cifrovacky.name,
			MinSize: Size{Width: 70},
		},
	}
	for i := range cifrovacky.le {
		cifrovackyWidget = append(cifrovackyWidget, LineEdit{
			AssignTo: &cifrovacky.le[i],
		})
	}
	cifrovackyWidget = append(cifrovackyWidget, ToolButton{
		Text:      "X",
		OnClicked: func() { cifrovacky.Clear() },
	})
	// cifrovacky = &cifrovacky
	lines[cifrovacky.name] = cifrovacky

	widgets = append(widgets,
		Composite{
			Layout: HBox{
				MarginsZero: true,
			},
			Children: cifrovackyWidget,
		},
	)

	stlNtica = new(StlNtica)
	stlNtica.name = STLNtica
	stlNtica.filter = func() (psl.Filter, error) {
		tica, err := psl.ParseNtica(n(), ntica.lines[0].Text())
		if err != nil {
			return nil, errors.New("Nebola zadaná Ntica")
		}
		return psl.NewFilterSTLNtica(n(), tica, stlNtica.Pozicie()), nil
	}
	stlNticaWidget := []Widget{
		Label{
			Text:    stlNtica.name,
			MinSize: Size{Width: 70},
		},
	}
	for i := range stlNtica.cb {
		stlNticaWidget = append(stlNticaWidget, CheckBox{
			AssignTo: &stlNtica.cb[i],
			Text:     fmt.Sprintf(":%2d", i+1),
			Enabled:  false,
		})
	}
	stlNticaWidget = append(stlNticaWidget,
		ToolButton{
			Text:      "X",
			OnClicked: func() { stlNtica.Clear() },
		},
	)
	lines[STLNtica] = stlNtica

	widgets = append(widgets,
		Composite{
			Layout: HBox{
				MarginsZero: true,
			},
			Children: stlNticaWidget,
		},
	)

	return Composite{
		Layout: VBox{
			MarginsZero: true,
		},
		Children: widgets,
	}
}

func Buttons() Widget {
	return Composite{
		Layout: HBox{
			MarginsZero: true,
		},
		Children: []Widget{
			PushButton{
				AssignTo:  &generujPB,
				Text:      "Generuj r+1",
				Enabled:   false,
				OnClicked: func() { Generuj() },
			},
			PushButton{
				AssignTo:  &filtrujPB,
				Text:      "Filtruj r+1",
				Enabled:   false,
				OnClicked: func() { Filtruj() },
			},
			PushButton{
				AssignTo:  &stopPB,
				Text:      "Stop",
				Enabled:   false,
				OnClicked: func() { Stop() },
			},
			PushButton{
				AssignTo:  &archivrPB,
				Text:      "Archív r",
				Enabled:   false,
				OnClicked: func() { ArchivR() },
			},
			PushButton{
				Text: "Zmaž limity",
				OnClicked: func() {
					for _, l := range lines {
						l.Clear()
					}
				},
			},
			ToolButton{
				Text:      "?",
				OnClicked: help,
			},
		},
	}
}

func help() {
	msg := `Navod k zadavaniu hodnot.
Vo vseobecnosti plati...
(a) oddelovac hodnot je: "," (ciarka)
(b) oddelovac pozicii je: ";" (bodkociarka)
(c) oddelovac pozicii od hodnot je: ":" (dvojbodka)

(1) jedina hodnota: c1, c2, c3
(2) rozmedzie: c1-c2 // c1 < c2
(3) cislovacky + zhoda: P, N, Pr, Mc, Vc, C19, C0, cC, Cc, CC, Zh
(4) zadanie pozicie: p1:(1),(2),(3); p2:(1),(2),(3) // p1 > 0 
Dolezite je ze hodnoty sa oddeluju pomocou (a) a pozicie pomocou (b). Pozicie a hodnoty oddeluje (c)

Priklad...
Povinne/Zakazane: 			1, 2, 4-5, P, N
Povinne STL/Zakazane STL: 	1:1,2,3-5,P,N; 5:Zh,23

Ntica: 	5 0 0 0 0
Xtica: 	1 2 0 2 1
Xcisla: 1:1, 2-4; 2:3-4; 6:1

Mod E pre P, N, ..., Zh. V tomto mode je mozne do stredneho policka zadavat presne hodnoty.`
	walk.MsgBox(mainWindow, "?", msg, walk.MsgBoxIconInformation)
}

func n() int {
	return int(nNE.Value())
}

func m() int {
	return int(mNE.Value())
}

func ArchivR() {
	for k, v := range lines {
		var f float64
		switch k {
		case R1:
			f = Archiv.R1
		case R2:
			f = Archiv.R2
		case STL1:
			f = Archiv.S1
		case STL2:
			f = Archiv.S2
		case HRX:
			f = Archiv.Riadok.Hrx
		case HHRX:
			f = Archiv.Riadok.HHrx
		case Sucet:
			f = float64(Archiv.Sucet)
		case P:
			f = float64(Archiv.C[0])
		case N:
			f = float64(Archiv.C[1])
		case Pr:
			f = float64(Archiv.C[2])
		case Mc:
			f = float64(Archiv.C[3])
		case Vc:
			f = float64(Archiv.C[4])
		case C19:
			f = float64(Archiv.C[5])
		case C0:
			f = float64(Archiv.C[6])
		case cC:
			f = float64(Archiv.C[7])
		case Cc:
			f = float64(Archiv.C[8])
		case CC:
			f = float64(Archiv.C[9])
		case Zhoda:
			f = float64(Archiv.Zh)

		case Kk:
			f = Archiv.Kk
		case Sm:
			f = Archiv.Sm

		case Ntica:
			v.Set(Archiv.Ntica.String(), 1)
			continue
		case Xtica:
			v.Set(Archiv.Xtica.String(), 1)
			continue

		case Cifrovacky:
			for i, c := range Archiv.Cifrovacky {
				v.Set(strconv.Itoa(int(c)), i)
			}
			continue

		default:
			continue
		}
		s := strconv.FormatFloat(f, 'f', -1, 64)
		s = strings.Replace(s, ".", ",", 1)
		v.Set(s, 1)
	}
}

func NacitajSuborMW(parent *walk.MainWindow) (string, error) {
	var fileDialog walk.FileDialog
	accepted, err := fileDialog.ShowOpen(parent)
	if err != nil {
		return "", err
	}
	if accepted {
		return fileDialog.FilePath, nil
	} else {
		return "", errors.New("Nebol zadaný súbor")
	}
}

func NacitajSubor() {
	csvPath, err := NacitajSuborMW(mainWindow)
	if err != nil {
		infoL.SetText(err.Error())
	} else {
		done := make(chan error)
		go func() {
			infoL.SetText("Vytvarám Archív")
			Archiv, err = psl.NewArchiv(csvPath, workingDir, n(), m())
			done <- err
		}()
		go func() {
			err := <-done
			if err != nil {
				infoL.SetText(err.Error())
				return
			}
			// Lock
			nacitajPB.SetEnabled(false)
			nNE.SetEnabled(false)
			mNE.SetEnabled(false)

			// Unlock
			generujPB.SetEnabled(true)
			filtrujPB.SetEnabled(true)
			archivrPB.SetEnabled(true)
			for i := 0; i < n() && i < 30; i++ {
				stlNtica.cb[i].SetEnabled(true)
			}

			riadokLE.SetText(Archiv.K.String())
			ucL.SetText(ucL.Text() + strconv.Itoa(int(Archiv.Cislo)))
			infoL.SetText("Archív úspešne vytvorený")
		}()
	}
}

func Filters() (psl.Filters, error) {
	var (
		f psl.Filters
		e []error
	)
	for _, line := range lines {
		if line.IsSet() {
			filter, err := line.Filter()
			if err != nil {
				e = append(e, err)
			} else {
				f = append(f, filter)
			}
		}
	}
	if len(f) == 0 && len(e) == 0 {
		e = append(e, errors.New("Nebol zadaný žiadny filter?"))
	}
	if len(e) != 0 {
		var errorString string
		for _, err := range e {
			errorString += "\n" + err.Error()
		}
		return nil, errors.New(errorString)
	}
	return f, nil
}

func lockBtns() {
	generujPB.SetEnabled(false)
	filtrujPB.SetEnabled(false)
	stopPB.SetEnabled(true)
}

func unlockBtns() {
	generujPB.SetEnabled(true)
	filtrujPB.SetEnabled(true)
	stopPB.SetEnabled(false)
}

func run(g psl.Generator) {
	g.Start()

	// buttons
	go func() {
		lockBtns()
		defer unlockBtns()

		msg := g.Progress()
		for {
			select {
			// pripad ze vsetko prebehne ako ma
			case m, ok := <-msg:
				if !ok {
					return
				}
				infoL.SetText(m)
				// pouzivatel stlacil stop
			case <-stop:
				g.Stop()
			}
		}
	}()
}

func Stop() {
	go func() {
		stop <- struct{}{}
	}()
	stopPB.SetEnabled(false)
}

func Generuj() {
	filters, err := Filters()
	if err != nil {
		walk.MsgBox(mainWindow, "Chyba", err.Error(), walk.MsgBoxIconWarning)
		return
	}
	g := psl.NewGenerator2(Archiv, filters)
	run(g)
}

func Filtruj() {
	filters, err := Filters()
	if err != nil {
		walk.MsgBox(mainWindow, "Chyba", err.Error(), walk.MsgBoxIconWarning)
		return
	}
	f := psl.NewFilter2(Archiv, filters)
	run(f)
}

func Main() (err error) {

	if wd, err := os.Getwd(); err != nil {
		return err
	} else {
		workingDir = wd
	}

	if err := (MainWindow{
		AssignTo: &mainWindow,
		Title:    "Generator",
		Layout:   VBox{},
		MinSize:  Size{Width: 1400}, // 1397
		Size:     Size{Width: 1405},
		Children: []Widget{
			Composite{
				Layout: HBox{
					MarginsZero: true,
				},
				Children: []Widget{
					//N
					NumberEdit{
						AssignTo: &nNE,

						MinValue: 1,
						MaxValue: 30,
						Value:    5.0,
					},
					//M
					NumberEdit{
						AssignTo: &mNE,
						MinValue: 2,
						MaxValue: 90,
						Value:    35.0,
					},
					PushButton{
						AssignTo:  &nacitajPB,
						Text:      "Načítaj súbor",
						OnClicked: func() { NacitajSubor() },
					},
					LineEdit{
						AssignTo: &riadokLE,
						Enabled:  false,
					},
					Label{
						AssignTo: &ucL,
						Text:     "Uc: ",
					},
				},
			},
			UpperFilters(),
			MiddleFilters(),
			DownFilters(),
			Buttons(),
			Label{
				AssignTo: &infoL,
			},
		},
	}).Create(); err != nil {
		return err
	}
	for _, v := range lines {
		v.Clear()
	}
	mainWindow.Run()
	return err
}

func main() {
	if err := Main(); err != nil {
		log.Fatal(err)
	}
}
