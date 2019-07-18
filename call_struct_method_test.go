package main

import (
	"reflect"
	"testing"
)

// Foo contains "Min" to show we haven't lost context when calling "Add()"
type Foo struct {
	Min int
}

func (f *Foo) Add(a int, b int) int {
	return a + b + f.Min
}

func TestCallStructMethod(t *testing.T) {

	foo := &Foo{2}
	fooValue := reflect.ValueOf(foo)

	// Input for function
	in := []reflect.Value{
		reflect.ValueOf(3),
		reflect.ValueOf(5),
	}

	result := fooValue.Method(0).Call(in)

	if result[0].Int() != 10 {
		t.Errorf("Invalid result: %d\n", result[0].Int())
	}

}
