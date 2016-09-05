package engine

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode/utf8"
)

func bytesToString(b []byte) string {
	var (
		space []byte
		buf   = make([]byte, 0, 128)
	)
	for _, u := range b {
		buf = append(buf, space...)
		buf = strconv.AppendUint(buf, uint64(u), 10)
		space = []byte(" ")
	}
	return string(buf[0:len(buf)])
}

func itoa(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func Ftoa(f float64) string {
	return ftoa(f)
}

func ftoa(f float64) string {
	buf := make([]byte, 0, 64)
	buf = strconv.AppendFloat(buf, f, 'g', -1, 64)
	for i, w := 0, 0; i < len(buf); i += w {
		runeValue, width := utf8.DecodeRune(buf[i:])
		if runeValue == '.' {
			buf[i] = ','
			break
		}
		w = width
	}
	return string(buf[:len(buf)])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// CopyFile copies a file from src to dst. If src and dst files exist, and are
// the same, then return success. Otherise, attempt to create a hard link
// between the two files. If that fail, copy the file contents from src to dst.
func CopyFile(dst, src string) (err error) {
	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if !sfi.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}
	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return
		}
	}
	// Dont want to create hard link..
	// if err = os.Link(src, dst); err == nil {
	// 	return
	// }

	// instead we copy file
	err = copyFileContents(src, dst)
	return
}

// copyFileContents copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}
