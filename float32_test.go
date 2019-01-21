package goptionals

import "testing"

func TestFloats32(t *testing.T) {
	b := Float32(1.5)
	if Float32Value(b) != 1.5 {
		t.Fatal()
	}
	if Float32Value(nil) != 0.0 {
		t.Fatal()
	}

	bl := []float32{1.5, 2.5}
	blp := Float32Slice(bl)
	if *blp[0] != 1.5 {
		t.Fatal()
	}
	if *blp[1] != 2.5 {
		t.Fatal()
	}
	blv := Float32ValueSlice(blp)
	if blv[0] != 1.5 {
		t.Fatal()
	}
	if blv[1] != 2.5 {
		t.Fatal()
	}
	if len(blp) != len(blv) {
		t.Fatal()
	}
	bm := Float32Map(map[string]float32{
		"v1": 1.5,
		"v2": 2.5,
	})
	if *bm["v1"] != 1.5 {
		t.Fatal()
	}
	if *bm["v2"] != 2.5 {
		t.Fatal()
	}
	bmp := Float32ValueMap(bm)
	if bmp["v1"] != 1.5 {
		t.Fatal()
	}
	if bmp["v2"] != 2.5 {
		t.Fatal()
	}
}
