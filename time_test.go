package goptionals

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now()
	b := Time(now)
	if TimeValue(b) != now {
		t.Fatal()
	}
	zero := time.Time{}
	if TimeValue(nil) != zero {
		t.Fatal()
	}

	bl := []time.Time{now}
	blp := TimeSlice(bl)
	if *blp[0] != now {
		t.Fatal()
	}
	blv := TimeValueSlice(blp)
	if blv[0] != now {
		t.Fatal()
	}
	if len(blp) != len(blv) {
		t.Fatal()
	}
	bm := TimeMap(map[string]time.Time{
		"v1": now,
	})
	if *bm["v1"] != now {
		t.Fatal()
	}
	bmp := TimeValueMap(bm)
	if bmp["v1"] != now {
		t.Fatal()
	}
}
