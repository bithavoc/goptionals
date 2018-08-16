package goptionals

import "testing"

func TestBools(t *testing.T) {
	b := Bool(true)
	if BoolValue(b) != true {
		t.Fatal()
	}
	if BoolValue(nil) != false {
		t.Fatal()
	}

	bl := []bool{true, false}
	blp := BoolSlice(bl)
	if *blp[0] != true {
		t.Fatal()
	}
	if *blp[1] != false {
		t.Fatal()
	}
	blv := BoolValueSlice(blp)
	if blv[0] != true {
		t.Fatal()
	}
	if blv[1] != false {
		t.Fatal()
	}
	if len(blp) != len(blv) {
		t.Fatal()
	}
	bm := BoolMap(map[string]bool{
		"v1": true,
		"v2": false,
	})
	if *bm["v1"] != true {
		t.Fatal()
	}
	if *bm["v2"] != false {
		t.Fatal()
	}
	bmp := BoolValueMap(bm)
	if bmp["v1"] != true {
		t.Fatal()
	}
	if bmp["v2"] != false {
		t.Fatal()
	}
}
