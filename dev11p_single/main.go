package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(str string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Goroutine Start:", str)
	time.Sleep(time.Second)
	fmt.Println("Goroutines end work!", str)
}

func main() {
	var wg sync.WaitGroup

	for n := range 5 {
		wg.Add(1)
		str := fmt.Sprintf("goroutine %d", n)
		go worker(str, &wg)
	}

	wg.Wait()
	fmt.Println("Program end!")
}
