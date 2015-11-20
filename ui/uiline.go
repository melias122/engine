package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/melias122/psl"
)

type Line interface {
	Filter() (psl.Filter, error)
	IsSet() bool
	Clear()
	Set(string, int)
}

type CifrovackyPanel struct {
	name   string
	le     [10]*walk.LineEdit
	filter func() (psl.Filter, error)
}

func (c CifrovackyPanel) Filter() (psl.Filter, error) {
	if c.filter == nil {
		return nil, fmt.Errorf("%s: nil filter", c.name)
	}
	filter, err := c.filter()
	if err != nil {
		return nil, fmt.Errorf("%s: %s", c.name, err)
	}
	return filter, nil
}

func (c CifrovackyPanel) String() string {
	var s string
	for i, le := range c.le {
		if i > 0 {
			s += " "
		}
		if strings.TrimSpace(le.Text()) == "" {
			s += "0"
		} else {
			s += le.Text()
		}
	}
	return s
}

func (c CifrovackyPanel) Clear() {
	for _, le := range c.le {
		if le != nil {
			le.SetText("")
		} else {
			log.Println(c.name, "nil line edit")
		}
	}
}

func (c CifrovackyPanel) IsSet() bool {
	var count int
	for i := range c.le {
		if c.le[i].Text() != "" {
			count++
		}
	}
	return count > 0
}

func (c CifrovackyPanel) Set(s string, i int) {
	if i >= 0 && i < len(c.le) {
		c.le[i].SetText(s)
	}
}

type StlNtica struct {
	name   string
	cb     [30]*walk.CheckBox
	filter func() (psl.Filter, error)
}

func (s StlNtica) Pozicie() []byte {
	var pozicie []byte
	for _, cb := range s.cb {
		if cb.Enabled() {
			var b byte
			if cb.Checked() {
				b = 1
			} else {
				b = 0
			}
			pozicie = append(pozicie, b)
		}
	}
	return pozicie
}

func (s StlNtica) Filter() (psl.Filter, error) {
	if s.filter == nil {
		return nil, fmt.Errorf("%s: nil filter", s.name)
	}
	filter, err := s.filter()
	if err != nil {
		return nil, fmt.Errorf("%s: %s", s.name, err)
	}
	return filter, nil
}

func (s StlNtica) IsSet() bool {
	var count int
	for i := range s.cb {
		if s.cb[i].Enabled() && s.cb[i].Checked() {
			count++
		}
	}
	return count > 0
}

func (s StlNtica) Clear() {
	for _, cb := range s.cb {
		if cb != nil {
			cb.SetChecked(false)
		} else {
			log.Println(STLNtica, "nil checkbox")
		}
	}
}

func (s StlNtica) Set(str string, i int) {

}

type UiLine struct {
	name     string
	rb1, rb2 *walk.RadioButton
	lines    []*walk.LineEdit
	filter   func() (psl.Filter, error)

	// cislovacky exact mode
	exactMode *walk.CheckBox

	// delta
	rbDelta0, rbDelta1, rbDelta2 *walk.RadioButton
}

func NewUiLine(name string, nLines int) *UiLine {
	return &UiLine{
		name:  name,
		lines: make([]*walk.LineEdit, nLines),
	}
}

func (u *UiLine) IsSet() bool {
	// delta
	if u.rbDelta0 != nil && u.rbDelta1 != nil && u.rbDelta2 != nil {
		if u.rbDelta0.Checked() {
			return false
		}
		if u.rbDelta1.Checked() {
			return true
		}
		if u.rbDelta2.Checked() {
			return true
		}
		return false
	}

	var (
		count int
		lines []*walk.LineEdit
	)
	switch len(u.lines) {
	case 1:
		lines = u.lines
	case 3:
		// exact mode
		if u.exactMode != nil && u.exactMode.Checked() {
			lines = u.lines[1:2]
			break
		}
		if u.lines[0].Enabled() {
			lines = u.lines[0:2]
		} else {
			lines = u.lines[1:3]
		}
	default:
		log.Println("UiLine.IsSet: ", u.name)
	}
	for _, line := range lines {
		if line.Text() != "" {
			count++
		}
	}
	return count > 0
}

func (u *UiLine) Set(s string, i int) {
	switch i {
	case 0:
		u.setMin(s)
	case 1:
		u.setR(s)
	case 2:
		u.setMax(s)
	default:
		log.Println(s, i)
		// panic("unsuported")
	}
}

func (u *UiLine) setR(s string) {
	var i int
	if u.lines[0].Enabled() && len(u.lines) == 3 {
		i = 1
	}
	u.lines[i].SetText(s)
}
func (u *UiLine) setMin(s string) {
	if len(u.lines) != 3 {
		return
	}
	var i int
	if u.lines[0].Enabled() {
		i = 0
	} else {
		i = 1
	}
	u.lines[i].SetText(s)
}
func (u *UiLine) setMax(s string) {
	if len(u.lines) != 3 {
		return
	}
	var i int
	if u.lines[0].Enabled() {
		i = 1
	} else {
		i = 2
	}
	u.lines[i].SetText(s)
}

