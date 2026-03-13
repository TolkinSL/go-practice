package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 5)
	ch2 := make(chan int, 5)
	ch3 := make(chan int, 5)

	ch1 <- 1
	ch2 <- 2
	ch3 <- 3

	var val int

	select {
	case val = <- ch1:
		fmt.Println("ch1", val)
	case val = <- ch2:
		fmt.Println("ch2", val)
	case val = <- ch3:
		fmt.Println("ch3", val)
	}

	ch5 := make(chan int)
	ch6 := make(chan int)

	val = 37
	select {
	case val = <- ch5:
	case val = <- ch6:
	default:
		fmt.Println("Default for ch5, ch6:", val)
	}

	// Срабатывание сигнального канала stop, при закрытии канала, будет срабатывать

	task := make(chan int)
	stop := make(chan int)

	close(stop)

	select {
	case val = <- task:
		fmt.Println("Чтение task из канала")
	case <- stop:
		fmt.Println("Сработал сигнал stop при закрытии канала!")
		break
	}

	fmt.Println("before Select")

	go func() {
		time.Sleep(5 * time.Second)
		ch5 <- 5
	}()

	fmt.Println("Select wait chanel write")
	select {
	case val = <- ch5:
		fmt.Println("Goroutines write val:", val)
	}
}