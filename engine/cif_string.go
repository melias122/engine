// Code generated by "stringer -type Cif -trimprefix Cif"; DO NOT EDIT.

package engine

import "strconv"

const _Cif_name = "C0C1C2C3C4C5C6C7C8C9"

var _Cif_index = [...]uint8{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

func (i Cif) String() string {
	if i >= Cif(len(_Cif_index)-1) {
		return "Cif(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Cif_name[_Cif_index[i]:_Cif_index[i+1]]
}
