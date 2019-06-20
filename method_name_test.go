package main

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

// Careful when trying to find a method name using the reflect package. From my
// understanding it seems you can lose "context" for the real function.
//
// "Method returns a function value"
// https://golang.org/pkg/reflect/#Value.Method
// It isn't specified if the function value refers to the "real" method
// function, or to something else.

type TestService struct{}

func (s *TestService) Foo() {}

func TestMethodName(t *testing.T) {

	// Struggle to find name based on reflect.Value
	serviceValue := reflect.ValueOf(&TestService{})
	fmt.Printf("runtime: %s\n", runtime.FuncForPC(serviceValue.Method(0).Pointer()).Name())
	fmt.Printf("reflect.Value: %s\n", serviceValue.Method(0).Type().Name())

	// Easy if you start with a reflect.Type
	serviceType := serviceValue.Type()
	fmt.Printf("reflect.Type: %s\n", serviceType.Method(0).Name)
}
