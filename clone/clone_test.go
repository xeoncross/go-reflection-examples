package main

import (
	"fmt"
	"reflect"
	"testing"
)

type A struct {
	S string
}

// https://play.golang.org/p/XHaffI2jAV9
func cloneValue(v reflect.Value) reflect.Value {
	aType := reflect.TypeOf(v)

	// If pointer, reduce to value
	if aType.Kind() == reflect.Ptr {
		aType = aType.Elem()
	}

	return reflect.New(aType)
}

func TestClone(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error", r)
		}
	}()

	a := &A{}

	newA := cloneValue(reflect.ValueOf(a)).Interface().(*A)

	a.S = "A"
	newA.S = "B"

	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", newA)

	// //
	// if fmt.Sprintf("%T", a) != fmt.Sprintf("%T", newA) {
	// 	t.Errorf("%T != %T", a, newA)
	// }

}
