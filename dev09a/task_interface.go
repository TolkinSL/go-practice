package main

import "fmt"
import "reflect"

type MyInterface interface {}

type Example struct {
	Value string
}

func Example1() MyInterface {
	var e *Example

	return e
}

func isNil(i interface{}) bool {
	if i == nil {
		return true
	}

	v := reflect.ValueOf(i)
	fmt.Printf("%#v\n",v)

	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.IsNil()
	}

	return false
}

func main() {
	e1 := Example1()
	// Типизированный nil - (*main.Example)(nil)
	fmt.Printf("%#v\n", e1)

	// Печать Value
	fmt.Printf("%v\n", e1)

	// Печать Type, он задан и равен *Example
	fmt.Printf("%T\n", e1)

	if e1 == nil {
		fmt.Println("интерфейс сам nil")
	} else if ex, ok := e1.(*Example); ok && ex == nil {
		fmt.Println("типизированный nil: *Example(nil)", ex, ok)
	}

	fmt.Println("---")
	fmt.Println(isNil(e1))
}
