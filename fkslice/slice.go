package fkslice

import "reflect"

// FindInSlice takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
// original source https://golangcode.com/check-if-element-exists-in-slice
func FindInSlice(slice []interface{}, val interface{}) (int, bool) {

	for i, item := range slice {
		if reflect.DeepEqual(item, val) {
			return i, true
		}
	}
	return -1, false
}

// InterfaceSlice cast any slice to []interface{}
// original source https://stackoverflow.com/questions/12753805/type-converting-slices-of-interfaces
func InterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	// Keep the distinction between nil and empty slice input
	if s.IsNil() {
		return nil
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}
