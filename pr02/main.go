package main

import (
	"fmt"

	"github.com/TolkinSL/go-practice/pr02/unpacker"
)

func main() {
	// s := "a3bc2"
	// s := "3"
	// s := "a@3bc2"
	s := ""
	u, err := unpacker.Unpack(s)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s: %s\n", s, u)
}