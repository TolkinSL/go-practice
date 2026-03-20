package main

import (
	"fmt"
	"math/rand"
	"time"
)

const maxWaitSeconds = 5

// Эмуляция API
func randomWait() int {
	workSeconds := rand.Intn(maxWaitSeconds)
	time.Sleep(time.Duration(workSeconds) * time.Second)
	fmt.Println("workSeconds:", workSeconds)

	return workSeconds
}

func main() {
	mainSeconds := make(chan int)
	totalSeconds := 0
	
	start := time.Now()

	for range 5 {
		go func() {
			mainSeconds <- randomWait()
		}()
	}
	
	for range 5 {
		totalSeconds += <-mainSeconds
	}

	fmt.Println("\nProgram work", time.Since(start))
	fmt.Println("totalSeconds:", totalSeconds)
}
