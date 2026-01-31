package main

import (
	"fmt"
	"github.com/TolkinSL/go-practice/dev02/internal/usecase"
)

func main() {
	myHello := usecase.NewHello()
	fmt.Println(myHello.Say())
}
