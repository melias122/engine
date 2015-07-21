package filter

import "github.com/melias122/psl/komb"

type povinne struct {
	n     int
	cisla []byte
}

func NewPovinne(n int, cisla []byte) Filter {
	return povinne{
		n:     n,
		cisla: cisla,
	}
}

func (p povinne) Check(k komb.Kombinacia) bool {
	return true
}

type povinneStl struct {
	cisla [][]bool
}

func NewPovinneStl(cisla [][]bool) Filter {
	return povinneStl{
		cisla: cisla,
	}
}

func (p povinneStl) Check(k komb.Kombinacia) bool {
	return true
}

// bool GLimits::checkPovinne(){
//     uint p_size = povinne.size();
//     uint r_size = result.size();
//     for(uint lev=0; lev < p_size && lev < r_size; ++lev){
//         if(result[lev] > povinne[lev])
//             return false;
//     }
//
//     if(level == r_size){
//         foreach (const uint &pov, povinne) {
//             if(!result.contains(pov))
//                 return false;
//         }
//     }
//     return true;
// }
