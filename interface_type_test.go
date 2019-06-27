package main

import (
	"context"
	"reflect"
	"testing"
)

// The example belows shows how to get the reflect.Type of an interface.
// Also included is an example of checking that a reflect.Type matches a known
// interface (context.Context)
//
// Thanks to Petar Maymounkov for this trick
// https://groups.google.com/forum/#!topic/golang-nuts/qgJy_H2GysY

func TestInterface(t *testing.T) {

	// First, create a nil instance of the interface
	var interfaceInstance interface{} = (*context.Context)(nil)

	// The reflect.Type we get from our actual codebase
	// Here we just make one up
	mytype := reflect.TypeOf(context.TODO())

	if !mytype.Implements(reflect.TypeOf(interfaceInstance).Elem()) {
		t.Error("Type does not implement context.Context")
	}

	if !isContext(mytype) {
		t.Error("Type does not implement context.Context")
	}

}

func isContext(r reflect.Type) bool {
	return r.Implements(reflect.TypeOf((*context.Context)(nil)).Elem())
}
