package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
	// "time"
)

// CGO_ENABLED=1 go run -race main.go
func main() {
	runtime.GOMAXPROCS(4)
	// Mutex      → защищает counter
	// WaitGroup  → ждёт завершения goroutines
	var wg sync.WaitGroup
	var mu sync.Mutex

	now := time.Now()

	var counter int
	for range 10000 {
		wg.Add(1)

		go func() {
			defer wg.Done()

			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("Mutex Counter time:", time.Since(now))
	fmt.Println("Mutex Counter:", counter)

	// time.Sleep(time.Second)

	now = time.Now()
	var acounter atomic.Int64
	for range 10000 {
		wg.Add(1)
		go func() {
			defer wg.Done()

			acounter.Add(1)
		}()
	}

	wg.Wait()
	fmt.Println("Atomic Counter time:", time.Since(now))
	fmt.Println("Atomic Counter:", acounter.Load())

	// Выполнится один раз
	var once sync.Once

	counter = 0
	for range 10000 {
		wg.Add(1)

		go func() {
			defer wg.Done()

			once.Do(Init)
			
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("Mutex Counter:", counter)

}

func Init() {
	fmt.Println("Initial setup!")
}