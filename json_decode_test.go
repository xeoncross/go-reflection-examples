package main

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

// In this example we are creating a struct on-the-fly and populating it.
// Imagine this was a client providing a http.Request.Body JSON payload.

type DecodeA struct {
	A string `json:"a"`
}

func TestJSONDecode(t *testing.T) {

	// Assume we were handed a reflect.Type we need to use
	// We will just make one up here
	x := reflect.TypeOf(DecodeA{})

	// Start with a blank reflect.Value
	object := reflect.New(x)

	if object.CanInterface() {
		i := object.Interface()
		err := json.NewDecoder(strings.NewReader(`{"a":"foo"}`)).Decode(&i)
		if err != nil {
			t.Error(err)
		}

		// fmt.Printf("%#v\n", object.Interface().(*DecodeA))
		// fmt.Printf("%#v\n", i.(*DecodeA))

		if a, ok := i.(*DecodeA); ok {
			if a.A != "foo" {
				t.Errorf("Failed to set struct field: %q", a.A)
			}
		}

	} else {
		t.Error("Can't Interface reflect.Type")
	}
}
