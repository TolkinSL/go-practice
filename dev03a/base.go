package main

import (
	"fmt"
	"unsafe"
)

const (
	MyX1 = 4
	MyX2 int = 7
	MyHello = "Hello"
	MyHelloStr string = "Hello Go!"
)

const (
	_ = iota
	Mon
	Tue
	Wed
)

func main() {

	// Область видимости
	var myY int
	myY = 10

	{
		myY = 20 // присвоение значения

		myY := "Go" // новая переменная внутри скобок затемнение
		_ = myY
	}

	_ = myY

	var str string
	fmt.Println(unsafe.Sizeof(str)) //16

	// type string struct {
	// 	Data uintptr // 8 byte
	// 	Len int // 8 byte
	// }
	
	
}