func (u *UiLine) MinMax() (float64, float64, error) {
	var (
		minMax = [2]float64{0, 9999}
		err    error
		lines  []*walk.LineEdit
	)
	if u.lines[0].Enabled() {
		lines = u.lines[0:2]
	} else {
		lines = u.lines[1:3]
	}
	for i, line := range lines {
		s := strings.TrimSpace(line.Text())
		if len(s) > 0 {
			minMax[i], err = psl.ParseFloat(s)
			if err != nil {
				return 0, 0, err
			}
		} else {
			log.Printf("%s[%d]: prazdny vstup", u.name, i)
		}
	}
	if minMax[0] > minMax[1] {
		return 0, 0, fmt.Errorf("hodnota %f > %f", minMax[0], minMax[1])
	}
	return minMax[0], minMax[1], nil
}

func (u *UiLine) Filter() (psl.Filter, error) {
	if u.filter == nil {
		return nil, fmt.Errorf("%s: nil filter", u.name)
	}
	filter, err := u.filter()
	if err != nil {
		return nil, fmt.Errorf("%s: %s", u.name, err)
	}
	return filter, nil
}

func (u *UiLine) Clear() {
	if u.rbDelta0 != nil && u.rbDelta1 != nil && u.rbDelta2 != nil {
		u.rbDelta0.SetChecked(true)
	}
	if u.exactMode != nil {
		u.exactMode.SetChecked(false)
	}
	if u.rb1 != nil {
		u.rb1.SetChecked(true)
	}
	if u.rb2 != nil {
		u.rb2.SetChecked(false)
	}
	for _, line := range u.lines {
		if line != nil {
			line.SetText("")
		}
	}
}

func (u *UiLine) setLeft() {
	u.lines[0].SetEnabled(true)
	u.lines[2].SetEnabled(false)
	u.rb1.SetChecked(true)
	u.rb2.SetChecked(false)
}

func (u *UiLine) setRight() {
	u.lines[0].SetEnabled(false)
	u.lines[2].SetEnabled(true)
	u.rb1.SetChecked(false)
	u.rb2.SetChecked(true)
}

func (u *UiLine) setExact() {
	if u.exactMode.Checked() {
		u.lines[0].SetEnabled(false)
		u.lines[2].SetEnabled(false)
		u.rb1.SetEnabled(false)
		u.rb2.SetEnabled(false)
	} else {
		u.rb1.SetEnabled(true)
		u.rb2.SetEnabled(true)
		if u.rb1.Checked() {
			u.lines[0].SetEnabled(true)
		} else {
			u.lines[2].SetEnabled(true)
		}
	}
}

func UiLineToWidget(uiLine *UiLine, labelWidth int) Widget {
	widget := Composite{
		Layout: HBox{
			MarginsZero: true,
		},
		Children: []Widget{
			Label{
				MinSize: Size{Width: labelWidth},
				Text:    uiLine.name,
			},
			LineEdit{
				AssignTo: &uiLine.lines[0],
			},
			RadioButton{
				AssignTo:  &uiLine.rb1,
				Text:      "<=",
				OnClicked: uiLine.setLeft,
			},
			LineEdit{
				AssignTo: &uiLine.lines[1],
			},
			RadioButton{
				AssignTo:  &uiLine.rb2,
				Text:      "<=",
				OnClicked: uiLine.setRight,
			},
			LineEdit{
				AssignTo: &uiLine.lines[2],
				Enabled:  false,
			},
			ToolButton{
				Text:      "X",
				OnClicked: func() { uiLine.Clear() },
			},
		},
	}
	return widget
}

func UiLineToWidgetWithExact(uiLine *UiLine, labelWidth int) Widget {
	widget := Composite{
		Layout: HBox{
			MarginsZero: true,
		},
		Children: []Widget{
			Label{
				MinSize: Size{Width: labelWidth},
				Text:    uiLine.name,
			},
			LineEdit{
				AssignTo: &uiLine.lines[0],
			},
			RadioButton{
				AssignTo:  &uiLine.rb1,
				Text:      "<=",
				OnClicked: uiLine.setLeft,
			},
			LineEdit{
				AssignTo: &uiLine.lines[1],
			},
			RadioButton{
				AssignTo:  &uiLine.rb2,
				Text:      "<=",
				OnClicked: uiLine.setRight,
			},
			LineEdit{
				AssignTo: &uiLine.lines[2],
				Enabled:  false,
			},
			CheckBox{
				AssignTo:  &uiLine.exactMode,
				Text:      "E",
				OnClicked: uiLine.setExact,
			},
			ToolButton{
				Text:      "X",
				OnClicked: func() { uiLine.Clear() },
			},
		},
	}
	return widget
}

