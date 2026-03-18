package main

import (
	"fmt"
	// "sync"
	"time"
)

func main() {
	// var wg sync.WaitGroup
	c := make(chan int)
	arr := []int{2, 4, 6, 8, -1}

	// wg.Add(1)
	// go printCount(c, &wg)
	go printCount(c)

	for _, num := range arr {
		c <- num
		time.Sleep(time.Second)
	}

	// wg.Wait()
	fmt.Println("Done!")
	time.Sleep(1 * time.Second)
}

func printCount(c chan int) {
	var num int

	// for num > 0  deadlock так как при инициализации num = 0 !
	for num >= 0 {
		num = <-c
		fmt.Println("num:", num)
	}
}

// func printCount(c chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	var num int

// 	num = <-c
// 	fmt.Println("num:", num)
// }
