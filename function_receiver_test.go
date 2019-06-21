package main

import (
	"reflect"
	"testing"
)

// "For a non-interface type T or *T, the returned Method's Type and Func
// fields describe a function whose first argument is the receiver."
//  - https://golang.org/pkg/reflect/#Type

type ReceiverFoo struct{}

func (r *ReceiverFoo) Run(word string) {}

func TestFunctionReceiver(t *testing.T) {

	receiverType := reflect.TypeOf(&ReceiverFoo{})
	methodType, _ := receiverType.MethodByName("Run")
	method := methodType.Func

	if method.Type().NumIn() != 2 {
		t.Errorf("%v has %d params\n", methodType.Name, method.Type().NumIn())
	}

	if method.Type().In(0).String() != "*main.ReceiverFoo" {
		t.Errorf("Unexpected type: %q", method.Type().In(0).String())
	}

	if method.Type().In(1).String() != "string" {
		t.Errorf("Unexpected type: %q", method.Type().In(1).String())
	}
}
