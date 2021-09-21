package disjoint_set

import (
	"reflect"
	"testing"
)

func TestAllConnected(t *testing.T) {
	ds := New(100)
	for i := 0; i < 99; i++ {
		ds.Union(i, i+1)
		if ds.Size(0) != i+2 {
			t.Errorf("%d != %d", ds.Size(0), i+2)
		}
	}
}

func TestEqual(t *testing.T) {
	ds := New(100)
	ds2 := New(0)
	for i := 0; i < 100; i++ {
		ds2.Add()
	}
	if !reflect.DeepEqual(ds, ds2) {
		t.Errorf("ds != ds2")
	}
}

func TestInit(t *testing.T) {
	ds := New(100)
	for i := 0; i < 100; i++ {
		a := ds.Get(i)
		if a != i {
			t.Errorf("%d != %d", a, i)
		}
	}
}
