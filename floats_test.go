package goptionals

import "testing"

func TestFloats(t *testing.T) {
	b := Float64(1.5)
	if Float64Value(b) != 1.5 {
		t.Fatal()
	}
	if Float64Value(nil) != 0.0 {
		t.Fatal()
	}

	bl := []float64{1.5, 2.5}
	blp := Float64Slice(bl)
	if *blp[0] != 1.5 {
		t.Fatal()
	}
	if *blp[1] != 2.5 {
		t.Fatal()
	}
	blv := Float64ValueSlice(blp)
	if blv[0] != 1.5 {
		t.Fatal()
	}
	if blv[1] != 2.5 {
		t.Fatal()
	}
	if len(blp) != len(blv) {
		t.Fatal()
	}
	bm := Float64Map(map[string]float64{
		"v1": 1.5,
		"v2": 2.5,
	})
	if *bm["v1"] != 1.5 {
		t.Fatal()
	}
	if *bm["v2"] != 2.5 {
		t.Fatal()
	}
	bmp := Float64ValueMap(bm)
	if bmp["v1"] != 1.5 {
		t.Fatal()
	}
	if bmp["v2"] != 2.5 {
		t.Fatal()
	}
}
