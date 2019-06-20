package main

import (
	"fmt"
	"reflect"
	"testing"
)

// Clone a new empty struct using a struct, pointer or reflect.Type.

type CloneFoo struct {
	S string
}

func cloneType(t reflect.Type) reflect.Value {
	// Dereference pointers
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return reflect.New(t)
}

func TestClone(t *testing.T) {

	a := &CloneFoo{}

	// In a real program this type would probably come from scaning an existing
	// structure for arguments or fields
	aType := reflect.TypeOf(a)

	// New reflect.Value can be asserted back to the underlying instance
	newA := cloneType(aType).Interface().(*CloneFoo)

	a.S = "A"
	newA.S = "B"

	// fmt.Printf("%#v\n", a)
	// fmt.Printf("%#v\n", newA)

	if fmt.Sprintf("%T", a) != fmt.Sprintf("%T", newA) {
		t.Errorf("%T != %T", a, newA)
	}
}
