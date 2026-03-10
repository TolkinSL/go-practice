package main

import "fmt"

type Myinterface interface {
	method()
}

type T struct {
	Payload string
}

func (t *T) method() {
	t.Payload = "Value"
	fmt.Println("method() called")
}

func fn(m Myinterface) {
	m.method()
}

func main() {
	t := &T{}
	fn(t)

	fmt.Println("t.Payload:", t.Payload)
}