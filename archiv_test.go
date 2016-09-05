package engine

import (
	"fmt"
	"path/filepath"
	"testing"
)

var testdir = "testdata"

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
