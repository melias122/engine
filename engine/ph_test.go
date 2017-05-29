package engine

import (
	"strconv"
	"testing"
)

func TestNewPh(t *testing.T) {
	ph := newph(1, 1, 5, 35)
	if ph.Pocet() != 0 {
		t.Errorf("Excepted: (0), Got: (%d)", ph.Pocet())
	}
	if ph.PocetNext() != 1 {
		t.Errorf("Excepted: (1), Got: (%d)", ph.PocetNext())
	}
	if ph.Hodnota() != 0.0 {
		t.Errorf("Excepted: (0.0), Got: (%f)", ph.Hodnota())
	}
	if strconv.FormatFloat(ph.HodnotaNext(), 'f', 10, 64) != "0.0000215629" {
		t.Errorf("Excepted: (0.0000215629), Got: (%.10f)", ph.HodnotaNext())
	}
}

func TestInc(t *testing.T) {
	ph := newph(1, 1, 5, 35)
	ph.inc()

	if ph.Pocet() != 1 {
		t.Errorf("Excepted: (1), Got: (%d)", ph.Pocet())
	}
	if ph.PocetNext() != 2 {
		t.Errorf("Excepted: (2), Got: (%d)", ph.PocetNext())
	}
	if strconv.FormatFloat(ph.Hodnota(), 'f', 10, 64) != "0.0000215629" {
		t.Errorf("Excepted: (0.0000215629), Got: (%f)", ph.Hodnota())
	}
	if strconv.FormatFloat(ph.HodnotaNext(), 'f', 10, 64) != "0.0000431258" {
		t.Errorf("Excepted: (0.0000431258), Got: (%.10f)", ph.HodnotaNext())
	}
}
