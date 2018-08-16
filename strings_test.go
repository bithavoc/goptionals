package goptionals

import "testing"

func TestString(t *testing.T) {
	b := String("a")
	if StringValue(b) != "a" {
		t.Fatal()
	}
	if StringValue(nil) != "" {
		t.Fatal()
	}

	bl := []string{"1", "2"}
	blp := StringSlice(bl)
	if *blp[0] != "1" {
		t.Fatal()
	}
	if *blp[1] != "2" {
		t.Fatal()
	}
	blv := StringValueSlice(blp)
	if blv[0] != "1" {
		t.Fatal()
	}
	if blv[1] != "2" {
		t.Fatal()
	}
	if len(blp) != len(blv) {
		t.Fatal()
	}
	bm := StringMap(map[string]string{
		"v1": "1",
		"v2": "2",
	})
	if *bm["v1"] != "1" {
		t.Fatal()
	}
	if *bm["v2"] != "2" {
		t.Fatal()
	}
	bmp := StringValueMap(bm)
	if bmp["v1"] != "1" {
		t.Fatal()
	}
	if bmp["v2"] != "2" {
		t.Fatal()
	}
}
