package main

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

// Convert a struct with scaler values into a url.Values map

func structToURLValues(i interface{}) (url.Values, error) {
	values := url.Values{}
	iValue := reflect.ValueOf(i)

	if iValue.Kind() == reflect.Ptr {
		iValue = iValue.Elem()
	}
	iType := iValue.Type()

	for i := 0; i < iValue.NumField(); i++ {
		f := iValue.Field(i)
		switch f.Interface().(type) {
		case []string:
			// TODO
		default:
			values.Set(iType.Field(i).Name, fmt.Sprint(iValue.Field(i)))
		}
	}
	return values, nil
}

type StructForURLValues struct {
	A string
	B int
	C []string
}

func TestStructToMap(t *testing.T) {
	object := &StructForURLValues{
		A: "A",
		B: 12,
		// C: []string{"Foo", "Bar"},
	}

	got, err := structToURLValues(object)

	if err != nil {
		t.Error(err)
	}

	want := url.Values{"A": []string{"A"}, "B": []string{"12"}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected: %v\nGot: %v", want, got)
	}
}
