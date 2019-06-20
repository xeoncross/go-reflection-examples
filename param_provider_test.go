package main

import (
	"fmt"
	"reflect"
	"testing"
)

// Read the parameters of a struct pointer method and generate empty structs
// that can be used to successfully call it.

type ProviderA struct {
	a string
}

type ProviderFoo struct{}

// Run demos needing struct, struct pointer, and string
func (p ProviderFoo) Run(a ProviderA, ptr *ProviderA, foo string) string {
	return a.a + ptr.a + foo
}

func TestParamProvider(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error", r)
		}
	}()

	handler := &ProviderFoo{}
	method := reflect.ValueOf(handler).MethodByName("Run")
	in := make([]reflect.Value, method.Type().NumIn())

	for i := 0; i < method.Type().NumIn(); i++ {
		paramType := method.Type().In(i)

		// Create a new instance of each param
		var object reflect.Value

		switch paramType.Kind() {
		case reflect.Struct:
			object = newReflectType(paramType).Elem()
		case reflect.Ptr:
			object = newReflectType(paramType)
		case reflect.String:
			object = reflect.New(paramType).Elem()
		default:
			t.Errorf("Unknown type: %s", paramType.Kind().String())
		}

		in[i] = object
	}

	if method.Type().NumOut() != 1 {
		t.Errorf("Too many return values")
	}

	response := method.Call(in)
	if response[0].String() != "" {
		t.Errorf("Unexpected result: %q", response[0].String())
	}
}

func newReflectType(t reflect.Type) reflect.Value {
	// Dereference pointers
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return reflect.New(t)
}
