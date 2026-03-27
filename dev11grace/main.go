package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("worker start:", id)

	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				fmt.Println("worker jobs closed:", id)
				return
			}

			fmt.Println("worker", id, "processing job:", job)

			// эмуляция работы
			select {
			case <-time.After(2 * time.Second):
				fmt.Println("worker", id, "done job:", job)
			case <-ctx.Done():
				fmt.Println("worker", id, "cancel during job")
				return
			}

		case <-ctx.Done():
			fmt.Println("worker ctx cancel:", id)
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jobs := make(chan int)

	var wg sync.WaitGroup

	// запускаем 2 worker-а
	workers := 2
	wg.Add(workers)
	for i := 1; i <= workers; i++ {
		go worker(ctx, i, jobs, &wg)
	}

	// producer
	go func() {
		i := 1
		for {
			select {
			case jobs <- i:
				fmt.Println("send job:", i)
				i++
				time.Sleep(500 * time.Millisecond)

			case <-ctx.Done():
				fmt.Println("producer stop")
				close(jobs) // закрывает producer
				return
			}
		}
	}()

	// ловим сигналы ОС
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	<-sig // ждём Ctrl+C
	fmt.Println("\nshutdown signal received")

	// graceful shutdown
	cancel() // 1. сигнал всем goroutine
	// jobs закроется внутри producer

	wg.Wait() // 2. ждём worker-ов

	fmt.Println("all workers stopped")
}