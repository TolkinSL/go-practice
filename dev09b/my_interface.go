package main

import "fmt"

func main() {
	var a any = "hello"

	// Приведение Типа
	s, ok := a.(string)
	if ok {
		fmt.Println("String:", s)
	} else {
		fmt.Println("Type is not string")
	}

	// Приведение Типа Type Switch 

	switch v := a.(type) {
	case int:
		fmt.Printf("Number: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	case error:
		fmt.Printf("Error: %v\n", v)
	default:
		fmt.Printf("Dont know\n")
	}
}
