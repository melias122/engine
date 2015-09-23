package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/melias122/psl/filter"
)

type Line interface {
	Filter() (filter.Filter, error)
	IsSet() bool
	Clear()
	Set(string, int)
}

type CifrovackyPanel struct {
	name   string
	le     [10]*walk.LineEdit
	filter func() (filter.Filter, error)
}

func (c CifrovackyPanel) Filter() (filter.Filter, error) {
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
	filter func() (filter.Filter, error)
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

func (s StlNtica) Filter() (filter.Filter, error) {
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
	name string
	// min, max float64
	lines  []*walk.LineEdit
	filter func() (filter.Filter, error)
}

func NewUiLine(name string, nLines int) UiLine {
	return UiLine{
		name:  name,
		lines: make([]*walk.LineEdit, nLines),
	}
}

func (u UiLine) IsSet() bool {
	var (
		count int
		lines []*walk.LineEdit
	)
	switch len(u.lines) {
	case 1:
		lines = u.lines
	case 3:
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

func (u UiLine) Set(s string, i int) {
	switch i {
	case 0:
		u.setMin(s)
	case 1:
		u.setR(s)
	case 2:
		u.setMax(s)
	default:
		panic("unsuported")
	}
}

func (u UiLine) setR(s string) {
	var i int
	if u.lines[0].Enabled() && len(u.lines) == 3 {
		i = 1
	}
	u.lines[i].SetText(s)
}
func (u UiLine) setMin(s string) {
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
func (u UiLine) setMax(s string) {
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
			minMax[i], err = filter.ParseFloat(s)
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

func (u UiLine) Filter() (filter.Filter, error) {
	if u.filter == nil {
		return nil, fmt.Errorf("%s: nil filter", u.name)
	}
	filter, err := u.filter()
	if err != nil {
		return nil, fmt.Errorf("%s: %s", u.name, err)
	}
	return filter, nil
}

func (u UiLine) Clear() {
	for _, line := range u.lines {
		if line != nil {
			line.SetText("")
		}
	}
}

func UiLineToWidget(uiLine UiLine, labelWidth int) Widget {
	var (
		rb1, rb2 *walk.RadioButton
	)
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
				AssignTo: &rb1,
				Text:     "<=",
				OnClicked: func() {
					uiLine.lines[0].SetEnabled(true)
					uiLine.lines[2].SetEnabled(false)
					rb1.SetChecked(true)
					rb2.SetChecked(false)
				},
			},
			LineEdit{
				AssignTo: &uiLine.lines[1],
			},
			RadioButton{
				AssignTo: &rb2,
				Text:     "<=",
				OnClicked: func() {
					uiLine.lines[0].SetEnabled(false)
					uiLine.lines[2].SetEnabled(true)
					rb1.SetChecked(false)
					rb2.SetChecked(true)
				},
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

func UiLineToWidget2(uiLine UiLine) Widget {
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
