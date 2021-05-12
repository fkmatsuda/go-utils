package fkslice

import "testing"

func TestFindInSlice_int(t *testing.T) {
	slice := []int{1, 5, 7, 8}
	var (
		idx   int
		found bool
	)
	idx, found = FindInSlice(InterfaceSlice(slice), 1)
	if !found {
		t.Error("\"1\" exists in slice")
	}
	if idx != 0 {
		t.Error("\"1\" must be in index 0")
	}
	idx, found = FindInSlice(InterfaceSlice(slice), 7)
	if !found {
		t.Error("\"7\" exists in slice")
	}
	if idx != 2 {
		t.Error("\"7\" must be in index 2")
	}
	idx, found = FindInSlice(InterfaceSlice(slice), 9)
	if found {
		t.Error("\"9\" not exists in slice")
	}
}

func TestFindInSlice_string(t *testing.T) {
	slice := []string{"alpha", "beta", "gamma", "theta"}
	var (
		idx   int
		found bool
	)
	idx, found = FindInSlice(InterfaceSlice(slice), "alpha")
	if !found {
		t.Error("\"alpha\" exists in slice")
	}
	if idx != 0 {
		t.Error("\"alpha\" must be in index 0")
	}
	idx, found = FindInSlice(InterfaceSlice(slice), "gamma")
	if !found {
		t.Error("\"gamma\" exists in slice")
	}
	if idx != 2 {
		t.Error("\"gamma\" must be in index 2")
	}
	idx, found = FindInSlice(InterfaceSlice(slice), "delta")
	if found {
		t.Error("\"delta\" not exists in slice")
	}
}

type testStruct struct {
	str   string
	value int
}

func TestFindInSlice_struct(t *testing.T) {
	slice := []testStruct{{"alpha", 1}, {"beta", 5}, {"gamma", 7}, {"theta", 8}}
	var (
		idx   int
		found bool
	)
	idx, found = FindInSlice(InterfaceSlice(slice), testStruct{"alpha", 1})
	if !found {
		t.Error("\"alpha\" exists in slice")
	}
	if idx != 0 {
		t.Error("\"alpha\" must be in index 0")
	}
	idx, found = FindInSlice(InterfaceSlice(slice), testStruct{"gamma", 7})
	if !found {
		t.Error("\"gamma\" exists in slice")
	}
	if idx != 2 {
		t.Error("\"gamma\" must be in index 2")
	}
	idx, found = FindInSlice(InterfaceSlice(slice), testStruct{"delta", 9})
	if found {
		t.Error("\"delta\" not exists in slice")
	}
}
