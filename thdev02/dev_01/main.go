package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello Go!")
}

func Basic1() {
	slice := []int{1, 2, 3, 4}
	for i, v := range slice {
		fmt.Println(i, v)
	}

	str := "Привет Go!"
	for i, v := range str {
		fmt.Println(i, v, string(v))
	}
}

func Basic2() {
	var s string = "abc"
	type myString string

	var myS myString = myString(s)
	_ = myS

	fmt.Println(s, "Type:", reflect.TypeOf(s))
	fmt.Println(myS, "Type:", reflect.TypeOf(myS))

}