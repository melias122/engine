package komb

// type Tica []byte
//
// func (t Tica) String() string {
// 	var buf bytes.Buffer
// 	for i, el := range t {
// 		if i > 0 {
// 			buf.WriteString(" ")
// 		}
// 		buf.WriteString(strconv.Itoa(int(el)))
// 	}
// 	return buf.String()
// }
//
// func Ntica(vect []byte) Tica {
//
// 	if len(vect) == 0 {
// 		return Tica{}
// 	}
//
// 	tice := make(Tica, len(vect))
// 	vect = append(vect, 0)
// 	var tica int
// 	for i := 0; i < len(vect); i++ {
//
// 		if i == len(vect)-2 {
// 			if tica > 0 {
// 				tice[tica]++
// 			} else {
// 				tice[0]++
// 			}
// 			break
// 		}
//
// 		c := int(vect[i]) - int(vect[i+1])
// 		if c == int(vect[i]) {
// 			if tica > 0 {
// 				tice[tica]++
// 			}
// 			break
// 		}
// 		if c < -1 {
// 			if tica == 0 {
// 				tice[0]++
// 			} else {
// 				tice[tica]++
// 				tica = 0
// 			}
// 		} else {
// 			tica++
// 		}
// 	}
// 	zero, vect = vect[len(vect)-1], vect[:len(vect)-1]
// 	return tice
// }
