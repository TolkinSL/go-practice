package main

import (
	"fmt"

	"github.com/TolkinSL/go-practice/dev09/pkg/facebook"
	"github.com/TolkinSL/go-practice/dev09/pkg/stdout"
)

type Writer interface {
	Write(p []byte) (n int, err error)
}

func worker(w Writer, msg string) {
	n, err := w.Write([]byte(msg))
	if err != nil {
		fmt.Printf("worker: write error: %v\n", err)
	}

	fmt.Printf("worker: success write %d bytes\n", n)
}

func main() {
	fb, err := facebook.NewWriter("my_facebook.txt")
	if err != nil {
		fmt.Printf("error in main")
	}

	worker(fb, "Hello FB!")

	st, err := stdout.NewWriter("My message")
	if err != nil {
		fmt.Printf("error in main: %v\n", err)
	}

	worker(st, "Hello StdOut!")

	fb1, err := facebook.NewWriter("")
	if err != nil {
		fmt.Printf("error in main: %v\n", err)
	}
	
	_ = fb1
	// worker(fb1, "Hello FB!")
}
