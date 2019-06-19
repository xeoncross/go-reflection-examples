package main

import (
	"reflect"
	"testing"
)

func TestZero(t *testing.T) {

	a := struct{ A string }{"Foo"}

	// These kind of look like they do the same thing,
	// but one gives you a pointer to a zeroed value
	// and one gives you a zeroed pointer (aka, nil pointer)

	zeroedValue := reflect.New(reflect.TypeOf(a))
	zeroedPointer := reflect.Zero(reflect.TypeOf(&a))

	// fmt.Printf("%#v\n", zeroedValue)
	// fmt.Printf("%#v\n", zeroedPointer)

	if !zeroedPointer.IsNil() {
		t.Errorf("%#v should be nil pointer", zeroedPointer)
	}

	if zeroedValue.IsNil() {
		t.Errorf("%#v should not be nil pointer", zeroedValue)
	}

}
