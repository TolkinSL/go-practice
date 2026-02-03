package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) GetName() string {
	return u.Name
}

type MemUser32b struct {
	Id   int16
	Name string
	Age  int16
}

type MemUser24b struct {
	Name string
	Id   int16
	Age  int16
}

func main() {
	myUser := User{
		Id:   121,
		Name: "Vasia",
		Age:  37,
	}

	fmt.Println(myUser.GetName())

	myMem32b := MemUser32b{
		Id:   121,
		Name: "Vasia",
		Age:  37,
	}

	myMem24b := MemUser24b{
		Id:   121,
		Name: "Vasia",
		Age:  37,
	}

	_ = myMem32b
	fmt.Println("Размер структуры MemUser32b в байтах:", unsafe.Sizeof(MemUser32b{}))
	fmt.Println("Размер int16 в байтах:", unsafe.Sizeof(myMem32b.Id))
	fmt.Println("Размер string в байтах:", unsafe.Sizeof(myMem32b.Name))

	fmt.Println("Размер структуры MemUser24b в байтах:", unsafe.Sizeof(MemUser24b{}))
	fmt.Println("Размер структуры MemUser24b в байтах:", unsafe.Sizeof(myMem24b.Id))
	
}
