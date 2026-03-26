package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Goroutune start")

	arr := []string{"msg1", "msg2", "msg3", "msg4"}
	for _, msg := range arr {

		time.Sleep(time.Second)
		out <- msg
	}

	fmt.Println("Goroutines end work!")
}

func main() {
	var wg sync.WaitGroup
	var readerWG sync.WaitGroup

	wg.Add(1)

	out := make(chan string)

	go worker(out, &wg)

	readerWG.Add(1)
	go func() {
		defer readerWG.Done()

		for msg := range out {
			fmt.Println("Goroutine message:", msg)
		}
	}()

	go func() {
		wg.Wait()
		close(out)
	}()
	
	readerWG.Wait()

	fmt.Println("Program end work!")
}