func UiLineToWidget2(uiLine *UiLine) Widget {
	widget := Composite{
		Layout: HBox{
			MarginsZero: true,
		},
		Children: []Widget{
			Label{
				Text:    uiLine.name,
				MinSize: Size{Width: 70},
			},
			LineEdit{
				AssignTo: &uiLine.lines[0],
			},
			ToolButton{
				Text:      "X",
				OnClicked: func() { uiLine.Clear() },
			},
		},
	}
	return widget
}

func UiLineToWidgetDelta(uiLine *UiLine) Widget {
	return Composite{
		Layout: HBox{
			MarginsZero: true,
		},
		Children: []Widget{
			Label{
				Text:    uiLine.name,
				MinSize: Size{Width: 70},
			},
			RadioButtonGroup{
				Buttons: []RadioButton{
					RadioButton{
						AssignTo: &uiLine.rbDelta0,
						Text:     "Vyp",
					},
					RadioButton{
						AssignTo: &uiLine.rbDelta1,
						Text:     "+",
					},
					RadioButton{
						AssignTo: &uiLine.rbDelta2,
						Text:     "-",
					},
				},
			},
		},
	}
}

func (u1 *UiLine) syncLines(u2 *UiLine) {
	if u1.exactMode.Checked() && u2.exactMode.Checked() {
		s := u2.lines[1].Text()
		r := strings.NewReader(s)
		p := psl.NewParser(r, n(), m())
		ints, err := p.ParseInts()
		if err != nil {
			log.Println(err)
			return
		}
		sort.Ints(ints)
		var str []string
		for i := range ints {
			ints[i] = n() - ints[i]
		}
		sort.Ints(ints)
		for _, i := range ints {
			str = append(str, strconv.Itoa(i))
		}
		u1.lines[1].SetText(strings.Join(str, ","))
	} else {
		min, max, err := u2.MinMax()
		if err != nil {
			log.Println(err)
			return
		}
		if int(min) < 0 {
			min = 0
		}
		if int(max) > n() {
			max = float64(n())
		}
		maxstr := strconv.Itoa(n() - int(min))
		minstr := strconv.Itoa(n() - int(max))
		u1.setMin(minstr)
		u1.setMax(maxstr)
	}
}

func UiLineToWidgetPair(u1, u2 *UiLine, labelWidth int) (Widget, Widget) {
	w1 := Composite{
		Layout: HBox{
			MarginsZero: true,
		},
		Children: []Widget{
			Label{
				MinSize: Size{Width: labelWidth},
				Text:    u1.name,
			},
			LineEdit{
				AssignTo: &u1.lines[0],
				OnEditingFinished: func() {
					// fmt.Println("text changed", u1)
					u2.syncLines(u1)
				},
			},
			RadioButton{
				AssignTo: &u1.rb1,
				Text:     "<=",
				OnClicked: func() {
					u1.setLeft()
					u2.syncLines(u1)
				},
			},
			LineEdit{
				AssignTo: &u1.lines[1],
				OnEditingFinished: func() {
					u2.syncLines(u1)
				},
			},
			RadioButton{
				AssignTo: &u1.rb2,
				Text:     "<=",
				OnClicked: func() {
					u1.setRight()
					u2.syncLines(u1)
				},
			},
			LineEdit{
				AssignTo: &u1.lines[2],
				Enabled:  false,
				OnEditingFinished: func() {
					u2.syncLines(u1)
				},
			},
			CheckBox{
				AssignTo: &u1.exactMode,
				Text:     "E",
				OnClicked: func() {
					u2.exactMode.SetChecked(!u2.exactMode.Checked())
					u1.setExact()
					u2.setExact()
				},
			},
			ToolButton{
				Text:      "X",
				OnClicked: func() { u1.Clear() },
			},
		},
	}

	w2 := Composite{
		Layout: HBox{
			MarginsZero: true,
		},
		Children: []Widget{
			Label{
				MinSize: Size{Width: labelWidth},
				Text:    u2.name,
			},
			LineEdit{
				AssignTo: &u2.lines[0],
				OnEditingFinished: func() {
					u1.syncLines(u2)
				},
			},
			RadioButton{
				AssignTo: &u2.rb1,
				Text:     "<=",
				OnClicked: func() {
					u2.setLeft()
					u1.syncLines(u2)
				},
			},
			LineEdit{
				AssignTo: &u2.lines[1],
				OnEditingFinished: func() {
					u1.syncLines(u2)
				},
			},
			RadioButton{
				AssignTo: &u2.rb2,
				Text:     "<=",
				OnClicked: func() {
					u2.setRight()
					u1.syncLines(u2)
				},
			},
			LineEdit{
				AssignTo: &u2.lines[2],
				Enabled:  false,
				OnEditingFinished: func() {
					u1.syncLines(u2)
				},
			},
			CheckBox{
				AssignTo: &u2.exactMode,
				Text:     "E",
				OnClicked: func() {
					u1.exactMode.SetChecked(!u1.exactMode.Checked())
					u1.setExact()
					u2.setExact()
				},
			},
			ToolButton{
				Text:      "X",
				OnClicked: func() { u2.Clear() },
			},
		},
	}
	return w1, w2
}
