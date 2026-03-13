package main

import (
	"fmt"
	"time"
	"unsafe"
)

func main() {
	ch := make(chan int, 10)

	ch <- 37
	ch <- 45

	fmt.Println("Длина канала", len(ch))
	fmt.Println("Полный размер канала", cap(ch))
	fmt.Println("Размер канала в стеке в байтах указатель", unsafe.Sizeof(ch))

	val := <-ch
	fmt.Println("val:", val)

	val = <-ch
	fmt.Println("val:", val)

	// val = <-ch // deadlock
	// fmt.Println("val:", val)

	// Небуферизированный канал
	var uch = make(chan int)

	go func() {
		for {
			val, ok := <-uch
			if !ok {
				fmt.Println("Канал закрыт")
				return
			}

			fmt.Println("Прием значения:", val)
			time.Sleep(time.Second)
		}
	}()

	fmt.Println("Отправка значений")
	uch <- 3
	uch <- 7
	uch <- 12
	uch <- 14
	fmt.Println("Значения отправлены")
	close(uch)

	time.Sleep(2 * time.Second)
}
