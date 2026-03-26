package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(in <-chan string, out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Goroutine start:")

	for msg := range in {
		fmt.Println("In:", msg)
		workString := fmt.Sprintf("%s edit", msg)
		out <- workString
		time.Sleep(time.Second)
	}

	fmt.Println("Goroutines end work!")
}

func main() {
	var wg sync.WaitGroup
	var readerWG sync.WaitGroup

	in := make(chan string)
	out := make(chan string)
	
	wg.Add(1)
	go worker(in, out, &wg)

	readerWG.Add(1)
	go func() {
		defer readerWG.Done()

		for msg := range out {
			
			fmt.Println(msg)
		}
		
		fmt.Println("Reader end work!")
	}()

	in <- "work1"
	in <- "work2"
	
	close(in)
	
	go func() {
		wg.Wait()
		close(out)
	}()


	readerWG.Wait()
	fmt.Println("Program end work!")
}
