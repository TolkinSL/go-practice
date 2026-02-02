package main

import (
	"fmt"
	"sync"
)

type AddFunc func(int, int) int
type IncFunc func(int) int

func main() {
	{
		var f AddFunc
		fmt.Printf("%#v\n", f)
	}

	// Функция как тип
	var f AddFunc = add // Указатель на функцию
	result := f(2, 3)
	fmt.Printf("%#v\n", f)
	fmt.Println(result)

	var i IncFunc = inc
	result = i(2)
	fmt.Printf("%#v\n", i)
	fmt.Println(result)

	tf := Foo(add)
	result = tf(5)
	fmt.Println(result)

	// Анонимная функция 
	fmt.Println("\nАнонимная функция")

	{
		f := func(x, y int) int {
			return x + y
		}
		
		result = f(2, 4)
		fmt.Println(result)
	}

	// Анонимная асинхронная функция 
	fmt.Println("\nАнонимная асинхронная функция")

	{
		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			fmt.Println("асинхронная функция1")
		}()

		go func() {
			defer wg.Done()
			fmt.Println("асинхронная функция2")
		}()

		wg.Wait()
	}

	// Замыкание - это когда анонимная функция использует внешние переменнные
	// во внешней области видимости
	// Чистая анонимная функция без сайд эффектов, использует только свои переменные
	fmt.Println("\nЗамыкание")

	{
		var sum int

		mySum := func(num int) {
			fmt.Printf("%#v\n", sum)
			sum += num // переменная неявно захватывается по казателю
		}

		fmt.Println("sum", sum)
		mySum(2)
		mySum(3)
		fmt.Println("sum", sum)


	}
}

func Foo(f AddFunc) IncFunc {
	return func(x int) int { // Возврат анонимной функции
		return f(x, 1)	// которая внутри выполнит f(x, 1) и делается возврат значения
	}
}

func add(x int, y int) int {
	return x + y
}

func inc(x int) int {
	return x + 1
}
