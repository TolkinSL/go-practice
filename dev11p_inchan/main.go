package main

import (
	"fmt"
	"sync"
	"time"
)

func workerIn(in chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Goroutune start")

	for msg := range in {
		fmt.Println("Message receive:", msg)
		time.Sleep(time.Second)
	}

	fmt.Println("Goroutune end")
}

func main() {
	var wg sync.WaitGroup
	in := make(chan string)

	wg.Add(1)
	go workerIn(in, &wg)

	in <- "Msg1"
	in <- "Msg2"
	in <- "Msg3"
	in <- "Msg4"
	in <- "Msg5"
	close(in)

	wg.Wait()
	fmt.Println("Program end!")
}
