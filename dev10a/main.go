package main

import (
	"fmt"
	"time"
)

// CGO_ENABLED=1 go run -race main.go
func main() {
	var counter int
	for range 1000 {
		go func() {
			counter++
		}()
	}

	time.Sleep(time.Second)
	
	fmt.Println("Counter:", counter)
}