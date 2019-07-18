package main

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

// In this example we are creating a struct instance and populating it.

type DecodeA struct {
	A string `json:"a"`
}

func TestJSONDecode(t *testing.T) {

	// Imagine this was a client providing a http.Request.Body JSON payload.
	r := strings.NewReader(`{"a":"foo"}`)

	// Assume we were handed a reflect.Type we need to use
	// We will just make one up here
	x := reflect.TypeOf(DecodeA{})

	// Start with a blank reflect.Value
	object := reflect.New(x)

	if object.CanInterface() {

		// Get the interface from the reflect.Value
		i := object.Interface()

		err := json.NewDecoder(r).Decode(&i)
		if err != nil {
			t.Error(err)
		}

		// fmt.Printf("%#v\n", object.Interface().(*DecodeA))
		// fmt.Printf("%#v\n", i.(*DecodeA))

		// type assertion to get our actual instance back
		if a, ok := i.(*DecodeA); ok {
			// Was the field actually set correctly?
			if a.A != "foo" {
				t.Errorf("Failed to set struct field: %q", a.A)
			}
		}

	} else {
		t.Error("Can't Interface reflect.Type")
	}
}
