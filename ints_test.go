package goptionals

import "testing"

func TestInt(t *testing.T) {
	b := Int(1)
	if IntValue(b) != 1 {
		t.Fatal()
	}
	if IntValue(nil) != 0 {
		t.Fatal()
	}

	bl := []int{1, 2}
	blp := IntSlice(bl)
	if *blp[0] != 1 {
		t.Fatal()
	}
	if *blp[1] != 2 {
		t.Fatal()
	}
	blv := IntValueSlice(blp)
	if blv[0] != 1 {
		t.Fatal()
	}
	if blv[1] != 2 {
		t.Fatal()
	}
	if len(blp) != len(blv) {
		t.Fatal()
	}
	bm := IntMap(map[string]int{
		"v1": 1,
		"v2": 2,
	})
	if *bm["v1"] != 1 {
		t.Fatal()
	}
	if *bm["v2"] != 2 {
		t.Fatal()
	}
	bmp := IntValueMap(bm)
	if bmp["v1"] != 1 {
		t.Fatal()
	}
	if bmp["v2"] != 2 {
		t.Fatal()
	}
}

func TestInt6464(t *testing.T) {
	b := Int64(1)
	if Int64Value(b) != 1 {
		t.Fatal()
	}
	if Int64Value(nil) != 0 {
		t.Fatal()
	}
	bl := []int64{1, 2}
	blp := Int64Slice(bl)
	if *blp[0] != 1 {
		t.Fatal()
	}
	if *blp[1] != 2 {
		t.Fatal()
	}
	blv := Int64ValueSlice(blp)
	if blv[0] != 1 {
		t.Fatal()
	}
	if blv[1] != 2 {
		t.Fatal()
	}
	if len(blp) != len(blv) {
		t.Fatal()
	}
	bm := Int64Map(map[string]int64{
		"v1": 1,
		"v2": 2,
	})
	if *bm["v1"] != 1 {
		t.Fatal()
	}
	if *bm["v2"] != 2 {
		t.Fatal()
	}
	bmp := Int64ValueMap(bm)
	if bmp["v1"] != 1 {
		t.Fatal()
	}
	if bmp["v2"] != 2 {
		t.Fatal()
	}
}
