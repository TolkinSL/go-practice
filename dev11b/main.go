package main

import (
	"fmt"
	"sync"
)

func main() {
	// runtime.GOMAXPROCS(1)
	ch := make(chan int, 20)
	wg := sync.WaitGroup{}

	wg.Add(10)
	for i := range 10 {
		go func() {
			defer wg.Done()

			for val := range ch {
				fmt.Println("Worker:", i, "Value:", val)
			}
		}()
	}

	for i := range 20 {
		ch <- i
	}

	close(ch)

	wg.Wait()
	fmt.Println("Program end work!")
}
