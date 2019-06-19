package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestZero(t *testing.T) {

	a := struct{ A string }{"Foo"}

	// These kind of look like they do the same thing,
	// but one gives you a pointer to a zeroed value
	// and one gives you a zeroed pointer (aka, nil pointer)

	zeroValue := reflect.New(reflect.TypeOf(a))
	zeroedPointer := reflect.Zero(reflect.TypeOf(&a))

	fmt.Printf("%#v\n", zeroValue)
	fmt.Printf("%#v\n", zeroedPointer)

}
