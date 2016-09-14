package engine

import (
	"flag"
	"fmt"
	"path/filepath"
	"testing"
)

var (
	testdir = "testdata"
	profile = flag.Bool("profile", false, "")
)

func TestProfile2080(t *testing.T) {
	if !*profile {
		t.SkipNow()
	}
	var (
		n, m    = 20, 80
		csvpath = filepath.Join(testdir, fmt.Sprintf("%d%d.csv", n, m))
	)
	_, err := NewArchiv(csvpath, testdir, n, m)
	if err != nil {
		t.Fatal(err)
	}
	// os.RemoveAll(a.WorkingDir)
}

func TestNewArchiv(t *testing.T) {
	var (
		n, m    = 5, 35
		csvpath = filepath.Join(testdir, fmt.Sprintf("%d%d.csv", n, m))
	)
	_, err := NewArchiv(csvpath, testdir, n, m)
	if err != nil {
		t.Fatal(err)
	}
	// os.RemoveAll(a.WorkingDir)
}
