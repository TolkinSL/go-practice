// GODEBUG=schedtrace=1000 ./my_sheduler

package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(runtime.NumCPU())

	fmt.Println(runtime.GOMAXPROCS(0))
	// fmt.Println(runtime.GOMAXPROCS(2))
	
	go func() {
		time.Sleep(5 * time.Second)
	}()

	go func() {
		time.Sleep(5 * time.Second)
	}()

	fmt.Println("NumGoroutine", runtime.NumGoroutine())

	time.Sleep(15 * time.Second)
}
