package main

import "fmt"

type Apple struct {
	Payload string
}

func (a Apple) Hello() {
	fmt.Println("Hi ", a.Payload)
}

type Banana1 struct {
	Apple Apple // Composition
}

type Banana2 struct {
	Payload string
	Apple // Embeding
}

func main() {
	a := Apple{
		Payload: "Apple",
	}

	b1 := Banana1{
		Apple: a,
	}

	b2 := Banana2{
		Payload: "Banana",
		Apple: a,
	}

	fmt.Printf("%+v\n", b1)
	fmt.Printf("%+v\n", b2)
	fmt.Printf("%+v\n", b2.Payload)
	fmt.Printf("%+v\n", b2.Apple)
	b2.Hello()
	b1.Apple.Hello()
}