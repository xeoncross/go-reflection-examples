package main

import (
	"fmt"
	"reflect"
	"testing"
)

// Create a new struct with fields which require initializing set to new empty
// values recursively.
//
// Not working.
//
// Based on: https://stackoverflow.com/a/16179271/99923

// Example nested struct
type InitializeFoo struct {
	S    string
	Meta struct {
		Desc       string
		Properties map[string]string
		Users      []string
	}
}

func TestInit(t *testing.T) {

	a := &InitializeFoo{}

	newA := initializeStruct(reflect.TypeOf(a)).Interface().(*InitializeFoo)

	// TODO map creation not working...
	// a.Meta.Properties["a"] = "A"
	// newA.Meta.Properties["b"] = "B"

	// fmt.Printf("%#v\n", a)
	// fmt.Printf("%#v\n", newA)

	if fmt.Sprintf("%T", a) != fmt.Sprintf("%T", newA) {
		t.Errorf("%T != %T", a, newA)
	}
}

func initializeStruct(t reflect.Type) reflect.Value {

	// If pointer, reduce to value
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	r := reflect.New(t)
	v := r.Elem()

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		ft := t.Field(i)
		switch ft.Type.Kind() {
		case reflect.Map:
			f.Set(reflect.MakeMap(ft.Type))
		case reflect.Slice:
			f.Set(reflect.MakeSlice(ft.Type, 0, 0))
		case reflect.Chan:
			f.Set(reflect.MakeChan(ft.Type, 0))
		case reflect.Struct:
			initializeStruct(ft.Type)
		case reflect.Ptr:
			fv := reflect.New(ft.Type.Elem())
			initializeStruct(ft.Type.Elem())
			f.Set(fv)
		default:
		}
	}

	return r
}
