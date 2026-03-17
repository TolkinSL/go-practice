package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup, worker int, ch chan <- string) {
	defer wg.Done()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <- ticker.C:
			ch <- fmt.Sprintf("Hi worker : %d", worker)
		case <- ctx.Done():
			return
		}
		
	}
}

func main() {
	tasks := make(chan string, 10)

	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	for i:= range 10 {
		wg.Add(1)
		go worker(ctx, &wg, i, tasks)
	}

	go func() {
		wg.Wait()
		close(tasks)
	}()

	for str := range tasks {
		fmt.Println(str)
	}

}
