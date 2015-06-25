package komb

// func TestNew(t *testing.T) {
// 	k := New(5, 35)
// 	if k.String() != "" {
// 		t.Fatalf("Excepted: (%s), Have: (%s)", "", k.String())
// 	}
// 	if k.Len() != 0 {
// 		t.Fatalf("Excepted: (%d), Have: (%d)", 0, k.Len())
// 	}
// 	if k.C().String() != "0 0 0 0 0 0 0 0 0 0" {
// 		t.Fatalf("Excepted: (%s), Have: (%s)", "0 0 0 0 0 0 0 0 0 0", k.C().String())
// 	}
// 	if k.R1() != 0 {
// 		t.Fatalf("Excepted: (%f), Have: (%f)", 0, k.R1())
// 	}
// 	if k.R2() != 0 {
// 		t.Fatalf("Excepted: (%f), Have: (%f)", 0, k.R2())
// 	}
// 	if k.Sucet() != 0 {
// 		t.Fatalf("Excepted: (%d), Have: (%d)", 0, k.Sucet())
// 	}
// 	if k.Ntica().String() != "0 0 0 0 0" {
// 		t.Fatalf("Excepted: (%s), Have: (%s)", "0 0 0 0 0", k.Ntica().String())
// 	}
// 	if k.Xtica().String() != "0 0 0 0" {
// 		t.Fatalf("Excepted: (%s), Have: (%s)", "0 0 0 0", k.Xtica().String())
// 	}
// }

// var cisla = []*num.N{
// 	num.New(1, 5, 35),
// 	num.New(14, 5, 35),
// 	num.New(15, 5, 35),
// 	num.New(17, 5, 35),
// 	num.New(19, 5, 35),
// }

// func TestNtica(t *testing.T) {
// 	k := New(5, 35)
// 	for _, n := range cisla {
// 		k.Push(n)
// 	}
// 	if k.Ntica().String() != "3 1 0 0 0" {
// 		t.Fatalf("Excepted: (%s), Have: (%s)", "3 1 0 0 0", k.Ntica().String())
// 	}
// }
