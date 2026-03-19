package main

import (
	"fmt"
	"time"
)

func main() {
	task := make(chan struct{})

	select {
	case <-task:
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout 3sec")
	}
}
