package main

import (
	"fmt"
	"reflect"
	"testing"
)

// Clone a struct knowing it's type

type A struct {
	S string
}

// https://play.golang.org/p/XHaffI2jAV9
func cloneValue(t reflect.Type) reflect.Value {
	// If pointer, reduce to value
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return reflect.New(t)
}

func TestClone(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error", r)
		}
	}()

	a := &A{}

	newA := cloneValue(reflect.TypeOf(a)).Interface().(*A)

	a.S = "A"
	newA.S = "B"

	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", newA)

	if fmt.Sprintf("%T", a) != fmt.Sprintf("%T", newA) {
		t.Errorf("%T != %T", a, newA)
	}
}
