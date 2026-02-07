package main

import "fmt"

type ClosureMem struct {
	i *int
}

func (c ClosureMem) MyFunc() {
	*(c.i) += 1
}

func main() {
	var i int
	
	closure := func() {
		i++
	}

	closure()
	closure()
	fmt.Println(i)

	// In memory realize struct
	i = 0

	closureMem := ClosureMem{
		i: &i,
	}
	fmt.Println(i)
	closureMem.MyFunc()
	fmt.Println(i)
}
