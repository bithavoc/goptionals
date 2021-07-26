package goptionals

import (
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	sec := time.Second
	b := Duration(sec)
	if DurationValue(b) != sec {
		t.Fatal()
	}
	zero := time.Duration(0)
	if DurationValue(nil) != zero {
		t.Fatal()
	}

	bl := []time.Duration{sec}
	blp := DurationSlice(bl)
	if *blp[0] != sec {
		t.Fatal()
	}
	blv := DurationValueSlice(blp)
	if blv[0] != sec {
		t.Fatal()
	}
	if len(blp) != len(blv) {
		t.Fatal()
	}
	min := time.Minute
	bm := DurationMap(map[string]time.Duration{
		"v1": min,
	})
	if *bm["v1"] != min {
		t.Fatal()
	}
	bmp := DurationValueMap(bm)
	if bmp["v1"] != min {
		t.Fatal()
	}
}
