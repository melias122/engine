// Copyright 2013 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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
	"github.com/melias122/psl/archiv"
	"github.com/melias122/psl/filter"
	"github.com/melias122/psl/generator"
	"github.com/melias122/psl/num"
)

const (
	R1    = "ƩR 1-DO"
	R2    = "ƩR OD-DO"
	STL1  = "ƩSTL 1-DO"
	STL2  = "ƩSTL OD-DO"
	HRX   = "HRX"
	HHRX  = "HHRX"
	Sucet = "ƩKombinacie"

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
	Cifrovacky = "Cifrovacky"
	STLNtica   = "STL Ntica"
)

type Ui struct {
	//UI
	mainWindow *walk.MainWindow
	nacitajPB  *walk.PushButton
	generujPB  *walk.PushButton
	filtrujPB  *walk.PushButton
	archivrPB  *walk.PushButton
	mNE, nNE   *walk.NumberEdit
	riadokLE   *walk.LineEdit
	ucL        *walk.Label
	infoL      *walk.Label

	stlNtica   *StlNtica
	cifrovacky *CifrovackyPanel

	//Vars
	Archiv     *archiv.Archiv
	workingDir string

	lines map[string]Line
}

func (u *Ui) UpperFilters() Widget {
	var widgets []Widget

	r1 := NewUiLine(R1, 3)
	r1.filter = func() (filter.Filter, error) {
		min, max, err := r1.MinMax()
		if err != nil {
			return nil, err
		}
		return filter.NewR(u.n(), min, max, u.Archiv.HHrx.Cisla, r1.name), nil
	}

	r2 := NewUiLine(R2, 3)
	r2.filter = func() (filter.Filter, error) {
		min, max, err := r2.MinMax()
		if err != nil {
			return nil, err
		}
		return filter.NewR(u.n(), min, max, u.Archiv.Hrx.Cisla, r2.name), nil
	}

	s1 := NewUiLine(STL1, 3)
	s1.filter = func() (filter.Filter, error) {
		min, max, err := s1.MinMax()
		if err != nil {
			return nil, err
		}
		return filter.NewStl(u.n(), min, max, u.Archiv.HHrx.Cisla, s1.name), nil
	}

	s2 := NewUiLine(STL2, 3)
	s2.filter = func() (filter.Filter, error) {
		min, max, err := s2.MinMax()
		if err != nil {
			return nil, err
		}
		return filter.NewStl(u.n(), min, max, u.Archiv.Hrx.Cisla, s2.name), nil
	}

	hhrx := NewUiLine(HHRX, 3)
	hhrx.filter = func() (filter.Filter, error) {
		min, max, err := hhrx.MinMax()
		if err != nil {
			return nil, err
		}
		return filter.NewHrx(u.n(), min, max, u.Archiv.HHrx, hhrx.name), nil
	}

	hrx := NewUiLine(HRX, 3)
	hrx.filter = func() (filter.Filter, error) {
		min, max, err := hrx.MinMax()
		if err != nil {
			return nil, err
		}
		return filter.NewHrx(u.n(), min, max, u.Archiv.Hrx, hrx.name), nil
	}

	sucet := NewUiLine(Sucet, 3)
	sucet.filter = func() (filter.Filter, error) {
		min, max, err := sucet.MinMax()
		if err != nil {
			return nil, err
		}
		return filter.NewSucet(u.n(), int(min), int(max)), nil
	}

	for _, line := range []UiLine{
		r1,
		r2,
		s1,
		s2,
		hhrx,
		hrx,
		sucet,
	} {
		u.lines[line.name] = line
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

func (u *Ui) MiddleFilters() Widget {
	var widgets []Widget
	p := NewUiLine(P, 3)
	p.filter = func() (filter.Filter, error) {
		min, max, err := p.MinMax()
		if err != nil {
			return nil, err
		}
		return filter.NewCislovacky(u.n(), int(min), int(max), num.IsP, p.name), nil
	}

	mc := NewUiLine(Mc, 3)
	mc.filter = func() (filter.Filter, error) {
		min, max, err := mc.MinMax()
		if err != nil {
			return nil, err
		}
		return filter.NewCislovacky(u.n(), int(min), int(max), num.IsMc, mc.name), nil
	}

	c0 := NewUiLine(C0, 3)
	c0.filter = func() (filter.Filter, error) {
		min, max, err := c0.MinMax()
		if err != nil {
			return nil, err
		}
		return filter.NewCislovacky(u.n(), int(min), int(max), num.IsC0, c0.name), nil
	}

	fCC := NewUiLine(CC, 3)
	fCC.filter = func() (filter.Filter, error) {
		min, max, err := fCC.MinMax()
		if err != nil {
			return nil, err
		}
		return filter.NewCislovacky(u.n(), int(min), int(max), num.IsCC, fCC.name), nil
	}

	n := NewUiLine(N, 3)
	n.filter = func() (filter.Filter, error) {
		min, max, err := n.MinMax()
		if err != nil {
			return nil, err
		}
		return filter.NewCislovacky(u.n(), int(min), int(max), num.IsN, n.name), nil
	}

	vc := NewUiLine(Vc, 3)
	vc.filter = func() (filter.Filter, error) {
		min, max, err := vc.MinMax()
		if err != nil {
			return nil, err
		}
		return filter.NewCislovacky(u.n(), int(min), int(max), num.IsVc, vc.name), nil
	}

	fcC := NewUiLine(cC, 3)
	fcC.filter = func() (filter.Filter, error) {
		min, max, err := fcC.MinMax()
		if err != nil {
			return nil, err
		}
		return filter.NewCislovacky(u.n(), int(min), int(max), num.IscC, fcC.name), nil
	}

	zhoda := NewUiLine(Zhoda, 3)
	zhoda.filter = func() (filter.Filter, error) {
		min, max, err := zhoda.MinMax()
		if err != nil {
			return nil, err
		}
		return filter.NewZhoda(u.n(), int(min), int(max), u.Archiv.K), nil
	}

	pr := NewUiLine(Pr, 3)
	pr.filter = func() (filter.Filter, error) {
		min, max, err := pr.MinMax()
		if err != nil {
			return nil, err
		}
		return filter.NewCislovacky(u.n(), int(min), int(max), num.IsPr, pr.name), nil
	}

	c19 := NewUiLine(C19, 3)
	c19.filter = func() (filter.Filter, error) {
		min, max, err := c19.MinMax()
		if err != nil {
			return nil, err
		}
		return filter.NewCislovacky(u.n(), int(min), int(max), num.IsC19, c19.name), nil
	}

	fCc := NewUiLine(Cc, 3)
	fCc.filter = func() (filter.Filter, error) {
		min, max, err := fCc.MinMax()
		if err != nil {
			return nil, err
		}
		return filter.NewCislovacky(u.n(), int(min), int(max), num.IsCc, fCc.name), nil
	}

	kk := NewUiLine(Kk, 3)
	kk.filter = func() (filter.Filter, error) {
		min, max, err := kk.MinMax()
		if err != nil {
			return nil, err
		}
		return filter.NewKorelacia(u.n(), u.m(), min, max, u.Archiv.K), nil
	}

	sm := NewUiLine(Sm, 3)
	sm.filter = func() (filter.Filter, error) {
		min, max, err := sm.MinMax()
		if err != nil {
			return nil, err
		}
		return filter.NewSmernica(u.n(), u.m(), min, max), nil
	}

	for _, line := range []UiLine{
		p,
		mc,
		c0,
		fCC,

		n,
		vc,
		fcC,
		zhoda,

		pr,
		c19,
		fCc,

		kk,
		sm,
	} {
		widgets = append(widgets, UiLineToWidget(line, 35))
		u.lines[line.name] = line
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

func (u *Ui) DownFilters() Widget {
	var widgets []Widget
	povinne := NewUiLine("Povinne", 1)
	povinne.filter = func() (filter.Filter, error) {
		cisla, err := filter.ParseBytes(povinne.lines[0].Text())
		if err != nil {
			return nil, err
		}
		return filter.NewPovinne(u.n(), cisla), nil
	}
	povinneSTL := NewUiLine("Povinne STL", 1)
	povinneSTL.filter = func() (filter.Filter, error) {
		cisla, err := filter.ParseNBytes(u.n(), povinneSTL.lines[0].Text())
		if err != nil {
			return nil, err
		}
		return filter.NewPovinneStl(u.n(), cisla), nil
	}
	zakazane := NewUiLine("Zakazane", 1)
	zakazane.filter = func() (filter.Filter, error) {
		cisla, err := filter.ParseBytes(zakazane.lines[0].Text())
		if err != nil {
			return nil, err
		}
		return filter.NewZakazane(u.m(), cisla), nil
	}

	zakazaneSTL := NewUiLine("Zakazane STL", 1)
	zakazaneSTL.filter = func() (filter.Filter, error) {
		cisla, err := filter.ParseNBytes(u.n(), zakazaneSTL.lines[0].Text())
		if err != nil {
			return nil, err
		}
		return filter.NewZakazaneStl(u.n(), cisla), nil
	}

	ntica := NewUiLine(Ntica, 1)
	ntica.filter = func() (filter.Filter, error) {
		tica, err := filter.ParseNtica(u.n(), ntica.lines[0].Text())
		if err != nil {
			return nil, err
		}
		return filter.NewNtica(u.n(), tica), nil
	}

	xtica := NewUiLine(Xtica, 1)
	xtica.filter = func() (filter.Filter, error) {
		tica, err := filter.ParseXtica(u.n(), u.m(), xtica.lines[0].Text())
		if err != nil {
			return nil, err
		}
		return filter.NewXtica(u.n(), u.m(), tica), nil
	}

	for _, line := range []UiLine{
		povinne,
		povinneSTL,
		zakazane,
		zakazaneSTL,
		ntica,
		xtica,
	} {
		widgets = append(widgets, UiLineToWidget2(line))
		u.lines[line.name] = line
	}

	cifrovacky := new(CifrovackyPanel)
	cifrovacky.name = Cifrovacky
	cifrovacky.filter = func() (filter.Filter, error) {
		c, err := filter.ParseCifrovacky(u.n(), u.m(), cifrovacky.String())
		if err != nil {
			return nil, err
		}
		return filter.NewCifrovacky(u.n(), u.m(), c)
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
	// u.cifrovacky = &cifrovacky
	u.lines[cifrovacky.name] = cifrovacky

	widgets = append(widgets,
		Composite{
			Layout: HBox{
				MarginsZero: true,
			},
			Children: cifrovackyWidget,
		},
	)

	stlNtica := new(StlNtica)
	stlNtica.name = STLNtica
	stlNtica.filter = func() (filter.Filter, error) {
		tica, err := filter.ParseNtica(u.n(), ntica.lines[0].Text())
		if err != nil {
			return nil, errors.New("Nebola zadaná Ntica")
		}
		return filter.NewStlNtica(u.n(), tica, stlNtica.Pozicie()), nil
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
	u.stlNtica = stlNtica // kvoli zakrtavaniu...
	u.lines[STLNtica] = stlNtica

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

func (u *Ui) Buttons() Widget {
	return Composite{
		Layout: HBox{
			MarginsZero: true,
		},
		Children: []Widget{
			PushButton{
				AssignTo:  &u.generujPB,
				Text:      "Generuj r+1",
				Enabled:   false,
				OnClicked: func() { u.Generuj() },
			},
			PushButton{
				AssignTo:  &u.filtrujPB,
				Text:      "Filtruj r+1",
				Enabled:   false,
				OnClicked: func() { u.Filtruj() },
			},
			// PushButton{
			// Text: "Limity r+1",
			// },
			PushButton{
				AssignTo:  &u.archivrPB,
				Text:      "Archív r",
				Enabled:   false,
				OnClicked: func() { u.ArchivR() },
			},
			PushButton{
				Text: "Zmaž limity",
				OnClicked: func() {
					for _, l := range u.lines {
						l.Clear()
					}
				},
			},
		},
	}
}

func (u *Ui) n() int {
	return int(u.nNE.Value())
}

func (u *Ui) m() int {
	return int(u.mNE.Value())
}

func (u *Ui) ArchivR() {
	for k, v := range u.lines {
		var f float64
		switch k {
		case R1:
			f = u.Archiv.R1
		case R2:
			f = u.Archiv.R2
		case STL1:
			f = u.Archiv.S1
		case STL2:
			f = u.Archiv.S2
		case HRX:
			f = u.Archiv.Riadok.Hrx
		case HHRX:
			f = u.Archiv.Riadok.HHrx
		case Sucet:
			f = float64(u.Archiv.Sucet)
		case P:
			f = float64(u.Archiv.C[0])
		case N:
			f = float64(u.Archiv.C[1])
		case Pr:
			f = float64(u.Archiv.C[2])
		case Mc:
			f = float64(u.Archiv.C[3])
		case Vc:
			f = float64(u.Archiv.C[4])
		case C19:
			f = float64(u.Archiv.C[5])
		case C0:
			f = float64(u.Archiv.C[6])
		case cC:
			f = float64(u.Archiv.C[7])
		case Cc:
			f = float64(u.Archiv.C[8])
		case CC:
			f = float64(u.Archiv.C[9])
		case Zhoda:
			f = float64(u.Archiv.Zh)

		case Kk:
			f = u.Archiv.Kk
		case Sm:
			f = u.Archiv.Sm

		case Ntica:
			v.Set(u.Archiv.Ntica.String(), 1)
			continue
		case Xtica:
			v.Set(u.Archiv.Xtica.String(), 1)
			continue

		case Cifrovacky:
			for i, c := range u.Archiv.Cifrovacky {
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

func NacitajSubor(parent *walk.MainWindow) (string, error) {
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

func (u *Ui) NacitajSubor() {
	path, err := NacitajSubor(u.mainWindow)
	if err != nil {
		u.infoL.SetText(err.Error())
	} else {
		done := make(chan error)
		go func() {
			u.infoL.SetText("Vytvarám Archív")
			u.Archiv, err = archiv.Make(path, u.workingDir, u.n(), u.m())
			done <- err
		}()
		go func() {
			err := <-done
			if err != nil {
				u.infoL.SetText(err.Error())
			} else {
				// Lock
				u.nacitajPB.SetEnabled(false)
				u.nNE.SetEnabled(false)
				u.mNE.SetEnabled(false)

				// Unlock
				u.generujPB.SetEnabled(true)
				u.filtrujPB.SetEnabled(true)
				u.archivrPB.SetEnabled(true)
				for i := 0; i < u.n() && i < 30; i++ {
					u.stlNtica.cb[i].SetEnabled(true)
				}

				u.riadokLE.SetText(u.Archiv.K.String())
				u.ucL.SetText(u.ucL.Text() + strconv.Itoa(int(u.Archiv.Cislo)))
				u.infoL.SetText("Archív úspešne vytvorený")
			}
		}()
	}
}

func (u *Ui) Filters() (filter.Filters, error) {
	var (
		f filter.Filters
		e []error
	)
	for _, line := range u.lines {
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

func (u *Ui) Generuj() {
	filters, err := u.Filters()
	if err != nil {
		walk.MsgBox(u.mainWindow, "Chyba", err.Error(), walk.MsgBoxIconWarning)
		return
	}
	msg := make(chan string)
	go func() {
		u.infoL.SetText(fmt.Sprintf("Generovanie dokončené. %s", <-msg))
	}()
	go func() {
		// btn lock
		u.generujPB.SetEnabled(false)
		defer u.generujPB.SetEnabled(true)

		// info
		u.infoL.SetText("Generujem kombinácie")

		generator.GenerateKombinacie(u.n(), u.Archiv, filters, msg)
	}()
}

func (u *Ui) Filtruj() {
	filters, err := u.Filters()
	if err != nil {
		walk.MsgBox(u.mainWindow, "Chyba", err.Error(), walk.MsgBoxIconWarning)
		return
	}
	msg := make(chan string)
	go func() {
		msg := <-msg
		u.infoL.SetText(fmt.Sprintf("Filtrovanie dokončené. %s", msg))
	}()
	go func() {
		u.filtrujPB.SetEnabled(false)
		defer u.filtrujPB.SetEnabled(true)
		u.infoL.SetText("Vytváram filter pre r+1")

		generator.GenerateFilter(u.n(), u.Archiv, filters, msg)
	}()
}

func Main() error {

	var (
		workingDir, _ = os.Getwd()
		ui            = Ui{
			workingDir: workingDir,
			lines:      make(map[string]Line),
		}
	)

	_, err := MainWindow{
		AssignTo: &ui.mainWindow,
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
						AssignTo: &ui.nNE,

						MinValue: 1,
						MaxValue: 89,
						Value:    5.0,
					},
					//M
					NumberEdit{
						AssignTo: &ui.mNE,
						MinValue: 2,
						MaxValue: 99,
						Value:    35.0,
					},
					PushButton{
						AssignTo:  &ui.nacitajPB,
						Text:      "Načítaj súbor",
						OnClicked: func() { ui.NacitajSubor() },
					},
					LineEdit{
						AssignTo: &ui.riadokLE,
						Enabled:  false,
					},
					Label{
						AssignTo: &ui.ucL,
						Text:     "Uc: ",
					},
				},
			},
			ui.UpperFilters(),
			ui.MiddleFilters(),
			ui.DownFilters(),
			ui.Buttons(),
			Label{
				AssignTo: &ui.infoL,
			},
		},
	}.Run()
	return err
}

func main() {
	if err := Main(); err != nil {
		log.Fatal(err)
	}
}
