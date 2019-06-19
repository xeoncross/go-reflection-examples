package main

import (
	"fmt"
	"reflect"
	"testing"
)

// https://stackoverflow.com/a/16179271/99923

// Example nested struct
type A struct {
	S    string
	Meta struct {
		Desc       string
		Properties map[string]string
		Users      []string
	}
}

func TestInit(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error", r)
		}
	}()

	a := &A{}

	newA := initializeStruct(reflect.TypeOf(a)).Interface().(*A)

	// TODO map creation not working...
	a.Meta.Properties["a"] = "A"
	newA.Meta.Properties["b"] = "B"

	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", newA)
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
