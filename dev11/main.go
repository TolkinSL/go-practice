package main

import (
	"fmt"
	"time"
)

func worker(worker int, ch chan <- string) {
	for range time.Tick(time.Second) {
		ch <- fmt.Sprintf("Hi worker : %d", worker)
	}
}

func main() {
	tasks := make(chan string, 10)

	for i:= range 10 {
		go worker(i, tasks)
	}

	for str := range tasks {
		fmt.Println(str)
	}
}
